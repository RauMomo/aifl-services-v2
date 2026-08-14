# AIFL Golang REST API Backend

Production-ready Golang REST API backend with entities for **Contact Us**, **Projects**, and **Testimonials**, optimized for local development and direct deployment to **Back4App Containers (CaaS)**.

---

## 🛠️ Entities & API Endpoints

### 1. Contact Us (`/api/v1/contact`)

- `POST /api/v1/contact` - Submit a new consultation / contact request
- `GET /api/v1/contact/list` - Retrieve all submitted contact requests

#### Sample Payload (`POST /api/v1/contact`):

```json
{
  "name": "Budi Santoso",
  "email": "budi@example.com",
  "company": "PT Maju Sejahtera",
  "subject": "Diskusi Digitalisasi Operasional",
  "message": "Kami ingin berkonsultasi mengenai otomatisasi workflow perusahaan."
}
```

---

### 2. Projects (`/api/v1/projects`)

- `GET /api/v1/projects` - List all portfolio projects / case studies
- `POST /api/v1/projects` - Add a new project

#### Sample Payload (`POST /api/v1/projects`):

```json
{
  "title": "Enterprise Resource Control System",
  "description": "Sistem terpusat pengelola workflow & audit perusahaan.",
  "category": "Operational Control",
  "client_name": "PT Corporate Indonesia",
  "image_url": "/aifl-software-mockup.png",
  "tags": ["Golang", "Ember.js", "Docker"],
  "impact_metrics": ["Efisiensi +80%", "Zero Data Loss"],
  "is_featured": true
}
```

---

### 3. Testimonials (`/api/v1/testimonials`)

- `GET /api/v1/testimonials` - Retrieve client testimonials
- `POST /api/v1/testimonials` - Submit a new client review

#### Sample Payload (`POST /api/v1/testimonials`):

```json
{
  "author_name": "Hendry Prasetya",
  "position": "Head of Operations",
  "company": "PT Nusantara Synergy",
  "quote": "Solusi digital AIFL mengubah cara kerja tim kami secara efisien.",
  "rating": 5,
  "avatar_url": "https://images.unsplash.com/photo-1534528741775-53994a69daeb?auto=format&fit=crop&w=150&q=80"
}
```

---

## 🚀 Running Locally

```bash
# Navigate to the backend directory
cd backend

# Run the server
go run main.go
```

The server will start at `http://localhost:8200`.

Test the health check endpoint:

```bash
curl http://localhost:8200/health
```

---

## ☁️ Deploying to Back4App Containers

1. Push your repository to **GitHub**.
2. Log in to [Back4App Dashboard](https://backend.back4app.com).
3. Click **Build new app** -> Select **Containers**.
4. Connect your GitHub repository.
5. Set **Root Directory** to `backend` (if in a monorepo) or standard root.
6. Set **Port** to `8200`.
7. Click **Deploy**. Back4App will automatically build the `Dockerfile` and deploy your Golang REST API service with free SSL and auto-scaling!
