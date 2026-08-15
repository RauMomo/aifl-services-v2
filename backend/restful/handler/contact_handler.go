package handler

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"

	"aifl-backend/common/models"
	"aifl-backend/proto/pb"
	"aifl-backend/restful/service"

	"google.golang.org/protobuf/proto"
)

type ContactHandler struct {
	contactService service.ContactService
}

func NewContactHandler(contactService service.ContactService) *ContactHandler {
	return &ContactHandler{contactService: contactService}
}

func (h *ContactHandler) CreateContact(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		respondJSON(w, http.StatusMethodNotAllowed, models.APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	var req models.ContactUs

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

	created, err := h.contactService.CreateContact(r.Context(), &req)
	if err != nil {
		respondJSON(w, http.StatusUnprocessableEntity, models.APIResponse{
			Success: false,
			Error:   err.Error(),
		})
		return
	}

	respondJSON(w, http.StatusCreated, models.APIResponse{
		Success: true,
		Message: "Pesan Anda telah berhasil dikirim. Tim AIFL akan menghubungi Anda segera.",
		Data:    created,
	})
}

func (h *ContactHandler) GetContacts(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		respondJSON(w, http.StatusMethodNotAllowed, models.APIResponse{
			Success: false,
			Error:   "Method not allowed",
		})
		return
	}

	contacts, err := h.contactService.GetAllContacts(r.Context())
	if err != nil {
		respondJSON(w, http.StatusInternalServerError, models.APIResponse{
			Success: false,
			Error:   err.Error(),
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
