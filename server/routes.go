package employee

import (
	"encoding/json"
	"net/http"
)

// Route struct to hold route data
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes - Slice of Route
type Routes []Route

// EmployeeRoutes defines the employee routes/endpoints
var EmployeeRoutes = Routes{
	Route{
		Name:        "GetAllEmployees",
		Method:      "GET",
		Pattern:     "/manage/employee",
		HandlerFunc: GetAllEmployeesHandler, // curl -sX GET http:localhost:8080/manage/employee | jq
	},
	Route{
		Name:        "GetEmployeeByNameORId",
		Method:      "GET",
		Pattern:     "/manage/employee/{idORname}",
		HandlerFunc: GetEmployeeHandler, // curl -sX GET http:localhost:8080/manage/employee/{name/id} | jq
	}, Route{
		Name:        "GetEmployees",
		Method:      "GET",
		Pattern:     "/manage/employees/{ids}",
		HandlerFunc: GetEmployeesHandler, //curl -sX GET http:localhost:8080/manage/employee/{ids} | jq
	}, Route{
		Name:        "AddEmployee",
		Method:      "POST",
		Pattern:     "/manage/employee",
		HandlerFunc: AddEmployeeHandler, // Ex: curl -sX POST http://localhost:8080/manage/employee -d '{"name": "Bob","age":32,"gender":"M"}' | jq
	}, Route{
		Name:        "UpdateEmployee",
		Method:      "PUT",
		Pattern:     "/manage/employee/{id}",
		HandlerFunc: UpdateEmployeeHandler, // Ex: curl -sX PUT http://localhost:8080/manage/employee/{id} -d '{"name": "Bob2","age":34}' | jq
	}, Route{
		Name:        "DeleteEmployee",
		Method:      "DELETE",
		Pattern:     "/manage/employee/{id}",
		HandlerFunc: DeleteEmployeeHandler, // Ex: curl -sX DELETE http://localhost:8080/manage/employee/{id} | jq
	},
}

func webJSONResponse(w http.ResponseWriter, status int, payload interface{}) {
	response, err := json.Marshal(payload)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(status)
	w.Write([]byte(response))
}

func webErrorResponse(w http.ResponseWriter, code int, message string) {
	webJSONResponse(w, code, map[string]string{"error": message})
}
