package model

import (
	"log"

	"github.com/spf13/viper"
)

//model customer

type CustomerModel struct {
	Id        int    `db:"Id" json:"Id"`
	FirstName string `db:"FirstName" json:"firstName"`
	LastName  string `db:"LastName" json:"lastName"`
	Address   string `db:"Address" json:"address"`
	Birthday  int    `db:"Birthday" json:"birthday"`
}

//controller

type Controller struct {
	Message string      `json:"message" example:"success"`
	Data    interface{} `json:"data"`
}

type QueryEffected struct {
	Row int64 `json:"row" example:"1"`
}

//system config

type SystemConfig struct {
	ConnectString string
	Host          string
	Protocal      string
}

func (c *SystemConfig) LoadConfig() {
	viper.SetConfigFile(".env")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatalf("Error while reading config file %s", err)
	}
	c.ConnectString = viper.GetString("CONNECTION_STRING")
	c.Host = viper.GetString("HOST")
	c.Protocal = viper.GetString("PROTOCAL")
}
