package app

import (
	"avtoService/psql"
	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"net/http"
)

type server struct {
	db     *pgxpool.Pool
	router *mux.Router
	svc    *psql.DB
}

func NewServer(db *pgxpool.Pool, router *mux.Router, svc *psql.DB) *server {
	return &server{db: db, router: router, svc: svc}
}

func (receiver *server) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	receiver.router.ServeHTTP(w, r)
}