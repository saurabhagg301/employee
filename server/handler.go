package employee

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

// Employee to hold employee data
type Employee struct {
	ID     int    `json:"id"`
	Name   string `json:"name"`
	Age    int    `json:"age"`
	Gender string `json:"gender"`
}

/* Test Data
var Employees = []Employee{
	{ID: 1, Name: "Ramesh", Age: 30, Gender: "M"},
	{ID: 2, Name: "Raj", Age: 34, Gender: "M"},
	{ID: 3, Name: "Bob", Age: 28, Gender: "M"},
}

*/

// GetAllEmployeesHandler to get details of all employees
// Ex : curl -sX GET http:localhost:8080/manage/employee | jq
func GetAllEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	res, err := EmpStore.GetAllEmployees()
	if err != nil {
		webErrorResponse(w, http.StatusBadRequest, err.Error())
	}
	webJSONResponse(w, http.StatusOK, res)
}

// GetEmployeeHandler to get details of a specific employees using id or name
//Ex : curl -sX GET http:localhost:8080/manage/employee/{name} | jq
//Ex : curl -sX GET http:localhost:8080/manage/employee/{id} | jq
func GetEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	var res Employee
	idORname := vars["idORname"]
	id, errParseInt := strconv.ParseInt(idORname, 10, 0)
	if errParseInt != nil {
		name := idORname
		res, err = EmpStore.GetEmployeeByName(name)

	} else {
		res, err = EmpStore.GetEmployeeByID(int(id))

	}
	if err != nil {
		webErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	webJSONResponse(w, http.StatusOK, res)
}

// GetEmployeesHandler to get details of multiple employees using id's
//curl -sX GET http:localhost:8080/manage/employee/1,2,4 | jq
func GetEmployeesHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var err error
	var res []Employee
	ids := vars["ids"]

	res, err = EmpStore.GetEmployeesByIDs(ids)

	if err != nil {
		webErrorResponse(w, http.StatusBadRequest, err.Error())
		return
	}
	webJSONResponse(w, http.StatusOK, res)
}

// AddEmployeeHandler to add an employee
// Ex: curl -sX POST http://localhost:8080/manage/employee -d '{"name": "Bob","age":32,"gender":"M"}'
func AddEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	e := Employee{}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&e); err != nil {
		webErrorResponse(w, http.StatusBadRequest, "invalid json")
		return
	}
	id, errAddEmploye := EmpStore.AddEmployee(e)
	if errAddEmploye != nil {
		webErrorResponse(w, http.StatusBadRequest, errAddEmploye.Error())
		return
	}
	webJSONResponse(w, http.StatusCreated, "msg: added employee "+e.Name+" with ID "+strconv.Itoa(id))
}

// UpdateEmployeeHandler to update an employee
// Ex: curl -sX PUT http://localhost:8080/manage/employee/{id} -d '{"name": "Bob2","age":34}'
func UpdateEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	var empNew Employee
	id, errParseInt := strconv.ParseInt(idStr, 10, 8)
	if errParseInt != nil {
		webErrorResponse(w, http.StatusBadRequest, "invalid id")
		return
	}
	empOld, errFetchEmpDetails := EmpStore.GetEmployeeByID(int(id))
	if errFetchEmpDetails != nil {
		webErrorResponse(w, http.StatusBadRequest, errFetchEmpDetails.Error())
		return
	}
	decoder := json.NewDecoder(r.Body)
	if err := decoder.Decode(&empNew); err != nil {
		webErrorResponse(w, http.StatusBadRequest, "invalid json")
		return
	}
	empNew.ID = int(id)

	if empNew.Name == "" {
		empNew.Name = empOld.Name
	}
	if empNew.Age == 0 {
		empNew.Age = empOld.Age
	}
	if empNew.Gender == "" {
		empNew.Gender = empOld.Gender
	}
	errUpdateEmployee := EmpStore.UpdateEmployee(empNew)
	if errUpdateEmployee != nil {
		webErrorResponse(w, http.StatusBadRequest, errUpdateEmployee.Error())
		return
	}
	webJSONResponse(w, http.StatusOK, "msg: employee id "+strconv.Itoa(empNew.ID)+" updated successfully")
}

// DeleteEmployeeHandler to delete an employee
// Ex: curl -sX DELETE http://localhost:8080/manage/employee/{id}
func DeleteEmployeeHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	idStr := vars["id"]
	id, errParseInt := strconv.ParseInt(idStr, 10, 8)
	if errParseInt != nil {
		webErrorResponse(w, http.StatusBadRequest, "invalid id")
		return
	}
	_, errFetchEmpDetails := EmpStore.GetEmployeeByID(int(id))
	if errFetchEmpDetails != nil {
		webErrorResponse(w, http.StatusBadRequest, errFetchEmpDetails.Error())
		return
	}

	errDeleteEmployee := EmpStore.DeleteEmployeeByID(int(id))
	if errDeleteEmployee != nil {
		webErrorResponse(w, http.StatusBadRequest, errDeleteEmployee.Error())
		return
	}
	webJSONResponse(w, http.StatusOK, "msg: employee id "+strconv.Itoa(int(id))+" deleted successfully")
}
