package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

type Response struct {
	Movie []Movie `json:"Search"`
}

type Movie struct {
	ImdbID string `json:"imdbID"`
	Title  string `json:"Title"`
	Year   int    `json:"Year"`
	Type   string `json:"Type"`
	Poster string `json:"Poster"`
}

type MovieDetail struct {
	ImdbID     string `json:"imdbID"`
	Title      string `json:"Title"`
	Year       int    `json:"Year"`
	Type       string `json:"Type"`
	Poster     string `json:"Poster"`
	Rated      string `json:"Rated"`
	Released   string `json:"Released"`
	Runtime    string `json:"Runtime"`
	Genre      string `json:"Genre"`
	Director   string `json:"Director"`
	Writer     string `json:"Writer"`
	Actors     string `json:"Actors"`
	Plot       string `json:"Plot"`
	Country    string `json:"Country"`
	ImdbRating string `json:"imdbRating"`
	ImdbVotes  string `json:"imdbVotes"`
}

type ResponseDataMovieList struct {
	Data  []Movie
	Total int
}

func goDotEnvVariable(key string) string {
	// load .env file
	err := godotenv.Load(".env")

	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	return os.Getenv(key)
}

func main() {
	log.Printf("Running")

	//Init Routes
	router := mux.NewRouter()
	router.HandleFunc("/", welcome)
	router.HandleFunc("/movie", getMovie).Methods("GET")
	router.HandleFunc("/movie/{id}", getMovieDetail).Methods("GET")

	//Log
	file, err := os.OpenFile("logActivity.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	log.SetOutput(file)
	log.Print("Logging to a file in Go!")

	// Init Serve
	serve := &http.Server{
		Handler:      router,
		Addr:         ":5000",
		WriteTimeout: 30 * time.Second,
		ReadTimeout:  40 * time.Second,
	}
	log.Fatal(serve.ListenAndServe())
}

func welcome(w http.ResponseWriter, r *http.Request) {
	log.Printf("/welcome")
	fmt.Println("Hello World")
}

func getMovie(w http.ResponseWriter, r *http.Request) {
	log.Printf("/movie")
	searchQuery := r.URL.Query().Get("searchword")
	pageQuery := r.URL.Query().Get("pagination")
	if len(searchQuery) == 0 || len(pageQuery) == 0 {
		log.Printf("The parameter is not defined")
		w.WriteHeader(http.StatusNoContent)
		return
	}
	apiUrl := goDotEnvVariable("API_URL") + "?apikey=" + goDotEnvVariable("API_KEY") + "&s=" + searchQuery + "&page=" + pageQuery

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	log.Printf("go to %s", apiUrl)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject Response
	json.Unmarshal(responseData, &responseObject)
	listMovie := ResponseDataMovieList{responseObject.Movie, len(responseObject.Movie)}
	data, err := json.Marshal(listMovie)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(data)
}

func getMovieDetail(w http.ResponseWriter, r *http.Request) {
	log.Printf("/detailMovie")
	movieId := mux.Vars(r)["id"]
	apiUrl := goDotEnvVariable("API_URL") + "?apikey=" + goDotEnvVariable("API_KEY") + "&i=" + movieId

	response, err := http.Get(apiUrl)
	if err != nil {
		fmt.Print(err.Error())
		os.Exit(1)
	}

	log.Printf("go to %s", apiUrl)

	responseData, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}

	var responseObject MovieDetail
	json.Unmarshal(responseData, &responseObject)
	json.NewEncoder(w).Encode(responseObject)
}
