package route

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

// BaseURL is the base end-point for star wars api
const BaseURL = "https://swapi.dev/api/"

// Planet : details about the planet
type Planet struct {
	Name       string `json:"name"`
	Population string `json:"population"`
	Terrain    string `json:"terrain"`
}

// Person : acting in the star wars
type Person struct {
	Name         string `json:"name"`
	HomeworldURL string `json:"homeworld"`
	Homeworld    Planet `json:""`
}

// AllPeople : list of persons acting in the movie
// because we have "results" in out json that we want to use as "People"
type AllPeople struct {
	People []Person `json:"results"`
}

// GetPeople returns list of people in star war franchise and their details
func GetPeople() func(w http.ResponseWriter, r *http.Request) {

	getPeople := func(w http.ResponseWriter, r *http.Request) {
		// fmt.Fprint(w, "getting people")

		res, err := http.Get(BaseURL + "people")

		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("Failed to request star wars people")
		}

		// fmt.Println(res)

		var bytes []byte
		if bytes, err = ioutil.ReadAll(res.Body); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			log.Println("Failed to parse response body")
		}

		// fmt.Println(bytes)
		// fmt.Println(string(bytes))

		var people AllPeople
		if err := json.Unmarshal(bytes, &people); err != nil {
			log.Println("Error parsing json")
		}

		// fmt.Println(people)

		for _, pers := range people.People {
			pers.getHomeWorld()
			fmt.Println(pers)
		}
	}

	return getPeople
}

func (p *Person) getHomeWorld() {
	res, err := http.Get(p.HomeworldURL)

	if err != nil {
		log.Println("Error fetching homeworld", err)
	}

	var bytes []byte
	if bytes, err = ioutil.ReadAll(res.Body); err != nil {
		log.Println("Failed to parse response body", err)
	}

	// fmt.Println(bytes)
	// fmt.Println(string(bytes))

	json.Unmarshal(bytes, &p.Homeworld)
}
