package dalbase

import (
	"database/sql"

	_ "github.com/proullon/ramsql/driver"
)

func GetDB() *sql.DB {
	dbBase := []string{
		`CREATE TABLE userAds (id int NOT NULL AUTO_INCREMENT, userID VARCHAR(256), videoID VARCHAR(256), timestamp TEXT, PRIMARY KEY (id));`,
	}

	db, err := sql.Open("ramsql", "PlayerWon")
	if err != nil {
		panic(err)
	}

	for _, dbData := range dbBase {
		_, err := db.Exec(dbData)
		if err != nil {
			panic(err)
		}
	}

	return db
}
