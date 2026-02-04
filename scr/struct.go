package jeux

type Game struct {
	Title    string `json:"title"`
	Platform string `json:"platform"`
}

// type Map struct {
// 	Data map[string]Games
// }

func InitGames() Game {
	return Game{}
}
