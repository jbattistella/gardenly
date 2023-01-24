package server

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/jbattistella/capstone-project/database"
	"github.com/jbattistella/capstone-project/engine"
)

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
func getVegetables(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	DB, err := database.ConnectDB()
	if err != nil {
		log.Fatal("error connecting to DB")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	var vegs []database.Vegetable
	DB.Find(&vegs)
	err = json.NewEncoder(w).Encode(&vegs)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func getVegetable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)

	DB, err := database.ConnectDB()
	if err != nil {
		log.Fatal("error connecting to DB")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var veg database.Vegetable
	DB.First(&veg, "common_name = ?", params["name"])
	err = json.NewEncoder(w).Encode(&veg)
	if err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
func createVegetable(w http.ResponseWriter, r *http.Request) {

	DB, err := database.ConnectDB()
	if err != nil {
		log.Fatal("error connecting to DB")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var veg database.Vegetable
	if err := json.NewDecoder(r.Body).Decode(&veg); err != nil {
		log.Printf("error decoding product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
	}

	DB.Create(&veg)

	w.WriteHeader(http.StatusCreated)
}
func updateVegetable(w http.ResponseWriter, r *http.Request) {
	DB, err := database.ConnectDB()
	if err != nil {
		log.Fatal("error connecting to DB")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var veg database.Vegetable
	if err := json.NewDecoder(r.Body).Decode(&veg); err != nil {
		log.Printf("error decoding product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	DB.Model(&veg).Where("common_name = ?", veg.CommonName).Update("dtm", veg.DTM)

	w.WriteHeader(http.StatusCreated)
}
func deleteVegetable(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	DB, err := database.ConnectDB()
	if err != nil {
		log.Fatal("error connecting to DB")
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	var veg database.Vegetable
	if err := json.NewDecoder(r.Body).Decode(&veg); err != nil {
		log.Printf("error decoding product: %v", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	DB.Where("common_name = ?", params["name"]).Delete(&veg)

	w.WriteHeader(http.StatusOK)
}
func GardenAPI() {

	r := mux.NewRouter()

	r.HandleFunc("/gardenly/{zipcode}", getGardenMsgHandler).Methods("GET")
	r.HandleFunc("/", GardenlyHome).Methods("GET")
	r.HandleFunc("/", GardenlyHomeSubmission).Methods("POST")

	//database
	r.HandleFunc("/vegetables", getVegetables).Methods("GET")
	r.HandleFunc("/vegetables/{name}", getVegetable).Methods("GET")
	r.HandleFunc("/vegetables", createVegetable).Methods("POST")
	r.HandleFunc("/vegetables", updateVegetable).Methods("PUT")
	r.HandleFunc("/vegetables/{name}", deleteVegetable).Methods("DELETE")

	log.Println("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
