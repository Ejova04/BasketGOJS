package service

import (
	"balldontlie/dto"
	"balldontlie/model"
	"balldontlie/repository"
	"fmt"
	"time"
)

func Get() ([]dto.Player, error) {
	players, err := repository.GetPlayers()
	if err != nil {
		return nil, err
	}

	dtoPlayers, err := createDtos(players)
	if err != nil {
		return nil, err
	}

	return dtoPlayers, nil
}

func Get2() ([]dto.Team, error) {
	teams, err := repository.GetTeams()
	if err != nil {
		return nil, err
	}

	dtoTeams, err := CreateDtosteam(teams)
	if err != nil {
		return nil, err
	}

	return dtoTeams, nil
}

func Get3() ([]dto.Game, error) {
	games, err := repository.GetGames()
	if err != nil {
		return nil, err
	}
	dtoGames, err := CreateDtosgame(games)
	if err != nil {
		return nil, err
	}
	return dtoGames, nil
}

func GetPlayerById(id int) (*dto.Player, error) {
	player, err := repository.GetPlayerById(id)
	if err != nil {
		return nil, err
	}

	fmt.Println(player)

	dtoPlayer, err := createDto(*player)
	if err != nil {
		return nil, err
	}

	return dtoPlayer, nil
}

func createDto(player model.Player) (*dto.Player, error) {
	dtoPlayer := &dto.Player{}

	dtoPlayer.ID = player.ID
	dtoPlayer.Weight = player.Weight
	dtoPlayer.Position = player.Position
	dtoPlayer.LastName = player.LastName
	dtoPlayer.FirstName = player.FirstName
	dtoPlayer.HeightInches = player.HeightInches
	dtoPlayer.HeightFeet = player.HeightFeet
	return dtoPlayer, nil
}

func createDto2(team model.Team) (*dto.Team, error) {
	dtoTeam := &dto.Team{}

	dtoTeam.ID = team.ID
	dtoTeam.Abbr = team.Abbr
	dtoTeam.City = team.City
	dtoTeam.Conference = team.Conference
	dtoTeam.Division = team.Division
	dtoTeam.FullName = team.FullName
	dtoTeam.Name = team.Name
	return dtoTeam, nil
}

func createDto3(game model.Game) (*dto.Game, error) {
	dtoGame := &dto.Game{}

	dtoGame.ID = game.ID
	dtoGame.Date = game.Date
	dtoGame.Season = game.Season
	dtoGame.HomeTeamScore = game.HomeTeamScore
	dtoGame.VisitorTeamScore = game.VisitorTeamScore
	dtoGame.Period = game.Period
	dtoGame.Status = game.Status
	dtoGame.Time = game.Time
	dtoGame.Postseason = game.Postseason
	return dtoGame, nil
}

func parallel(a model.Player, chanArt chan<- dto.Player) {

	dtoPlayer := dto.Player{}

	dtoPlayer.ID = a.ID
	dtoPlayer.Position = a.Position
	dtoPlayer.LastName = a.LastName
	dtoPlayer.FirstName = a.FirstName
	dtoPlayer.Weight = a.Weight
	dtoPlayer.HeightFeet = a.HeightFeet
	dtoPlayer.HeightInches = a.HeightInches

	chanArt <- dtoPlayer
}

func parallel2(a model.Team, chanTea chan<- dto.Team) {

	dtoTeam := dto.Team{}

	dtoTeam.ID = a.ID
	dtoTeam.Abbr = a.Abbr
	dtoTeam.City = a.City
	dtoTeam.Conference = a.Conference
	dtoTeam.Division = a.Division
	dtoTeam.FullName = a.FullName
	dtoTeam.Name = a.Name

	chanTea <- dtoTeam
}

func parallel3(a model.Game, chanGam chan<- dto.Game) {

	dtoGame := dto.Game{}

	dtoGame.ID = a.ID
	dtoGame.Date = a.Date
	dtoGame.Season = a.Season
	dtoGame.Period = a.Period
	dtoGame.Status = a.Status
	dtoGame.HomeTeamScore = a.HomeTeamScore
	dtoGame.VisitorTeamScore = a.VisitorTeamScore
	dtoGame.Postseason = a.Postseason

	chanGam <- dtoGame
}

func createDtos(players []model.Player) ([]dto.Player, error) {
	//var err error
	var dtoPlayers []dto.Player

	//wg := &sync.WaitGroup{}
	start := time.Now()

	chanArt := make(chan dto.Player)

	for _, a := range players {
		go parallel(a, chanArt)
	}

	for len(dtoPlayers) != len(players) { //54
		select {
		case elem := <-chanArt:
			{
				dtoPlayers = append(dtoPlayers, elem)
			}
		}
	}

	elapsed := time.Since(start)

	fmt.Printf("took %s \n", elapsed)

	return dtoPlayers, nil
}

func CreateDtosteam(teams []model.Team) ([]dto.Team, error) {
	var dtoTeams []dto.Team

	start := time.Now()
	chanTea := make(chan dto.Team)

	for _, a := range teams {
		go parallel2(a, chanTea)
	}

	for len(dtoTeams) != len(teams) {
		select {
		case elem := <-chanTea:
			{
				dtoTeams = append(dtoTeams, elem)
			}
		}
	}
	elapsed := time.Since(start)

	fmt.Printf("took %s \n", elapsed)

	return dtoTeams, nil
}

func CreateDtosgame(games []model.Game) ([]dto.Game, error) {
	var dtoGames []dto.Game

	start := time.Now()
	chanGam := make(chan dto.Game)

	for _, a := range games {
		go parallel3(a, chanGam)
	}

	for len(dtoGames) != len(games) {
		select {
		case elem := <-chanGam:
			{
				dtoGames = append(dtoGames, elem)
			}
		}
	}
	elapsed := time.Since(start)

	fmt.Printf("took %s \n", elapsed)

	return dtoGames, nil
}
