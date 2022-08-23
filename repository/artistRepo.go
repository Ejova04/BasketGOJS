package repository

import (
	"balldontlie/model"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"time"
)

const playersUrl string = "https://www.balldontlie.io/api/v1/players"
const teamsUrl string = "https://www.balldontlie.io/api/v1/teams"
const gamesUrl string = "https://www.balldontlie.io/api/v1/games"

var client = &http.Client{}

func GetPlayers() ([]model.Player, error) {
	var players []model.Player
	err := getbench(playersUrl, &players)
	if err != nil {
		return nil, err
	}

	return players, nil
}

func GetPlayerById(id int) (*model.Player, error) {
	player := &model.Player{}

	err := get(playersUrl+"/"+strconv.Itoa(id), &player)
	if err != nil {
		return nil, err
	}

	return player, nil
}

func GetTeams() ([]model.Team, error) {
	var teams []model.Team
	err := getbench(teamsUrl, &teams)
	if err != nil {
		return nil, err
	}

	return teams, nil
}

func GetTeamById(id int) (*model.Team, error) {
	team := &model.Team{}

	err := get(teamsUrl+"/"+strconv.Itoa(id), &team)
	if err != nil {
		return nil, err
	}

	return team, nil
}

func GetGames() ([]model.Game, error) {
	var games []model.Game
	err := getbench(gamesUrl, &games)
	if err != nil {
		return nil, err
	}

	return games, nil
}

func GetGameById(id int) (*model.Game, error) {
	game := &model.Game{}

	err := get(gamesUrl+"/"+strconv.Itoa(id), &game)
	if err != nil {
		return nil, err
	}

	return game, nil
}

func get(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	err = json.NewDecoder(r.Body).Decode(target)
	if err != nil {
		return err
	}

	return nil
}

func getbench(url string, target interface{}) error {
	r, err := client.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	start := time.Now()

	err = json.NewDecoder(r.Body).Decode(target)

	elapsed := time.Since(start)
	fmt.Printf("all took %s \n", elapsed)

	if err != nil {
		return err
	}

	return nil
}
