package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jbattistella/capstone-project/engine"
)

func getGardenMsgHandler(w http.ResponseWriter, r *http.Request) {

	type Reply struct {
		Messages  string
		VegHeader string
		V1        string
		V2        string
		V3        string
		V4        string
		V5        string
		V6        string
		V7        string
		V8        string
		V9        string
		V10       string
		V11       string
		V12       string
		V13       string
		V14       string
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

	var vegNames []string
	for _, v := range res.Vegetables {

		vegNames = append(vegNames, v.CommonName)
	}

	var rep Reply

	if len(vegNames) == 0 {
		rep = Reply{
			Messages:  msg,
			VegHeader: "You can seed the following:",
		}
	} else {
		rep = Reply{
			Messages:  msg,
			VegHeader: "You can seed the following:",
			V1:        vegNames[0],
			V2:        vegNames[1],
			V3:        vegNames[2],
			V4:        vegNames[3],
			V5:        vegNames[4],
			V6:        vegNames[5],
			V7:        vegNames[6],
			V8:        vegNames[7],
			V9:        vegNames[8],
			V10:       vegNames[9],
			V11:       vegNames[10],
			V12:       vegNames[11],
			V13:       vegNames[12],
			V14:       vegNames[13],
		}
	}

	t, _ := template.ParseFiles("gardenly.html")
	if err := t.Execute(w, rep); err != nil {
		log.Fatal(err)
	}
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

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
