package core

import (
	"fmt"
	"log"
	"os"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/joho/godotenv"
)

type (
	Application struct {
		Name    	string   `json:"name"`
		Port    	string   `json:"port"`
		Version 	string   `json:"version"`
		JWT_Secret  string `json:"jwt_secret"`
		Config  	Config   `json:"app_config"`
		DB      	*gorm.DB `json:"db"`
	}

	Config struct {
		Port        string
		DB_Host     string
		DB_User     string
		DB_Password string
		DB_Name     string
		DB_Log      string
		DB_Port 	string
	}
)

var App *Application

func init() {
	var err error
	App = &Application{}
	App.LoadAppConfig()
	
	if err = App.ConnectDatabase(); err != nil {
		log.Println("Error connecting to database: ", err)
	}
}

func (app *Application) LoadAppConfig() {
	_ = godotenv.Overload()

	app.Name = "Advisree-API"
	app.Version = os.Getenv("API_VERSION")
	app.Port = os.Getenv("API_PORT")

	app.Config = Config{
		Port:        os.Getenv("API_PORT"),
		DB_Host:     os.Getenv("DB_HOST"),
		DB_User:     os.Getenv("DB_USER"),
		DB_Password: os.Getenv("DB_PASS"),
		DB_Name:     os.Getenv("DB_NAME"),
		DB_Log:      os.Getenv("DB_LOG"),
		DB_Port:     os.Getenv("DB_PORT"),
	}
}

func (app *Application) ConnectDatabase() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local", app.Config.DB_User, app.Config.DB_Password, app.Config.DB_Host, app.Config.DB_Port, app.Config.DB_Name)

	db, err := gorm.Open("mysql", dsn)
	db.LogMode(app.Config.DB_Log == "1")
	app.DB = db

	return err
}

func (app *Application) Close() error {
	if err := app.DB.Close(); err != nil {
		return err
	}

	return nil
}