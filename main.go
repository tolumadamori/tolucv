package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//database code
const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "..toluwani"
	dbname   = "visitors_db"
)

var tpl *template.Template
var tpl2 *template.Template

//parse index file with form
func init() {
	tpl = template.Must(template.ParseFiles("assets/index.gohtml"))
	tpl2 = template.Must(template.ParseFiles("assets/submit.gohtml"))
}

type datatoinput struct {
	Name    string
	Email   string
	Message string
}

func index(w http.ResponseWriter, r *http.Request) {
	submittedName := r.FormValue("yname")
	submittedEmail := r.FormValue("yemail")
	submittedMessage := r.FormValue("ymessage")
	d := datatoinput{

		Name:    submittedName,
		Email:   submittedEmail,
		Message: submittedMessage,
	}
	w.Header().Set("Content-Type", "text/html")
	if err := tpl.Execute(w, d); err != nil {
		panic(err)
	}

}

func submit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tpl2.Execute(w, nil); err != nil {
		panic(err)
	}

}

func main() {

	// //connect to database
	// psqlconn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// db, err := sql.Open("postgres", psqlconn)
	// if err != nil {
	//     panic(err.Error())
	// defer db.Close()

	// insertStmt := `insert into "Visitors"("Name", "email", "message") values(submittedName, submittedEmail, submittedMessage)`
	// _, e := db.Exec(insertStmt)
	// CheckError(e)

	// insertDynStmt := `insert into "Visitors"("Name", "email", "message") values($1, $2,$3, )`
	// _, a := db.Exec(insertDynStmt, "toba", "tobamadamori", "bro")
	// CheckError(a)
	// insert, err := db.Query("INSERT INTO Visitors VALUES (submittedName, submittedEmail, submittedMessage )")
	// if err != nil {
	//     panic(err.Error())
	//initialize routers
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/submit", submit).Methods("POST")

	//serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))

	log.Fatal(http.ListenAndServe(":8080", r))

}
