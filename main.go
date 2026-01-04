package main

import (
	"fmt"
	"workshop2/database"
	"workshop2/routes"

	"github.com/gofiber/fiber/v2"
    m "workshop2/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func initDatabase() {
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=Local",
		"root",
		"",
		"127.0.0.1",
		"3306",
		"accout",
	)
	var err error
	database.DBConn, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	fmt.Println("Database connected!")
	database.DBConn.AutoMigrate(&m.Register{})
}

func main() {
	app := fiber.New()
	routes.Rountes(app)
	app.Listen(":3000")
}
