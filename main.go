package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jbattistella/capstone-project/engine"
)

func getGardenMsgHandler(w http.ResponseWriter, r *http.Request) {

	type Reply struct {
		Messages   string
		VegHeader  string
		Vegetables []string
	}

	type ErrPage struct {
		Message error
	}

	params := mux.Vars(r)
	UserId := params["zipcode"]

	res, err := engine.Engine(UserId)
	if err != nil {
		er := ErrPage{Message: err}
		t, _ := template.ParseFiles("htmlPages/errpage.html")
		if err := t.Execute(w, er); err != nil {
			log.Fatal(err)
		}
		return
	}

	var rep Reply

	msg := res.Msg1 + res.Msg2 + res.Msg3

	var vegNames []string
	for _, v := range res.Vegetables {

		vegNames = append(vegNames, v.CommonName)
	}

	if len(vegNames) == 0 {
		rep = Reply{
			Messages: msg,
		}
	} else {
		rep = Reply{
			Messages:   msg,
			VegHeader:  "You can seed the following:",
			Vegetables: vegNames,
		}
	}

	t, _ := template.ParseFiles("html/gardenly.html")
	if err := t.Execute(w, rep); err != nil {
		log.Fatal(err)
	}
}

func GardenlyHome(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "html/homepage.html")
}

func getDemoGardenMsgHandler(w http.ResponseWriter, r *http.Request) {

	type Reply struct {
		Messages   string
		VegHeader  string
		Vegetables []string
	}

	type ErrPage struct {
		Message error
	}

	rep := Reply{
		Messages:   "There are 60 days until the first frost",
		Vegetables: []string{"carrot", "beats", "cilantro", "parsley"},
	}

	t, _ := template.ParseFiles("html/gardenly.html")
	if err := t.Execute(w, &rep); err != nil {
		log.Fatal(err)
	}
}
func GardenlyHomeSubmission(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	zipcode := r.FormValue("zipcode")

	http.Redirect(w, r, fmt.Sprintf("/gardenly/%s", zipcode), http.StatusFound)

}

func notZero(s string) string {
	if len(s) == 0 {
		return ""
	}
	return s
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/gardenly/{zipcode}", getGardenMsgHandler)
	r.HandleFunc("/test/", getDemoGardenMsgHandler)
	r.HandleFunc("/", GardenlyHome).Methods("GET")
	r.HandleFunc("/", GardenlyHomeSubmission).Methods("POST")

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
