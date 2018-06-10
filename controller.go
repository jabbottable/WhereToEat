package wheretoeat

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

//Controller ...
type Controller struct {
	PlaceAPI PlaceAPI
}

// FindFood GET /
func (c *Controller) FindFood(w http.ResponseWriter, r *http.Request) {
	var location Location
	vars := mux.Vars(r)

	longitude, err := strconv.ParseFloat(vars["lon"], 64)

	if err != nil {
		log.Fatalln("Error FindFood Longitude", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	location.Longitude = longitude

	latitude, err := strconv.ParseFloat(vars["lat"], 64)

	if err != nil {
		log.Fatalln("Error FindFood Latitude", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	location.Latitude = latitude

	log.Println(location)
	placesSearchResponse := c.PlaceAPI.FindFood(location)
	data, _ := json.Marshal(placesSearchResponse)
	success := true
	if !success {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}
