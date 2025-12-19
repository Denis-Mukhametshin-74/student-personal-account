package main

import (
	"log"
	"net/http"
	"os"
	"student-personal-account/internal/app/handler"
	"student-personal-account/internal/app/middleware"
	"student-personal-account/internal/app/repository"
	"student-personal-account/internal/app/service"
	"student-personal-account/pkg/database"
	"student-personal-account/pkg/jwt"
	"time"

	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Println("No .env file found, using environment variables")
	}

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	jwtConfig := jwt.Config{
		SecretKey:  os.Getenv("JWT_SECRET"),
		Expiration: os.Getenv("JWT_EXPIRATION"),
	}

	if err := jwt.InitJWT(jwtConfig); err != nil {
		log.Fatal("Failed to initialize JWT:", err)
	}

	dbConfig := database.Config{
		Host:     os.Getenv("DB_HOST"),
		Port:     os.Getenv("DB_PORT"),
		User:     os.Getenv("DB_USER"),
		Password: os.Getenv("DB_PASSWORD"),
		Name:     os.Getenv("DB_NAME"),
	}

	if err := database.InitDB(dbConfig); err != nil {
		log.Fatal("Failed to connect to database:", err)
	}

	scheduleRepo := repository.NewScheduleRepository(database.DB)
	studentRepo := repository.NewStudentRepository(database.DB)
	subjectRepo := repository.NewSubjectRepository(database.DB)

	authService := service.NewAuthService(studentRepo)
	profileService := service.NewProfileService(studentRepo, subjectRepo)
	scheduleService := service.NewScheduleService(scheduleRepo, studentRepo)

	authHandler := handler.NewAuthHandler(authService)
	profileHandler := handler.NewProfileHandler(profileService)
	scheduleHandler := handler.NewScheduleHandler(scheduleService)

	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("./web/static/"))

	mux.Handle("/", fs)

	mux.HandleFunc("POST /api/auth/login", authHandler.Login)
	mux.HandleFunc("POST /api/auth/register", authHandler.Register)

	mux.Handle("GET /api/profile", middleware.AuthMiddleware(
		http.HandlerFunc(profileHandler.GetProfile)))

	mux.Handle("GET /api/schedule", middleware.AuthMiddleware(
		http.HandlerFunc(scheduleHandler.GetSchedule)))

	server := &http.Server{
		Addr:         ":" + port,
		Handler:      mux,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
	}

	log.Println("The server is running on port", port)
	log.Fatal(server.ListenAndServe())
}
