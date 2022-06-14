package handler

import (
	"customer/model"
	"database/sql"
	"log"

	_ "github.com/denisenkom/go-mssqldb"
	// _ "github.com/samonzeweb/godb"
	// _ "github.com/samonzeweb/godb/adapters/mssql"
)

func Init_DB() (*sql.DB, error) {
	configModel := model.SystemConfig{}
	configModel.LoadConfig()
	return sql.Open("sqlserver", configModel.ConnectString)
}

func Close_DB(dbc *sql.DB) error {
	err := dbc.Close()
	if err != nil {
		log.Fatal("Error While load Config : ", err.Error())
		return err
	}
	return nil
}
