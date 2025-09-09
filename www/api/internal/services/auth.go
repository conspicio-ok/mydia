package service

import (
	"log"
	"database/sql"
	"encoding/json"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
// 	"github.com/go-chi/chi/v5"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

func CreateJWT(username string, isAdmin bool) (string, error) {
    // Définir les claims
    claims := JWTClaims{
        Username: username,
        Admin:  isAdmin,
        RegisteredClaims: jwt.RegisteredClaims{
            ExpiresAt: jwt.NewNumericDate(time.Now().Add(30 * 24 * time.Hour)), // Expire dans 30 * 24h
            IssuedAt:  jwt.NewNumericDate(time.Now()),
            NotBefore: jwt.NewNumericDate(time.Now()),
            Issuer:    "mydia", // Nom de ton application
        },
    }
    
    // Créer le token avec les claims
    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    
    // Signer le token avec la clé secrète
    tokenString, err := token.SignedString(jwtSecret)
    if err != nil {
        return "", err
    }
    
    return tokenString, nil
}

func	Login(db *sql.DB) http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        if (r.Header.Get("Content-Type") != "application/json") {
            http.Error(w, "Content-Type must be application/json", http.StatusBadRequest)
            return
        }

        var req User
		var	userPassword string
		err := json.NewDecoder(r.Body).Decode(&req)
        if (err != nil) {
            http.Error(w, "Invalid JSON", http.StatusBadRequest)
            return
        }
        hashedPassword := hashPassword(req.Password)

		getUserQuery := "SELECT passwd FROM USERS WHERE pseudo = ?;"
        err = db.QueryRow(getUserQuery, req.Pseudo).Scan(&userPassword)
		if (err == sql.ErrNoRows) {
			http.Error(w, "User not found", http.StatusNotFound)
			return
		}
        if (err != nil) {
			log.Println("Error : ", err)
            http.Error(w, "Server Error", http.StatusInternalServerError)
            return
        }

		if (userPassword != hashedPassword) {
			http.Error(w, "Wrong password", http.StatusConflict)
			return
		}

		jwt, err := CreateJWT(req.Pseudo, true)
		if (err != nil) {
			log.Println("Error : ", err)
			http.Error(w, "Server Error", http.StatusInternalServerError)
			return
		}

        response := MessageResponse{
            Status:     "success",
            Message:  	jwt,
        }

        w.Header().Set("Content-Type", "application/json")
        w.WriteHeader(http.StatusCreated)
        json.NewEncoder(w).Encode(response)
    }
}


// func ValidateToken(w http.ResponseWriter, r *http.Request) {
//     authHeader := r.Header.Get("Authorization")
//     if !strings.HasPrefix(authHeader, "Bearer ") {
//         w.WriteHeader(http.StatusUnauthorized)
//         return
//     }
    
//     token := strings.TrimPrefix(authHeader, "Bearer ")
//     claims, err := ValidateJWT(token)
//     if err != nil {
//         w.WriteHeader(http.StatusUnauthorized)
//         return
//     }
    
//     // Retourner les infos utilisateur dans les headers
//     w.Header().Set("X-User", claims.Username)
//     w.Header().Set("X-User-ID", strconv.Itoa(claims.UserID))
//     w.WriteHeader(http.StatusOK)
// }




// // import (
// //     "fmt"
// //     "github.com/golang-jwt/jwt/v5"
// // )



// // func ValidateJWT(tokenString string) (*JWTClaims, error) {
// //     // Parser le token
// //     token, err := jwt.ParseWithClaims(tokenString, &JWTClaims{}, func(token *jwt.Token) (interface{}, error) {
// //         // Vérifier que la méthode de signature est bien HMAC
// //         if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
// //             return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
// //         }
// //         return jwtSecret, nil
// //     })
    
// //     if err != nil {
// //         return nil, err
// //     }
    
// //     // Vérifier que le token est valide
// //     if !token.Valid {
// //         return nil, fmt.Errorf("invalid token")
// //     }
    
// //     // Récupérer les claims
// //     claims, ok := token.Claims.(*JWTClaims)
// //     if !ok {
// //         return nil, fmt.Errorf("invalid token claims")
// //     }
    
// //     return claims, nil
// // }

// // func AuthJWT(next http.Handler) http.Handler {
// //     return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
// //         // Récupérer le token depuis le header
// //         authHeader := r.Header.Get("Authorization")
// //         if (!strings.HasPrefix(authHeader, "Bearer ")) {
// //             http.Error(w, "Missing or invalid authorization header", http.StatusUnauthorized)
// //             return
// //         }
        
// //         tokenString := strings.TrimPrefix(authHeader, "Bearer ")
        
// //         // Valider le JWT (tu devras implémenter cette fonction)
// //         claims, err := ValidateJWT(tokenString)
// //         if (err != nil) {
// //             http.Error(w, "Invalid token", http.StatusUnauthorized)
// //             return
// //         }
        
// //         // Stocker les infos dans le context
// //         ctx := context.WithValue(r.Context(), "user_id", claims.UserID)
// //         ctx = context.WithValue(ctx, "admin", claims.Admin)
        
// //         // Continuer vers le handler suivant
// //         next.ServeHTTP(w, r.WithContext(ctx))
// //     })
// // }

// // func	GetToken(db *sql.DB) http.HandlerFunc {
// // 	return func(w http.ResponseWriter, r *http.Request) {
// // 		var token string
// // 		username := chi.URLParam(r, "username")
// // 		query := "SELECT token FROM JWT jwt JOIN USERS u ON u.id = jwt.id_user WHERE pseudo = ?"

// // 		err := db.QueryRow(query, username).Scan(&token)
// // 		if (err == sql.ErrNoRows) {
// // 			http.Error(w, "JWT not found", http.StatusNotFound)
// // 			return
// // 		}
// // 		if (err != nil) {
// // 			log.Println("Error : ", err)
// // 			http.Error(w, "Server Error", http.StatusInternalServerError)
// // 			return
// // 		}

// // 		w.Header().Set("Content-Type", "application/json")
// // 		json.NewEncoder(w).Encode(token)
// // 	}
// // }

