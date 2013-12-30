package main

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"log"
	"os"
	"os/user"
	"path"
	"time"
)

var dbFile string

func init() {
	tuser, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	dbFile = path.Join(tuser.HomeDir, "karma.db")

	if _, err := os.Stat(dbFile); os.IsNotExist(err) {
		log.Fatalf("Database file: %[1]s does not exist. Create a database file using 'sqlite3 -init init.sql %[1]s'\n", dbFile)
	}
}

type Sport struct {
	Id   int64  `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

func getSport(sportId int64) *Sport {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	err = db.QueryRow("SELECT name_en FROM sport WHERE id=?", sportId).Scan(&name)
	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		log.Fatal(err)
		return nil
	default:
		return &Sport{Id: sportId, Name: name}
	}
}
func (sport *Sport) getLeagues() []*League {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	leagues := make([]*League, 0)

	rows, err := db.Query("SELECT id, name FROM league WHERE sport_id =? ORDER BY name", sport.Id)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int64
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		leagues = append(leagues, &League{Id: id, Name: name})
	}

	return leagues
}

type League struct {
	Id   int64  `json:"id" xml:"id"`
	Name string `json:"name" xml:"name"`
}

func getLeague(leagueId int64) (league *League) {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	var name string
	err = db.QueryRow("SELECT name FROM league WHERE id=?", leagueId).Scan(&name)
	switch {
	case err == sql.ErrNoRows:
		return nil
	case err != nil:
		log.Fatal(err)
		return nil
	default:
		return &League{Id: leagueId, Name: name}
	}
}
func (league *League) getCurrentSeasonTeams() []*Team {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	teams := make([]*Team, 0)

	seasonId := league.getCurrentSeasonId(db)
	// no current season id was found for the League
	if seasonId == 0 {
		return teams
	}

	rows, err := db.Query(
		`SELECT lt.id, lt.name, lt.short_name FROM league_team lt 
		JOIN league_team_group ltg ON (lt.id = ltg.league_team_id) 
		JOIN league_season_group lsg ON (lsg.id = ltg.league_season_group_id) 
		WHERE lt.league_id=? 
		AND   lsg.league_season_id=? 
		ORDER BY lt.name`,
		league.Id, seasonId)
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int64
		var name string
		var shortName string
		if err := rows.Scan(&id, &name, &shortName); err != nil {
			log.Fatal(err)
		}
		teams = append(teams, &Team{Id: id, Name: name, ShortName: shortName})
	}

	return teams
}
func (league *League) getCurrentSeasonId(db *sql.DB) (seasonId int32) {
	// get the league_season_id of the current season for the league
	query := `SELECT id 
		FROM league_season 
		WHERE league_id=? 
		AND start_date = (
		SELECT MAX(start_date) FROM league_season WHERE league_id=? AND start_date<=?
		)`
	err := db.QueryRow(query, league.Id, league.Id, time.Now()).Scan(&seasonId)
	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
	}
	return
}

type Team struct {
	Id        int64  `json:"id" xml:"id"`
	ShortName string `json:"short_name" xml:"short_name"`
	Name      string `json:"name" xml:"name"`
}

func getSports() []*Sport {
	db, err := sql.Open("sqlite3", dbFile)
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	sports := make([]*Sport, 0)

	rows, err := db.Query("SELECT id, name_en FROM sport")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var id int64
		var name string
		if err := rows.Scan(&id, &name); err != nil {
			log.Fatal(err)
		}
		sports = append(sports, &Sport{Id: id, Name: name})
	}

	return sports
}
