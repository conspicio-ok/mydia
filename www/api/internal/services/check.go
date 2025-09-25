package service

import (
	"database/sql"
	"log"
	"regexp"
	"crypto/sha256"
	"encoding/hex"
)

// Av -> Available
func	AvUsername(db *sql.DB, username string) bool {
	var existingID int
	checkQuery := "SELECT id FROM USERS WHERE pseudo = ?"
	err := db.QueryRow(checkQuery, username).Scan(&existingID)
	if (err == sql.ErrNoRows) {
		return true
	}
	if (err != nil) {
		log.Println("Servor Error")
		return false
	}
	return false
}

func	StrongPassword(password string) bool {
	strongRegex := regexp.MustCompile(`^[A-Za-z\d@$!%*?&]{8,}$`)
	if (!strongRegex.MatchString(password)) {
		return false
	}
	return true
}

func	hashPassword(password string) string {
    hash := sha256.Sum256([]byte(password))
    return hex.EncodeToString(hash[:])
}
