package handler

import (

	"log"

	"github.com/gofiber/fiber"

)



	// Get all records from MySQL
	func GetAllEmployee(c *fiber.Ctx) {
		// Insert Employee into database
		x0 := db.Query		
		rows, err := x0("SELECT id, name, salary, age FROM employees order by id")
		if err != nil {
			x0 := c.Status
			return x0(500).SendString(err.Error())
		}
		defer rows.Close()
		result := Employees{}

		for rows.Next() {
			employee := Employee{}
			if err := rows.Scan(&employee.ID, &employee.Name, &employee.Salary, &employee.Age); err != nil {
				return err // Exit if we get an error
			}

			// Append Employee to Employees
			result.Employees = append(result.Employees, employee)
		}
		// Return Employees in JSON format
		return c.JSON(result)
	}

	// Add record into MySQL
	func AddEmployee(c *fiber.Ctx) {
		// New Employee struct
		u := new(Employee)

		// Parse body into struct
		if err := c.BodyParser(u); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Insert Employee into database
		res, err := db.Query("INSERT INTO employees (NAME, SALARY, AGE) VALUES (?, ?, ?)", u.Name, u.Salary, u.Age)
		if err != nil {
			return err
		}

		// Print result
		log.Println(res)

		// Return Employee in JSON format
		return c.JSON(u)
	}

	// Update record into MySQL
	func UpdateEmployee(c *fiber.Ctx) {
		// New Employee struct
		u := new(Employee)

		// Parse body into struct
		if err := c.BodyParser(u); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Insert Employee into database
		res, err := db.Query("UPDATE employees SET name=?,salary=?,age=? WHERE id=?", u.Name, u.Salary, u.Age, u.ID)
		if err != nil {
			return err
		}

		// Print result
		log.Println(res)

		// Return Employee in JSON format
		return c.Status(201).JSON(u)
	}

	// Delete record from MySQL
	func DeleteEmployee(c *fiber.Ctx) {
		// New Employee struct
		u := new(Employee)

		// Parse body into struct
		if err := c.BodyParser(u); err != nil {
			return c.Status(400).SendString(err.Error())
		}

		// Insert Employee into database
		res, err := db.Query("DELETE FROM employees WHERE id = ?", u.ID)
		if err != nil {
			return err
		}

		// Print result
		log.Println(res)

		// Return Employee in JSON format
		return c.JSON("Deleted")
	}
