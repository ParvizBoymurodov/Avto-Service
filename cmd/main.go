package main

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"log"
	"net"
	"net/http"


	"avtoService/cmd/app"
	"avtoService/models"
	"avtoService/psql"

	"github.com/gorilla/mux"
	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/rs/cors"
)

var conf models.Config

func init() {
	data, err := ioutil.ReadFile("./cmd/config.json")
	if err != nil {
		log.Fatalf("Contact gate config file: %v\n", err)
	}

	if err := json.Unmarshal(data, &conf); err != nil {
		log.Fatalf("Contact gate config file: %v\n", err)
	}
}

func main() {
	addr := net.JoinHostPort(conf.Host, conf.Port)
	start(addr, conf.Dsn)
}

func start(addr string, dsn string) {
	pool, err := pgxpool.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	router:=mux.NewRouter()
	answersSvc := psql.NewDB(pool)
	server := app.NewServer(
		pool,
		router,
		answersSvc,
	)
	server.InitRoutes()
	handler := cors.Default().Handler(server)
	//answersSvc.Start()
	panic(http.ListenAndServe(addr,handler))
}

