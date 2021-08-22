package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

//database code
const (
	uri = "postgres://vuuqbghzjegyvk:5d0dc4f72d109d2bb5d2c652fe743db2570dcac033f0cb16b83c55464484da05@ec2-52-72-125-94.compute-1.amazonaws.com:5432/d8r669a0v2gin6"
	host= "ec2-52-72-125-94.compute-1.amazonaws.com"
	user   = "vuuqbghzjegyvk"
	password = "5d0dc4f72d109d2bb5d2c652fe743db2570dcac033f0cb16b83c55464484da05"
	dbname   = "d8r669a0v2gin6"
)

var tpl *template.Template
var tpl2 *template.Template

//parse index file with form
func init() {
	tpl = template.Must(template.ParseFiles("assets/index.gohtml"))
	tpl2 = template.Must(template.ParseFiles("assets/submit.gohtml"))
	//handle func called in main func with r.handlefunc this only a failsafe
	//http.HandleFunc("/save", submit)
}

type Visitor struct {
	gorm.Model
	Name    string
	Email   string
	Message string
}

func index(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := tpl.Execute(w, nil); err != nil {
		panic(err)
	}

}

func submit(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	submittedName := r.FormValue("yname")
	submittedEmail := r.FormValue("ymail")
	submittedMessage := r.FormValue("ymessage")

	//create visitors details struct
	data := &Visitor{

		Name:    submittedName,
		Email:   submittedEmail,
		Message: submittedMessage}

	//connect to database
	psqlconn := fmt.Sprintf( "uri=%s host=%s user=%s password=%s dbname=%s sslmode=require", uri, host, user, password, dbname)

	//create record
	db, err := gorm.Open(postgres.Open(psqlconn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	result := db.Create(&data)
	if err := tpl2.Execute(w, result); err != nil {
		panic(err)
	}

}

func main() {

	//initialize routers
	r := mux.NewRouter()
	r.HandleFunc("/", index).Methods("GET")
	r.HandleFunc("/submit", submit).Methods("POST")

	//serve static files
	r.PathPrefix("/").Handler(http.FileServer(http.Dir("./assets/")))

	//for local host uncomment below line
	//log.Fatal(http.ListenAndServe(":8080", r))
	//for web service
	log.Fatal(http.ListenAndServe(":"+os.Getenv("PORT"), r))

}
