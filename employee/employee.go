package employee

import (
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

type Employee struct {
	ID     int     `json:"id"`
	Name   string  `json:"name"`
	Salary float64 `json:"salary"`
	Age    int     `json:"age"`
}

type Employees struct {
	Employees []Employee `json:"employees"`
}