package models

type Market struct {
	Id        int64  `json:"id"`
	Name      string `json:"name" validate:"nonzero"`
	ServiceId int64  `json:"service_id" validate:"nonzero"`
	Title     string `json:"title" validate:"nonzero"`
	Service   Svc    `json:"service"`
}

