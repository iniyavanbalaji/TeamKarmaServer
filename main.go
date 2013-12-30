package main

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

func getApiRouter() (s *mux.Router) {
	r := mux.NewRouter()
	s = r.PathPrefix("/api").Subrouter()
	s.HandleFunc("/sports", sportsHandler).Methods("GET")
	s.HandleFunc("/sports/{sport_id:[0-9]+}/leagues", leaguesHandler).Methods("GET")
	s.HandleFunc("/leagues/{league_id:[0-9]+}/teams", teamsHandler).Methods("GET")
	return
}

func sportsHandler(response http.ResponseWriter, request *http.Request) {
	sports := getSports()
	jsonResp, err := json.Marshal(sports)
	if err != nil {
		log.Fatal(err)
	}
	response.Write(jsonResp)
}

func leaguesHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	sportIdS := vars["sport_id"]
	sportId, err := strconv.ParseInt(sportIdS, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	sport := getSport(sportId)
	if sport == nil {
		http.NotFound(response, request)
		return
	}

	leagues := sport.getLeagues()
	jsonResp, err := json.Marshal(leagues)
	if err != nil {
		log.Fatal(err)
	}
	response.Write(jsonResp)
}

func teamsHandler(response http.ResponseWriter, request *http.Request) {
	vars := mux.Vars(request)
	leagueIdS := vars["league_id"]
	leagueId, err := strconv.ParseInt(leagueIdS, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	league := getLeague(leagueId)
	if league == nil {
		http.NotFound(response, request)
		return
	}

	teams := league.getCurrentSeasonTeams()
	jsonResp, err := json.Marshal(teams)
	if err != nil {
		log.Fatal(err)
	}
	response.Write(jsonResp)
}

func main() {
	apiRouter := getApiRouter()
	http.Handle("/", apiRouter)
	http.ListenAndServe(":3000", nil)
}
