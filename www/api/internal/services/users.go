package service

import (
	"log"
	"database/sql"
	"encoding/json"
	"net/http"
	"encoding/hex"
	"crypto/sha256"
	"strings"

	"github.com/go-chi/chi/v5"
)

type User struct {
	Pseudo		string	`json:"pseudo"`
	Password	string	`json:"password"`
}

type CreateUserResponse struct {
	ID		int		`json:"id"`
	Pseudo	string	`json:"pseudo"`
	Message	string	`json:"message"`
}

type UpdateUserResponse struct {
	Pseudo	string	`json:"pseudo"`
	Message	string	`json:"message"`
}

func	hashPassword(password string) string {
    hash := sha256.Sum256([]byte(password))
    return hex.EncodeToString(hash[:])
}

func	ListUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var users []User
		query := "SELECT pseudo FROM USERS"
		rows, err := db.Query(query)
		if (err != nil) {
			log.Println("Error : select users - ", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}
		defer rows.Close()

		for rows.Next() {
			var user User
			err := rows.Scan(&user.Pseudo)
			if (err == sql.ErrNoRows) {
				http.Error(w, "User not found", http.StatusNotFound)
				return
			}
			if (err != nil) {
				log.Println("Error : ", err)
				http.Error(w, "Server Error", http.StatusInternalServerError)
				return
			}
			users = append(users, user)
		}

		err = rows.Err()
		if (err != nil) {
			log.Println(err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(users)
	}
}

func	GetUser(db *sql.DB) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var user User
		username := chi.URLParam(r, "username")
		query := "SELECT pseudo FROM USERS WHERE pseudo = ?"

		err := db.QueryRow(query, username).Scan(&user.Pseudo)
		if (err == sql.ErrNoRows) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
		if (err != nil) {
			log.Println("Error : ", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(user)
	}
}

func	CreateUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if (r.Header.Get("Content-Type") != "application/json") {
            http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
            return
        }

        var req User
		err := json.NewDecoder(r.Body).Decode(&req)
        if (err != nil) {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }

        if (strings.TrimSpace(req.Pseudo) == "") {
            http.Error(w, "Create Account Required Username", http.StatusBadRequest)
            return
        }
		if (!StrongPassword(req.Password)) {
			http.Error(w, "Create Account Required Strong Password", http.StatusBadRequest)
			return
		}

		if (AvUsername(db, req.Pseudo)) {
			http.Error(w, "User already exist, please change your username", http.StatusConflict)
			return
		}

        hashedPassword := hashPassword(req.Password)
		insertQuery := "INSERT INTO USERS (pseudo, passwd) VALUES (?, ?)"
        result, err := db.Exec(insertQuery, req.Pseudo, hashedPassword)
        if (err != nil) {
            http.Error(w, "User Creation Error", http.StatusInternalServerError)
            return
        }

        userID, err := result.LastInsertId()
        response := CreateUserResponse{
            ID:      int(userID),
            Pseudo:  req.Pseudo,
            Message: "User Created",
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(response)
    }
}

func	UpdateUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if (r.Header.Get("Content-Type") != "application/json") {
            http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
            return
        }
        username := chi.URLParam(r, "username")
        if (username == "") {
            http.Error(w, "Username required", http.StatusBadRequest)
            return
        }
        var req User
        err := json.NewDecoder(r.Body).Decode(&req)
        if (err != nil) {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
		if (AvUsername(db, username)) {
			http.Error(w, "User not found or Error Server", http.StatusNotFound)
			return
		}

        setParts := []string{}
        args := []interface{}{}
        if (strings.TrimSpace(req.Pseudo) != "") {
            if (!AvUsername(db, req.Pseudo)) {
                http.Error(w, "Username already exists or Error Server", http.StatusConflict)
                return
            }
            setParts = append(setParts, "pseudo = ?")
            args = append(args, req.Pseudo)
        }
        if (strings.TrimSpace(req.Password) != "") {
            if (!StrongPassword(req.Password)) {
                http.Error(w, "Modify Account Required Strong Password", http.StatusBadRequest)
                return
            }
            hashedPassword := hashPassword(req.Password)
            setParts = append(setParts, "passwd = ?")
            args = append(args, hashedPassword)
        }
        if (len(setParts) == 0) {
            http.Error(w, "No fields to update", http.StatusBadRequest)
            return
        }

        updateQuery := "UPDATE USERS SET " + strings.Join(setParts, ", ") + " WHERE pseudo = ?"
        args = append(args, username)
        result, err := db.Exec(updateQuery, args...)
        if (err != nil) {
            log.Println("Error updating user:", err)
            http.Error(w, "User modification error", http.StatusInternalServerError)
            return
        }
        rowsAffected, err := result.RowsAffected()
        if (err != nil || rowsAffected == 0) {
            http.Error(w, "No user was updated", http.StatusInternalServerError)
            return
        }

        finalPseudo := username
        if (strings.TrimSpace(req.Pseudo) != "") {
            finalPseudo = req.Pseudo
        }
        response := UpdateUserResponse{
            Pseudo:  finalPseudo,
            Message: "User updated successfully",
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}

func	DeleteUser(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        username := chi.URLParam(r, "username")
        if (username == "") {
            http.Error(w, "Username required", http.StatusBadRequest)
            return
        }
		if (AvUsername(db, username)) {
			http.Error(w, "User not found or Error Server", http.StatusNotFound)
			return
		}

        deleteQuery := "DELETE FROM USERS WHERE pseudo = ?"
        result, err := db.Exec(deleteQuery, username)
        if (err != nil) {
            log.Println("Error deleting user:", err)
            http.Error(w, "User deleting error", http.StatusInternalServerError)
            return
        }
		rowsAffected, err := result.RowsAffected()
        if (err != nil || rowsAffected == 0) {
            http.Error(w, "No user was deleted", http.StatusInternalServerError)
            return
        }
		
        response := map[string]string{
            "message": username + " deleted successfully",
        }
        w.Header().Set("Content-Type", "application/json")
        json.NewEncoder(w).Encode(response)
    }
}
