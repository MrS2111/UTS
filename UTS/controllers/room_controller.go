package controllers

import (
	"UTS/models"
	"encoding/json"
	"fmt"
	"net/http"
)

func GetAllRooms(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	query := "SELECT * FROM rooms"
	rows, err := db.Query(query)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !rows.Next() {
		var response models.RoomResponse
		response.Status = "404"
		response.Message = "Data not found"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	var rooms []models.Room
	for rows.Next() {
		var room models.Room
		rows.Scan(&room.ID, &room.Name, &room.ID_game)
		rooms = append(rooms, room)
	}
	var response models.RoomResponse
	response.Status = "200"
	response.Message = "Success"
	response.Data = rooms
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func GetRoomDetail(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	roomID := r.URL.Query().Get("id")
	query := "SELECT r.id, r.room_name, p.id, p.id_account, a.username FROM rooms r JOIN participants p ON p.id_room = r.id JOIN accounts a on a.id = p.id_account WHERE r.id = ?"
	rows, err := db.Query(query, roomID)
	if err != nil {
		fmt.Println(err)
		return
	}
	if !rows.Next() {
		var response models.RoomResponse
		response.Status = "404"
		response.Message = "Data not found"
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(response)
		return
	}
	var detailRoom models.DetailRoom
	var participants []models.Participant
	var accounts models.Account
	for rows.Next() {
		var participant models.Participant
		rows.Scan(&detailRoom.ID, &detailRoom.Name, &participant.ID, &participant.ID_room, &participant.ID_account, &accounts.Username)
		participants = append(participants, participant)
	}
	detailRoom.Participants = participants
	var response models.RoomDetailResponse
	response.Status = "200"
	response.Data = detailRoom
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func InsertRoom(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	var room models.Room
	json.NewDecoder(r.Body).Decode(&room)

	insertQuery := "INSERT INTO rooms (room_name, id_game) VALUES (?, ?)"
	_, err := db.Exec(insertQuery, room.Name, room.ID_game)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response models.RoomResponse
	response.Status = "200"
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}

func LeaveRoom(w http.ResponseWriter, r *http.Request) {
	db := Connect()
	defer db.Close()

	var participant models.Participant
	json.NewDecoder(r.Body).Decode(&participant)

	deleteQuery := "DELETE FROM participants WHERE id_account = ? AND id_room = ?"
	_, err := db.Exec(deleteQuery, participant.ID_account, participant.ID_room)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	var response models.RoomResponse
	response.Status = "200"
	response.Message = "Success"
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
