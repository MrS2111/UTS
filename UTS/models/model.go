package models

type Account struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
}
type Game struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Max_player int `json:"max_player"`
}
type Room struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	ID_game int `json:"id_game"`
}
type DetailRoom struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Participants []Participant `json:"participants"`
}

type Participant struct {
	ID       int    `json:"id"`
	ID_room int `json:"id_room"`
	ID_account int `json:"id_account"`
}
type AccountResponse struct {
	ID       int    `json:"id"`
	Username string `json:"username"`
	Password string `json:"password"`
}
type GameResponse struct {
	ID       int    `json:"id"`
	Name     string `json:"name"`
	Max_player int `json:"max_player"`
}
type RoomResponse struct {
	Status     string `json:"status"`
	Message    string `json:"message"`
	Data []Room `json:"data"`
}
type ParticipantResponse struct {
	ID       int    `json:"id"`
	ID_room int `json:"id_room"`
	ID_account int `json:"id_account"`
}
type RoomDetailResponse struct {
 	Status     string `json:"status"`
 	Data DetailRoom `json:"data"`
}
