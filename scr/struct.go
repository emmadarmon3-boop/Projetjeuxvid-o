package jeux

type User struct {
	Pseudo string
	Id     int
}

func InitUser() User {
	return User{}
}

type Game struct {
	Title    string `json:"title"`
	Platform string `json:"platform"`
}

func InitGames() Game {
	return Game{}
}
