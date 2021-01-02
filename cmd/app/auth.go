package app

import (
	"avtoService/models"
	"errors"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4"
	"gopkg.in/validator.v2"
	"log"
	"net/http"
	"strconv"
	"strings"
)


func (receiver *server) Registered(w http.ResponseWriter, r *http.Request) {
	var c models.Reg
	if !parseBody(r, &c) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("Registered can't read json:(")
		return
	}

	if errs := validator.Validate(c); errs != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("data is not  validate")
		return
	}


	err := receiver.svc.Registered(r.Context(), c)
	if err != nil {

		if strings.Contains(err.Error(), "ERROR: duplicate key value violates unique constraint") {
			http.Error(w, http.StatusText(http.StatusConflict), http.StatusConflict)
			log.Print(err)
			return
		}
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	err = writeJSONBody(w, &c)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	return
}

func (receiver *server) Login(w http.ResponseWriter, r *http.Request)  {
	var c models.Login
	if !parseBody(r, &c) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("Login can't read json:(")
		return
	}

	if errs := validator.Validate(c); errs != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("Login data is not  validate")
		return
	}

	_, err := receiver.svc.Login(r.Context(), c)
	if err != nil {
		if errors.Is(err, pgx.ErrNoRows) {
			http.Error(w, http.StatusText(http.StatusNotFound), http.StatusNotFound)
			log.Print("Not found person")
			return
		}

		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	err = writeJSONBody(w, "Пользователь найден")
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	return
}

func (receiver *server) ChangePass(w http.ResponseWriter, r *http.Request) {
	vars  := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	var pass models.ChangePass
	if !parseBody(r, &pass) {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("Login can't read json:(")
		return
	}

	if errs := validator.Validate(pass); errs != nil {
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		log.Printf("Login data is not  validate")
		return
	}

	err = receiver.svc.ChangePass(r.Context(), int64(id), pass)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}

	err = writeJSONBody(w, &pass)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}
	return
}

func (receiver *server) PersonalList(w http.ResponseWriter, r *http.Request) {
	list, err := receiver.svc.PersonalList(r.Context())
	if err != nil {
		log.Print(err)
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		return
	}

	err = writeJSONBody(w, list)
	if err != nil {
		http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
		log.Print(err)
		return
	}
	return
}

func (receiver *server) DeletePersonal(w http.ResponseWriter, r *http.Request) {
	vars  := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Println(err)
		http.Error(w, http.StatusText(http.StatusBadRequest), http.StatusBadRequest)
		return
	}

	err = receiver.svc.DeletePersonal(r.Context(), int64(id))
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