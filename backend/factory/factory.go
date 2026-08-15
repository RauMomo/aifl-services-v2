package factory

import (
	"fmt"

	"aifl-backend/common/config"
	"aifl-backend/common/logger"
	"aifl-backend/common/models"
	"aifl-backend/repository"
	"aifl-backend/restful/client"
	"aifl-backend/restful/handler"
	"aifl-backend/restful/router"
	"aifl-backend/restful/service"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
	gormlogger "gorm.io/gorm/logger"
)

type Container struct {
	Config      *config.Config
	DB          *gorm.DB
	HTTPHandler func() interface{}
	Router      interface{}
	AuthService service.AuthService
}

type Factory struct {
	Config        *config.Config
	DB            *gorm.DB
	RedisRepo     repository.RedisRepository
	AdminRepo     repository.AdminRepository
	ContactRepo   repository.ContactRepository
	ProjectRepo   repository.ProjectRepository
	TestiRepo     repository.TestimonialRepository
	IKClient      client.ImageKitClient
	AuthService   service.AuthService
	ContactService service.ContactService
	ProjectService service.ProjectService
	TestiService   service.TestimonialService
	AuthHandler   *handler.AuthHandler
	ContactHandler *handler.ContactHandler
	ProjectHandler *handler.ProjectHandler
	TestiHandler   *handler.TestimonialHandler
}

func NewFactory() (*Factory, error) {
	cfg, err := config.LoadConfig()
	if err != nil {
		logger.Log.Warn("Failed to load Viper configuration", "error", err.Error())
	}

	dbPath := "aifl.db"
	if cfg != nil && cfg.Database.Path != "" {
		dbPath = cfg.Database.Path
	}

	gormConfig := &gorm.Config{
		Logger: gormlogger.Default.LogMode(gormlogger.Silent),
	}

	db, err := gorm.Open(sqlite.Open(dbPath), gormConfig)
	if err != nil {
		logger.Log.Error("❌ Failed to connect to SQLite database", "error", err.Error())
		return nil, fmt.Errorf("Failed to initialize database: %w", err)
	}
	logger.Log.Info("🟢 Database connected successfully", "engine", "SQLite", "file", dbPath)

	if err := db.AutoMigrate(
		&models.AdminUser{},
		&models.ContactUs{},
		&models.Project{},
		&models.Testimonial{},
	); err != nil {
		logger.Log.Error("❌ Database AutoMigrate failed", "error", err.Error())
		return nil, err
	}
	logger.Log.Info("✅ Database schema migrated successfully")

	redisRepo := repository.NewRedisRepository(cfg)
	adminRepo := repository.NewAdminRepository(db)
	contactRepo := repository.NewContactRepository(db)
	projectRepo := repository.NewProjectRepository(db)
	testiRepo := repository.NewTestimonialRepository(db)

	_ = adminRepo.SeedDefaultAdmin(cfg)

	ikClient := client.NewImageKitClient(cfg)

	authService := service.NewAuthService(adminRepo, redisRepo, cfg)
	contactService := service.NewContactService(contactRepo)
	projectService := service.NewProjectService(projectRepo, redisRepo, ikClient)
	testiService := service.NewTestimonialService(testiRepo, redisRepo)

	authHandler := handler.NewAuthHandler(authService)
	contactHandler := handler.NewContactHandler(contactService)
	projectHandler := handler.NewProjectHandler(projectService)
	testiHandler := handler.NewTestimonialHandler(testiService)

	return &Factory{
		Config:         cfg,
		DB:             db,
		RedisRepo:      redisRepo,
		AdminRepo:      adminRepo,
		ContactRepo:    contactRepo,
		ProjectRepo:    projectRepo,
		TestiRepo:      testiRepo,
		IKClient:       ikClient,
		AuthService:    authService,
		ContactService: contactService,
		ProjectService: projectService,
		TestiService:   testiService,
		AuthHandler:    authHandler,
		ContactHandler: contactHandler,
		ProjectHandler: projectHandler,
		TestiHandler:   testiHandler,
	}, nil
}

func (f *Factory) SetupRouter() interface{} {
	deps := &router.RouterDependencies{
		AuthHandler:        f.AuthHandler,
		ContactHandler:     f.ContactHandler,
		ProjectHandler:     f.ProjectHandler,
		TestimonialHandler: f.TestiHandler,
		AuthService:        f.AuthService,
		RedisRepo:          f.RedisRepo,
	}
	return router.RegisterRoutes(deps)
}
