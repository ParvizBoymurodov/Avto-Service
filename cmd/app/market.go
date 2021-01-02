package app

import (
	"avtoService/models"
	"github.com/gorilla/mux"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
	"strconv"
)

func (receiver *server) AddMarket(w http.ResponseWriter, r *http.Request) {
	var  mrk models.Market

	if !parseBody(r, &mrk) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("AddMarket can't read json:(")
		return
	}

	if errs := validator.Validate(mrk); errs != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("data is not  validate")
		return
	}

	err := receiver.svc.AddMarket(r.Context(), mrk)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	err = writeJSONBody(w, &mrk)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	return
}

func (receiver *server) MarketList(w http.ResponseWriter, r *http.Request) {
	list, err := receiver.svc.MarketList(r.Context())
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	err = writeJSONBody(w, &list)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	return
}

func (receiver *server) RemoveMarket(w http.ResponseWriter, r *http.Request) {
	vars  := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = receiver.svc.RemoveMarket(r.Context(), int64(id))
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	err = writeJSONBody(w, "Deleted")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	return
}