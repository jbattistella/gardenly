package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jbattistella/gardenly/database"
	"github.com/jbattistella/gardenly/engine"
	"gorm.io/gorm"
)

type DataStore struct {
	db *gorm.DB
}

// gardenly ui request handlers
func getGardenMsgHandler(w http.ResponseWriter, r *http.Request) {

	type Reply struct {
		Messages       string
		PlantingSeason bool
		VegHeader      string
		Vegetables     []database.Vegetable
	}

	type ErrPage struct {
		Message error
	}

	params := mux.Vars(r)
	zipCode := params["zipcode"]

	fmt.Println(zipCode)

	res, err := engine.Engine(zipCode)
	if err != nil {
		er := ErrPage{Message: err}
		t, _ := template.ParseFiles("html/errpage.html")
		if err := t.Execute(w, er); err != nil {
			log.Fatal(err)
		}
		return
	}

	var rep Reply
	msg := res.Msg1 + res.Msg2 + res.Msg3

	if len(res.Vegetables) == 0 {
		rep = Reply{
			Messages:       msg,
			PlantingSeason: false,
		}
	} else {
		rep = Reply{
			Messages:       msg,
			PlantingSeason: true,
			VegHeader:      "You can seed the following:",
			Vegetables:     res.Vegetables,
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
func GardenlyHomeSubmission(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	zipcode := r.FormValue("zipcode")

	http.Redirect(w, r, fmt.Sprintf("/gardenly/%s", zipcode), http.StatusFound)

}

// gardenly database handlers
func (a *DataStore) getVegetables(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var vegs []database.Vegetable

	a.db.Find(&vegs)
	err := json.NewEncoder(w).Encode(&vegs)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func (a *DataStore) getVegetable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	var veg database.Vegetable
	a.db.First(&veg, "common_name = ?", params["name"])
	err := json.NewEncoder(w).Encode(&veg)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func (a *DataStore) createVegetable(w http.ResponseWriter, r *http.Request) {

	var veg database.Vegetable
	if err := json.NewDecoder(r.Body).Decode(&veg); err != nil {
		log.Printf("error decoding product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	a.db.Create(&veg)

	w.WriteHeader(http.StatusCreated)
}
func (a *DataStore) updateVegetable(w http.ResponseWriter, r *http.Request) {

	var veg database.Vegetable
	if err := json.NewDecoder(r.Body).Decode(&veg); err != nil {
		log.Printf("error decoding product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	a.db.Model(&veg).Where("common_name = ?", veg.CommonName).Update("dtm", veg.DTM)

	w.WriteHeader(http.StatusCreated)
}
func (a *DataStore) deleteVegetable(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	var veg database.Vegetable
	if err := json.NewDecoder(r.Body).Decode(&veg); err != nil {
		log.Printf("error decoding product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	a.db.Where("common_name = ?", params["name"]).Delete(&veg)

	w.WriteHeader(http.StatusOK)
}
func GardenAPI() {

	DB, err := database.ConnectDB()
	if err != nil {
		log.Fatal()
	}

	a := DataStore{db: DB}

	r := mux.NewRouter()

	r.HandleFunc("/gardenly/{zipcode}", getGardenMsgHandler).Methods("GET")
	r.HandleFunc("/", GardenlyHome).Methods("GET")
	r.HandleFunc("/", GardenlyHomeSubmission).Methods("POST")

	//database
	r.HandleFunc("/vegetables", a.getVegetables).Methods("GET")
	r.HandleFunc("/vegetables/{name}", a.getVegetable).Methods("GET")
	r.HandleFunc("/vegetables", a.createVegetable).Methods("POST")
	r.HandleFunc("/vegetables", a.updateVegetable).Methods("PUT")
	r.HandleFunc("/vegetables/{name}", a.deleteVegetable).Methods("DELETE")

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
