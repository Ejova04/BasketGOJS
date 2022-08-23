package dto

type Player struct {
	ID           int    `json:"id"`
	FirstName    string `json:"first_name"`
	LastName     string `json:"last_name"`
	Position     string `json:"position"`
	HeightFeet   int    `json:"height_feet"`
	HeightInches int    `json:"height_inches"`
	Weight       int    `json:"weight_pounds"`
	Team         Team   `json:"team"`
}

type Team struct {
	ID         int    `json:"id"`
	Abbr       string `json:"abbreviation"`
	City       string `json:"city"`
	Conference string `json:"conference"`
	Division   string `json:"division"`
	FullName   string `json:"full_name"`
	Name       string `json:"name"`
}

type Game struct {
	ID               int    `json:"id"`
	Date             string `json:"date"`
	Season           int    `json:"season"`
	HomeTeamScore    int    `json:"home_team_score"`
	VisitorTeamScore int    `json:"visitor_team_score"`
	Period           int    `json:"period"`
	Status           string `json:"status"`
	Time             string `json:"time"`
	Postseason       bool   `json:"postseason"`
	HomeTeam         Team   `json:"home_team"`
	VisitorTeam      Team   `json:"visitor_team"`
}
