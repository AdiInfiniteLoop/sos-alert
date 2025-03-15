package handlers

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"sos-alert/db"
	"sos-alert/pkg"
)

type SOS struct {
	Id      int    `json:"id"`
	Message string `json:"message"`
	Author  string `json:"author"`
}

func PostAlert(w http.ResponseWriter, r *http.Request) {
	var alert SOS

	//body, err := io.ReadAll(r.Body)
	//if err != nil {
	//	log.Println("Error reading request body:", err)
	//	http.Error(w, "Cannot read body", http.StatusBadRequest)
	//	return
	//}
	//fmt.Println(string(body)) //Printing the request bod //after reading body the r.Body becomes empty

	err := json.NewDecoder(r.Body).Decode(&alert)
	if err != nil {
		log.Println("Parsing to Struct Leading to Error", err)
	}
	//fmt.Println(alert)

	query := `INSERT INTO alerts(MESSAGE, AUTHOR) VALUES($1, $2) RETURNING id`
	var alertID int
	err = db.DB.QueryRow(context.Background(), query, alert.Message, alert.Author).Scan(&alertID)
	if err != nil {
		log.Println("❌ Database Insert Error:", err)
		http.Error(w, "Failed to store alert", http.StatusInternalServerError)
		return
	}
	pkg.RespondWithJSON(w, 200, alert)
}

func GetAlert(w http.ResponseWriter, request *http.Request) {
	query := "SELECT * FROM alerts"

	rows, err := db.DB.Query(context.Background(), query)
	if err != nil {
		log.Println("DB Error", err)
	}

	var alerts []SOS
	defer rows.Close()
	for rows.Next() {
		var alert SOS
		err := rows.Scan(&alert.Id, &alert.Message, &alert.Author)
		if err != nil {
			log.Println("error scanning the row", err)
			return
		}
		alerts = append(alerts, alert)
	}

	pkg.RespondWithJSON(w, 200, alerts)
}
