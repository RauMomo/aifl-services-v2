package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"aifl-backend/db"
	"aifl-backend/logger"
	"aifl-backend/models"
	"aifl-backend/proto/pb"

	"google.golang.org/protobuf/proto"
)

type ContactHandler struct{}

func NewContactHandler() *ContactHandler {
	return &ContactHandler{}
}

// CreateContact handles POST /api/v1/contact
func (h *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondJSON(w, http.StatusMethodNotAllowed, models.APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	var req models.ContactUs

	// Support both JSON & Protobuf payloads
	if strings.Contains(r.Header.Get("Content-Type"), "application/x-protobuf") {
		body, err := io.ReadAll(r.Body)
		if err == nil {
			var pbReq pb.CreateContactRequestProto
			if proto.Unmarshal(body, &pbReq) == nil {
				req.Name = pbReq.GetName()
				req.Email = pbReq.GetEmail()
				req.Company = pbReq.GetCompany()
				req.Subject = pbReq.GetSubject()
				req.Message = pbReq.GetMessage()
			}
		}
	} else {
		_ = json.NewDecoder(r.Body).Decode(&req)
	}

	req.Name = strings.TrimSpace(req.Name)
	req.Email = strings.TrimSpace(req.Email)
	req.Message = strings.TrimSpace(req.Message)

	if req.Name == "" || req.Email == "" || req.Message == "" {
		respondJSON(w, http.StatusUnprocessableEntity, models.APIResponse{
			Success: false,
			Error:   "Nama, email, dan pesan wajib diisi",
		})
		return
	}

	req.ID = fmt.Sprintf("ct-%d", time.Now().UnixNano())
	req.Status = "unread"
	req.CreatedAt = time.Now()

	if err := db.DB.Create(&req).Error; err != nil {
		logger.Log.Error("Failed to insert contact in database", "error", err.Error())
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Gagal menyimpan data pesan",
		})
		return
	}

	logger.Log.Info("📩 New contact submission created", "id", req.ID, "name", req.Name, "email", req.Email)

	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Pesan Anda telah berhasil dikirim. Tim AIFL akan menghubungi Anda segera.",
		Data:    req,
	})
}

// GetContacts handles GET /api/v1/contact/list
func (h *ContactHandler) GetContacts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondJSON(w, http.StatusMethodNotAllowed, models.APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	var contacts []models.ContactUs
	if err := db.DB.Order("created_at desc").Find(&contacts).Error; err != nil {
		logger.Log.Error("Failed to query contacts from database", "error", err.Error())
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   "Gagal mengambil data kontak",
		})
		return
	}

	acceptHeader := r.Header.Get("Accept")
	if strings.Contains(acceptHeader, "application/x-protobuf") {
		pbContacts := make([]*pb.ContactUsProto, 0, len(contacts))
		for _, c := range contacts {
			pbContacts = append(pbContacts, &pb.ContactUsProto{
				Id:        c.ID,
				Name:      c.Name,
				Email:     c.Email,
				Company:   c.Company,
				Subject:   c.Subject,
				Message:   c.Message,
				Status:    c.Status,
				CreatedAt: c.CreatedAt.Format(time.RFC3339),
			})
		}
		pbResp := &pb.ContactListResponseProto{
			Success:  true,
			Message:  fmt.Sprintf("Retrieved %d contacts", len(contacts)),
			Contacts: pbContacts,
		}
		data, err := proto.Marshal(pbResp)
		if err == nil {
			w.Header().Set("Content-Type", "application/x-protobuf")
			w.WriteHeader(http.StatusOK)
			_, _ = w.Write(data)
			return
		}
	}

	respondJSON(w, http.StatusOK, models.APIResponse{
		Success: true,
		Message: fmt.Sprintf("Berhasil mengambil %d pesan kontak", len(contacts)),
		Data:    contacts,
	})
}

func respondJSON(w http.ResponseWriter, statusCode int, payload models.APIResponse) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	_ = json.NewEncoder(w).Encode(payload)
}
