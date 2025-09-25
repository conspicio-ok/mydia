package main

import (
	"mydia/internal/db"
	"mydia/internal/services"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"

	// "errors"
	"log"
	"net/http"
	
	// "os/signal"
	// "syscall"
	"time"
)

// var conf = models.Config{
// 	Bind:			":8090",
// 	JWTSecret:		[]byte(os.Getenv("JWT_SECRET")),
// 	JWTIssuer:		os.Getenv("JWT_ISS"),
// 	JWTCookieName:	os.Getenv("JWT_COOKIE"),
// 	Domain:			os.Getenv("COOKIE_DOMAIN"),
// 	SecureCookie:	os.Getenv("COOKIE_SECURE"),
// 	SameSite:		http.SameSiteLaxMode,
// 	TTL:			14 * 24 * time.Hour,
// 	RenewIfLess:	7 * 24 * time.Hour,
// }

func main() {
	db := db.GetDB()
	if (db == nil) {
		return
	}
	defer db.Close()
	route := chi.NewRouter()
	route.Use(middleware.RequestID)
	route.Use(middleware.RealIP)
	route.Use(middleware.Recoverer)
	route.Use(middleware.Timeout(30 * time.Second))
	route.Use(middleware.Logger)

	route.Get("/users", service.ListUser(db))
	route.Get("/users/{username}", service.GetUser(db))

	// route.Group(func(r chi.Router) {
	// 	r.Use(AuthJWT)
		
		route.Post("/users", service.CreateUser(db))
		route.Put("/users/{username}", service.UpdateUser(db))
		route.Delete("/users/{username}", service.DeleteUser(db))
	// })	

	route.Post("/auth/login", service.Login(db))
	// Endpoint utilitaire pour proxy: valide le JWT cookie et renouvelle si nécessaire
	// route.Get("/auth/validate/{username}", service.GetToken(db))

	// srv := &http.Server{ Addr: conf.Bind, Handler: r }

	log.Println("Server starting on :8090")
    http.ListenAndServe(":8090", route)

	// go func() {
	// 	log.Println("listening on ", conf.Bind)
	// 	err := srv.ListenAndServe()
	// 	if (!errors.Is(err, http.ErrServerClosed)) {
	// 		log.Fatalf("server: %v", err)
	// 	}
	// }()

	// // Graceful shutdown
	// c := make(chan os.Signal, 1)
	// signal.Notify(c, syscall.SIGINT, syscall.SIGTERM)
	// <-c
	// ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	// defer cancel()
	// _ = srv.Shutdown(ctx)
	
}
