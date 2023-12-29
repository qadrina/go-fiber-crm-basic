package main

import (
	"fmt"

	"github.com/gofiber/fiber"
	"github.com/qadrina/go-fiber-crm-basic/database"
	"github.com/qadrina/go-fiber-crm-basic/lead"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
)

func setupRoutes(app *fiber.App) {
	app.Get("api/v1/lead", lead.GetLeads)
	app.Get("api/v1/lead/:id", lead.GetLead)
	app.Post("api/v1/lead", lead.NewLead)
	app.Delete("api/v1/lead/:id", lead.DeleteLead)
}

func initDatabase() {
	// before going to this part, import gorm in the database go file.
	var err error
	dsn := "sqlserver://uae:Demo1234@localhost:1434?database=TestDev"
	database.DBConn, err = gorm.Open(sqlserver.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect to database.")
	}
	fmt.Println("Connection opened to database.")
	database.DBConn.AutoMigrate(&lead.Lead{})
	fmt.Println("Database migrated.")
}

func main() {
	// call the fiber library
	app := fiber.New()
	// connect to the DB
	initDatabase()
	setupRoutes(app)
	// listen to port 3000
	app.Listen(3000)
	// start connection to DB
	//defer database.DBConn.Close()
}
