package service

import (
	"database/sql"
	"log"
	"regexp"
)

// Av -> Available
func	AvUsername(db *sql.DB, username string) bool {
	var existingID int
	checkQuery := "SELECT id FROM USERS WHERE pseudo = ?"
	err := db.QueryRow(checkQuery, username).Scan(&existingID)
	if (err != sql.ErrNoRows) {
		return false
	}
	if (err != nil) {
		log.Println("Servor Error")
		return false
	}
	return true
}

func	StrongPassword(password string) bool {
	strongRegex := regexp.MustCompile(`^[A-Za-z\d@$!%*?&]{8,}$`)
	if (!strongRegex.MatchString(password)) {
		return false
	}
	return true
}
