package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

// emp struct
type Employee struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Bg   string `json:"bg"`
}

func main() {

	//var e1 Employee := Employee{Id: 100, Name: "pushpa", Bg: "A+"};

	// inserting a rows
	insert(Employee{101, "pushpa1", "A+"})
	insert(Employee{102, "pushpa2", "B+"})

	// updating the emp by id
	updateById(Employee{101, "quest global", "A+"})

	// select all emp
	results := selectAll()

	// iterating a results
	for results.Next() {
		var e2 Employee
		results.Scan(&e2.Id, &e2.Name, &e2.Bg)
		fmt.Println(e2.Id, e2.Name, e2.Bg)
	}

	// select emp by id
	result := selectById(101)
	var e2 Employee
	result.Scan(&e2.Id, &e2.Name, &e2.Bg)
	fmt.Println(e2.Id, e2.Name, e2.Bg)

	// delete a emp by id
	delete(101)
}

// function to get a database connection
func connect() *sql.DB {
	db, err := sql.Open("mysql", "root:PP@gh.1995@tcp(127.0.0.1:3306)/empdb")
	if err != nil {
		fmt.Println("Error! Getting connection...")
	}
	return db
}

// function to insert a row in emp table
func insert(e2 Employee) {
	db := connect()
	insert, err := db.Query("INSERT INTO emp(id, name, bg) VALUES (?, ?, ?)", e2.Id, e2.Name, e2.Bg)
	if err != nil {
		fmt.Println("Error! Inserting records...")
	}
	defer insert.Close()
	defer db.Close()
}

// function to select all records from emp table
func selectAll() *sql.Rows {
	db := connect()
	results, err := db.Query("SELECT * FROM emp")
	if err != nil {
		fmt.Println("Error! Getting records...")
	}
	defer db.Close()
	return results
}

// function to select a emplyoee record from table by emp id
func selectById(id int) *sql.Row {
	db := connect()
	result := db.QueryRow("SELECT * FROM emp WHERE id=?", id)
	defer db.Close()
	return result
}

// function to update a emp record by emp id
func updateById(e2 Employee) {
	db := connect()
	db.QueryRow("UPDATE emp SET name=? WHERE id=?", e2.Name, e2.Id)
}

// function to delete a emp by emp id
func delete(id int) {
	db := connect()
	db.QueryRow("DELETE FROM emp WHERE id=?", id)
}
