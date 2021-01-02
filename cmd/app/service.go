package app

import (
	"avtoService/models"
	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
	"strconv"
)

func (receiver *server) AddService(w http.ResponseWriter, r *http.Request) {
	var svc models.Svc

	if !parseBody(r, &svc) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("AddService can't read json:(")
		return
	}

	if errs := validator.Validate(svc); errs != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("data is not  validate")
		return
	}

	err := receiver.svc.AddService(r.Context(), svc)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	err = writeJSONBody(w, &svc)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	return
}

func (receiver *server) ServiceList(w http.ResponseWriter, r *http.Request) {
	list, err := receiver.svc.ServiceList(r.Context())
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = writeJSONBody(w, &list)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	return
}

func (receiver *server) EditService(w http.ResponseWriter, r *http.Request) {
	vars  := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var svc models.Svc

	if !parseBody(r, &svc) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("EditService can't read json:(")
		return
	}

	if errs := validator.Validate(svc); errs != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("data is not  validate EditService")
		return
	}

	err = receiver.svc.EditService(r.Context(), svc, int64(id))
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = writeJSONBody(w, &svc)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	return
}

func (receiver *server) RemoveService(w http.ResponseWriter, r *http.Request) {
	vars  := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = receiver.svc.RemoveService(r.Context(), int64(id))
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = writeJSONBody(w, "Deleted")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	return
}