package service

import (
	"fmt"
	"time"
	"os"
	"database/sql"
	"log"

	"github.com/golang-jwt/jwt/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func	CreateJWT(username string, isAdmin bool) string {
    claims := JWTClaims{
        Username: username,
        Admin:  isAdmin,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)),
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
            Issuer:    "mydia",
        },
    }
    
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(jwtSecret)
    if (err != nil) {
		fmt.Println("Error token : ", err)
        return ""
    }
    
    return tokenString
}

func	UpdateJWT(db *sql.DB, userId int, jwt string) bool {
    updateQuery := `UPDATE JWT SET token = ?, expires = datetime('now', '+1 month') WHERE id_user = ?`
    result, err := db.Exec(updateQuery, jwt, userId)
    if (err != nil) {
        log.Println("JWT Update Error :", err)
        return false
    }
    
    rowsAffected, _ := result.RowsAffected()
    if (rowsAffected == 0) {
        insertQuery := "INSERT INTO JWT (id_user, token, expires) VALUES (?, ?, datetime('now', '+1 month'))"
        _, err = db.Exec(insertQuery, userId, jwt)
        if (err != nil) {
            log.Println("JWT Insertion Error :", err)
            return false
        }
    }
    return true
}

func	JWT(db *sql.DB, userId int, username string, admin bool) bool {
	var token string

	token = CreateJWT(username, admin)
	if (token == "") {
		return false
	}

	return UpdateJWT(db, userId, token)
}