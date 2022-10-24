package main

import (
	"context"
	"fishki/internal/app"
	"fishki/internal/db"
	"fishki/internal/rest"
	"fishki/pkg/core"
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"os"
)

const databaseName = "fishki"

func main() {
	router := gin.Default()

	config := cors.DefaultConfig()
	var origins = []string{"http://fishki-client:5174", "http://127.0.0.1:5174",
		"http://127.0.0.1:5173", "http://localhost:5173/"}

	if corsEnv, found := os.LookupEnv("CORS_ORIGIN"); found {
		origins = append(origins, corsEnv)
	}

	config.AllowOrigins = origins

	if gin.Mode() == gin.DebugMode {
		for _, origin := range config.AllowOrigins {
			fmt.Println("Origin allowed:", origin)
		}
	}
	config.AllowCredentials = true

	router.Use(cors.New(config))

	_ = router.SetTrustedProxies(nil)
	credentials := options.Credential{Username: os.Getenv("MONGO_USERNAME"), Password: os.Getenv("MONGO_PASSWORD")}
	mongoClient, err := mongo.Connect(context.Background(),
		options.Client().ApplyURI("mongodb://mongo:27017").SetAuth(credentials))
	if err != nil {
		panic(err)
	}

	defer func() {
		if err := mongoClient.Disconnect(context.TODO()); err != nil {
			panic(err)
		}
	}()

	mongoDb := mongoClient.Database(databaseName)

	userRepo := db.NewUserRepository(mongoDb)
	userService := app.NewUserService(userRepo)
	userHandler := rest.NewUserHandler(userService)

	sessionMiddleware := rest.SessionMiddleware{Service: userService}
	authorized := router.Group("/")
	authorized.Use(sessionMiddleware.Session)

	userHandler.Setup(router, authorized)

	cardRepo := db.NewCardRepository(mongoDb)
	cardService := app.NewCardService(cardRepo)
	cardHandler := rest.NewCardHandler(cardService)
	cardHandler.Setup(router, authorized)

	deckRepo := db.NewDeckRepository(mongoDb)
	deckService := app.NewDeckService(deckRepo, cardService, userService)
	deckHandler := rest.NewDeckHandler(deckService)
	deckHandler.Setup(router, authorized)

	cardService.AddCreatedCardObserver(deckService.AddCard)
	cardService.AddDeletedCardObserver(deckService.RemoveCard)

	reviewRepo := db.NewReviewRepository(mongoDb)
	reviewService := app.NewReviewService(reviewRepo, &userService)

	fixedRepo := db.NewFixedModeRepository(mongoDb)
	fixedService := app.NewFixedModeService(fixedRepo, deckService)
	fixedHandler := rest.NewFixedModeHandler(&fixedService)

	reviewService.RegisterModeService(core.Fixed, fixedService)

	reviewHandler := rest.NewReviewHandler(&reviewService)
	reviewHandler.RegisterModeHandler(core.Fixed, fixedHandler)
	reviewHandler.Setup(router, authorized)

	answerService := app.NewAnswerService(reviewService)
	answerHandler := rest.NewAnswerHandler(answerService)
	answerHandler.Setup(router, authorized)

	err = router.Run("0.0.0.0:8080")

	if err != nil {
		panic(err)
	}
}
