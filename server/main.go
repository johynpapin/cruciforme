package main

import (
	"log"
	"os"

	"github.com/getsentry/sentry-go"
	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/johynpapin/cruciforme/server/email"
	"github.com/johynpapin/cruciforme/server/handlers"
	"github.com/johynpapin/cruciforme/server/store"
)

func loadEnv() {
	env := os.Getenv("CRUCIFORME_ENV")
	if "" == env {
		env = "development"
	}

	godotenv.Load(".env." + env + ".local")
	if "test" != env {
		godotenv.Load(".env.local")
	}
	godotenv.Load(".env." + env)
	godotenv.Load()
}

func main() {
	loadEnv()

	if err := sentry.Init(sentry.ClientOptions{
		Dsn: os.Getenv("CRUCIFORME_SENTRY_DSN"),
	}); err != nil {
		log.Fatalf("sentry initialization failed: %s", err)
	}

	r := gin.Default()

	r.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))
	r.Use(errorHandlerMiddleware())

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AllowHeaders = []string{"Origin", "Content-Length", "Content-Type", "Authorization"}

	r.Use(cors.New(corsConfig))

	h := handlers.Handlers{
		JwtSigningKey: []byte(os.Getenv("CRUCIFORME_JWT_SIGNING_KEY")),
		Store:         store.New(),
		EmailSender: email.NewSender(
			os.Getenv("CRUCIFORME_SENDGRID_API_KEY"),
			email.EmailAddress{
				Name:  "Cruciforme",
				Email: "no-reply@crucifor.me",
			},
		),
	}

	authGroup := r.Group("/auth")

	authGroup.POST("/sign-up", h.SignUpHandler)
	authGroup.POST("/sign-in", h.SignInHandler)
	authGroup.POST("/refresh", h.RefreshHandler)
	authGroup.POST("/verify", h.UserVerificationHandler)

	formsGroup := r.Group("/forms", authRequiredMiddleware(h.JwtSigningKey))

	formsGroup.GET("", h.GetFormsHandler)
	formsGroup.POST("", h.CreateFormHandler)

	if err := h.Store.Open(); err != nil {
		log.Panicf("Error opening the store: %v\n", err)
	}

	if err := r.Run(":8000"); err != nil {
		log.Panicf("Error running the server: %v\n", err)
	}
}
