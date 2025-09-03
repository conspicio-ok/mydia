package db

import (
	"log"
	"database/sql"
	
	_ "github.com/mattn/go-sqlite3"
)

func checkTables(db *sql.DB) bool {
	createTablesSQL := `
		CREATE TABLE IF NOT EXISTS USERS (
			id		INTEGER PRIMARY KEY AUTOINCREMENT,
			pseudo	VARCHAR(42) NOT NULL UNIQUE,
			passwd	CHAR(64) NOT NULL
		);

		CREATE TABLE IF NOT EXISTS JWT (
			id_user	INTEGER PRIMARY	KEY REFERENCES USERS(id) ON DELETE CASCADE,
			token	TEXT NOT NULL,
    		expires	DATETIME
		);
		
		CREATE TABLE IF NOT EXISTS RIGHTS (
			id_user		INTEGER PRIMARY KEY REFERENCES USERS(id) ON DELETE CASCADE,
			admin		BOOLEAN DEFAULT FALSE
		);		
		`
    _, err := db.Exec(createTablesSQL)
    if (err != nil) {
        log.Println("Erreur lors de la création des tables: ", err)
        return false
    }
    log.Println("Tables vérifiées/créées avec succès")
    return true
}

func	GetDB() *sql.DB {
	db, err := sql.Open("sqlite3", "users.sqlite")
	if (err != nil) {
		log.Println(err)
		return nil
	}
	if (!checkTables(db)) {
		return nil
	}
	return db
}