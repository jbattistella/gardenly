package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jbattistella/capstone-project/database"
	"github.com/jbattistella/capstone-project/engine"
)

func getGardenMsgHandler(w http.ResponseWriter, r *http.Request) {

	type Reply struct {
		Messages   string
		Vegetables []database.Vegetable
	}

	type ErrPage struct {
		Message error
	}

	params := mux.Vars(r)
	UserId := params["zipcode"]

	res, err := engine.Engine(UserId)
	if err != nil {
		er := ErrPage{Message: err}
		t, _ := template.ParseFiles("errpage.html")
		if err := t.Execute(w, er); err != nil {
			log.Fatal(err)
		}
		return
	}

	msg := res.Msg1 + res.Msg2 + res.Msg3

	rep := Reply{
		Messages:   msg,
		Vegetables: res.Vegetables,
	}

	t, _ := template.ParseFiles("gardenly.html")
	if err := t.Execute(w, rep); err != nil {
		log.Fatal(err)
	}
}

func main() {

	r := mux.NewRouter()

	r.HandleFunc("/gardenly/{zipcode}", getGardenMsgHandler)

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
