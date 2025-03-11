package db

import (
	"context"
	"github.com/jackc/pgx/v5"
	"log"
	"time"
)

var DB *pgx.Conn

func OpenDatabaseConnection() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, "host=localhost user=adiinfiniteloop password=mysecretpassword dbname=sos-db port=5432 sslmode=disable")
	if err != nil {
		log.Println("Error in Connecting to DB", err)
	}
	DB = conn

	log.Println("Connected To DB Successfully")

}
func CloseDatabaseConnection() {
	if DB != nil {
		err := DB.Close(context.Background())
		if err != nil {
			log.Println("❌ Error Closing DB Connection:", err)
		} else {
			log.Println("✅ Database Connection Closed Successfully")
		}
	}
}
