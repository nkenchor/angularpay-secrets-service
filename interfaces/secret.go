package interfaces

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"
	"secret/application"
	"secret/domain/entity"

	"github.com/go-chi/chi/v5"
)

type SecretInterface struct {
	us application.SecretApplication
}

func NewSecret(us application.SecretApplication) SecretInterface {
	return SecretInterface{
		us: us,
	}
}

func (s *SecretInterface) CreateSecret(w http.ResponseWriter, r *http.Request) {
	var secret entity.SecretStruct
	err := decodeJSONBody(w, r, &secret)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			log.Println("400", err, &mr)
			ErrMessageController(w, r, "400", err.Error(), "entity.SecretStruct", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&mr)
		} else {
			log.Println("500", err, &mr)
			ErrMessageController(w, r, "500", &mr, "entity.SecretStruct", err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	postSecret, err := s.us.CreateSecret(secret)
	if err != nil {
		ErrMessageController(w, r, "400", postSecret, "entity.SecretStruct", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(postSecret)

}

func (s *SecretInterface) UpdateSecret(w http.ResponseWriter, r *http.Request) {
	var secret entity.SecretStruct
	ref := chi.URLParam(r, "reference")
	err := decodeJSONBody(w, r, &secret)
	if err != nil {
		var mr *malformedRequest
		if errors.As(err, &mr) {
			log.Println("400", err, &mr)
			ErrMessageController(w, r, "400", err.Error(), "entity.SecretStruct", err.Error())
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			json.NewEncoder(w).Encode(&mr)
		} else {
			log.Println("500", err, &mr)
			ErrMessageController(w, r, "500", &mr, "entity.SecretStruct", err.Error())
			http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		}
		return
	}
	putSecret, err := s.us.UpdateSecret(ref, secret)
	if err != nil {
		ErrMessageController(w, r, "400", putSecret, "entity.SecretStruct", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(putSecret)
}

func (s *SecretInterface) DeleteSecret(w http.ResponseWriter, r *http.Request) {
	var secret entity.SecretStruct
	ref := chi.URLParam(r, "reference")

	aSecret, err := s.us.DeleteSecret(ref, secret)
	if err != nil {
		ErrMessageController(w, r, "400", aSecret, "entity.SecretStruct", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	json.NewEncoder(w).Encode(aSecret)
}

func (s *SecretInterface) GetSecretByRef(w http.ResponseWriter, r *http.Request) {
	var secret entity.SecretStruct
	ref := chi.URLParam(r, "reference")

	aSecret, err := s.us.GetSecretByRef(ref, secret)
	if err != nil {
		ErrMessageController(w, r, "400", aSecret, "entity.SecretStruct", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(aSecret)
}

func (s *SecretInterface) GetAllSecret(w http.ResponseWriter, r *http.Request) {
	var secret entity.SecretStruct
	var secrets []entity.SecretStruct
	allSecret, err := s.us.GetAllSecret(secret, secrets)
	if err != nil {
		ErrMessageController(w, r, "400", allSecret, "entity.SecretStruct", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(allSecret)
}

func (s *SecretInterface) GetServiceSecretList(w http.ResponseWriter, r *http.Request) {
	var secret entity.SecretStruct
	ref := chi.URLParam(r, "reference")

	aSecret, err := s.us.GetServiceSecretList(ref, secret)
	if err != nil {
		ErrMessageController(w, r, "400", aSecret, "entity.SecretStruct", err.Error())
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(aSecret)
}
