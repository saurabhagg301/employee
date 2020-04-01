package employee

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	// Postgres native driver
	_ "github.com/lib/pq"
)

var serverStartTime = time.Now()

// Message struct to hold uptime heartbeat
type Message struct {
	Uptime float32 `json:"uptime"`
}

const (
	dbHost     = "localhost"
	dbPort     = 5432
	dbUser     = "postgres"
	dbPassword = "postgres"
	dbName     = "employee_db"
)

// StartWebServer to start webserver
func StartWebServer(host string, port int) {
	// start database server
	psqlConnStr := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName)

	db, err := sql.Open("postgres", psqlConnStr)
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err.Error())
		panic(err)
	}

	InitStore(&Store{Db: db})

	// Get mux router
	r := NewRouter(EmployeeRoutes)

	// Configure server properties
	srvr := &http.Server{
		Addr:         fmt.Sprintf("%s:%d", host, port),
		ReadTimeout:  15 * time.Second,
		WriteTimeout: 15 * time.Second,
		Handler:      r,
	}

	// Ping is handler for checking heartbeat
	r.HandleFunc("/ping", func(w http.ResponseWriter, r *http.Request) {
		t2 := time.Now()
		uptime := float32(t2.Sub(serverStartTime).Seconds())
		m := Message{Uptime: uptime}
		webJSONResponse(w, http.StatusOK, m)
	}).Methods("GET")
	fmt.Println("Listening at port ", port)
	// start webserver
	log.Fatal(srvr.ListenAndServe())
}

// NewRouter to get an instance of gorilla mux router
func NewRouter(routes Routes) *mux.Router {

	// Create an instance of the Gorilla router
	router := mux.NewRouter()
	var handler http.Handler

	for _, route := range routes {

		// Attach each route
		handler = route.HandlerFunc
		router.Name(route.Name).
			Methods(route.Method).
			Path(route.Pattern).
			Handler(handler)
	}
	return router
}
