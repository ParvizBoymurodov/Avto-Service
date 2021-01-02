package models

type Reg struct {
	Name     string `json:"name" validate:"nonzero"`
	Surname  string `json:"surname" validate:"nonzero"`
	Login    string `json:"login" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}

type Login struct {
	Login    string `json:"login" validate:"nonzero"`
	Password string `json:"password" validate:"nonzero"`
}

type ChangePass struct {
	Password string `json:"password" validate:"nonzero"`
}