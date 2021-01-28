// 🚀 Fiber is an Express inspired web framework written in Go with 💖
// 📌 API Documentation: https://fiber.wiki
// 📝 Github Repository: https://github.com/gofiber/fiber
package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gofiber/fiber/v2"
	"github.com/mysqlfibergo/employee"
)

// Database instance
var db *sql.DB

// Database settings
const (
	host     = "localhost"
	port     = 3306 // Default port
	user     = "root"
	password = "newpass"
	dbname   = "fiber_demo"
)

// Employee struct

// Employees struct

// Connect function
func Connect() error {
	var err error
	// Use DSN string to open
	db, err = sql.Open("mysql", fmt.Sprintf("%s:%s@/%s", user, password, dbname))
	if err != nil {
		return err
	}
	if err = db.Ping(); err != nil {
		return err
	}
	return nil
}

func setupRoutes(app *fiber.App) {
	app.Get("/api/v1/employee", employee.GetAllEmployee)
	app.Get("/api/v1/employee/:id", employee.GetEmployee)
	app.Post("/api/v1/employee", employee.AddEmployee)
	app.Delete("/api/v1/employee/:id", employee.DeleteEmployee)
}

func main() {
	// Connect with database
	if err := Connect(); err != nil {
		log.Fatal(err)
	}

	// Create a Fiber app
	app := fiber.New()


	log.Fatal(app.Listen(":3000"))
}
