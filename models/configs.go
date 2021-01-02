package models

type Config struct {
	Host string `json:"host"`
	Port string `json:"port"`
	Dsn  string `json:"dsn"`
	Key  string `json:"key"`
}
