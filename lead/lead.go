package lead

import (
	//"gorm.io/driver/sqlserver"
	"github.com/gofiber/fiber"
	"github.com/qadrina/go-fiber-crm-basic/database"
	"gorm.io/gorm"
)

type Lead struct {
	gorm.Model
	// the back ticks is used to tell golang what it's gonna look like in json
	Name    string `json:"name"`
	Company string `json:"company"`
	Email   string `json:"email"`
	Phone   int    `json:"phone"`
}

func GetLeads(c *fiber.Ctx) {
	// ctx = context
	// the function can start work with the data that's coming from the user
	db := database.DBConn
	// slice of Lead means an array of objects
	var leads []Lead
	db.Find(&leads)
	c.JSON(leads)
}

func GetLead(c *fiber.Ctx) {
	// fetch the data of the particular ID
	id := c.Params("id")
	// establish the db connection
	db := database.DBConn
	// similar to Lead lead in C#, Lead is the model, lead is the var name
	var lead Lead
	db.Find(&lead, id)
	// send the response in json format
	c.JSON(lead)
}

func NewLead(c *fiber.Ctx) {
	db := database.DBConn
	lead := new(Lead)
	if err := c.BodyParser(lead); err != nil {
		c.Status(503).Send(err)
		return
	}
	db.Create(&lead)
	c.JSON(lead)
}

func DeleteLead(c *fiber.Ctx) {
	id := c.Params("id")
	db := database.DBConn
	var lead Lead
	db.First(&lead, id)
	if lead.Name == "" {
		c.Status(500).Send("No lead found with ID")
	}
	db.Delete(&lead)
	c.Send("Lead successfully deleted")
}
