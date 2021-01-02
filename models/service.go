package models

type Svc struct {
	Name   string `json:"name" validate:"nonzero"`
	Region string `json:"region" validate:"nonzero"`
}

type SelectSvc struct {
	Id     int64  `json:"id"`
	Name   string `json:"name"`
	Region string `json:"region"`
}
