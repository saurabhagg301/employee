package employee

import (
	"database/sql"
	"errors"
	"strconv"
	"strings"
)

// Store struct will implement the MgtCfgStorer interface
type Store struct {
	Db *sql.DB
}

// EmpStorer interface to hold methods signatures
type EmpStorer interface {
	GetAllEmployees() ([]Employee, error)
	GetEmployeeByName(string) (Employee, error)
	GetEmployeeByID(int) (Employee, error)
	GetEmployeesByIDs(string) ([]Employee, error)
	AddEmployee(Employee) (int, error)
	UpdateEmployee(Employee) error
	DeleteEmployeeByID(int) error
}

// EmpStore - an instance of MgtCfgStorer
var EmpStore EmpStorer

// InitStore ...
func InitStore(e EmpStorer) {
	EmpStore = e
}

// GetAllEmployees to fetch all employees details
func (store *Store) GetAllEmployees() ([]Employee, error) {
	res := []Employee{}
	sql := "SELECT id, name, age, gender FROM emp.employee ORDER BY id"
	rows, err := store.Db.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		e := Employee{}
		if err := rows.Scan(&e.ID, &e.Name, &e.Age, &e.Gender); err != nil {
			return nil, err
		}
		res = append(res, e)
	}
	if len(res) == 0 {
		return res, errors.New("no records found")
	}

	return res, nil

}

// GetEmployeeByName to fetch a particular employee details by name
func (store *Store) GetEmployeeByName(name string) (Employee, error) {
	e := Employee{}
	sql := "SELECT id, name, age, gender FROM emp.employee WHERE lower(name) = $1"
	row, err := store.Db.Query(sql, strings.ToLower(name))
	if err != nil {
		return e, err
	}
	defer row.Close()

	for row.Next() {
		if err := row.Scan(&e.ID, &e.Name, &e.Age, &e.Gender); err != nil {
			return e, err
		}
	}

	// if no record found, return no record found
	if e == (Employee{}) {
		return e, errors.New("no record found for name = " + name)
	}
	return e, nil

}

// GetEmployeeByID to fetch a particular employee details by ID
func (store *Store) GetEmployeeByID(id int) (Employee, error) {
	e := Employee{}
	sql := "SELECT id, name, age, gender FROM emp.employee WHERE id = $1"
	row, err := store.Db.Query(sql, id)
	if err != nil {
		return e, err
	}
	defer row.Close()

	for row.Next() {
		if err := row.Scan(&e.ID, &e.Name, &e.Age, &e.Gender); err != nil {
			return e, err
		}
	}

	// if no record found, return no record found
	if e == (Employee{}) {
		return e, errors.New("no record found for Id=" + strconv.Itoa(id))
	}
	return e, nil

}

// GetEmployeesByIDs to fetch details of multiple employees
func (store *Store) GetEmployeesByIDs(ids string) ([]Employee, error) {
	res := []Employee{}
	idArr := strings.Split(ids, ",")
	sql := "SELECT id, name, age, gender FROM emp.employee WHERE id in "
	for k, v := range idArr {
		if k == 0 {
			sql += "(" + v
		} else {
			sql += ", " + v

		}
	}
	sql += ")"
	row, err := store.Db.Query(sql)
	if err != nil {
		return res, err
	}
	defer row.Close()
	for row.Next() {
		e := Employee{}
		if err := row.Scan(&e.ID, &e.Name, &e.Age, &e.Gender); err != nil {
			return res, err
		}
		res = append(res, e)

	}

	// if no record found, return no record found
	if len(res) == 0 {
		return res, errors.New("no record found for Id's = " + ids)
	}
	return res, nil

}

// AddEmployee to add a nemployee
func (store *Store) AddEmployee(e Employee) (int, error) {
	var id int64
	sql := "INSERT INTO emp.employee(name, age, gender) values ($1,$2,$3) RETURNING id"
	err := store.Db.QueryRow(sql, e.Name, e.Age, e.Gender).Scan(&id)
	if err != nil {
		return 0, err
	}
	return int(id), nil

}

// UpdateEmployee to update an employee
func (store *Store) UpdateEmployee(e Employee) error {

	sql := "UPDATE emp.employee SET name= $1, age= $2, gender= $3 WHERE id = $4"

	_, err := store.Db.Exec(sql, e.Name, e.Age, e.Gender, e.ID)
	if err != nil {
		return err
	}
	return nil
}

// DeleteEmployeeByID to delete an employee
func (store *Store) DeleteEmployeeByID(id int) error {

	sql := "DELETE FROM emp.employee WHERE id = $1"

	_, err := store.Db.Exec(sql, id)
	if err != nil {
		return err
	}
	return nil
}
