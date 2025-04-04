package main

import (
	"context"
	"log"
	"log/slog"
	"os"

	"ctf01d/internal/config"
	"ctf01d/internal/handler"
	"ctf01d/internal/httpserver"
	migration "ctf01d/internal/migrations/psql"
	"ctf01d/pkg/ginmiddleware"

	"github.com/getkin/kin-openapi/openapi3"
	"github.com/getkin/kin-openapi/openapi3filter"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
	_ "go.uber.org/automaxprocs"
)

func main() {
	path := "./configs/config.development.yml"
	if envPath, exists := os.LookupEnv("CONFIG_PATH"); exists {
		path = envPath
	}

	cfg, err := config.New(path)
	if err != nil {
		slog.Error("Config error: " + err.Error())
		os.Exit(1)
	}
	logger := slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: slog.Level(cfg.ParseLogLevel(cfg.Log.Level)),
	}))
	slog.SetDefault(logger)
	slog.Info("Config path is - " + path)
	db, err := migration.InitDatabase(cfg)
	if err != nil {
		slog.Error("Database connection error: " + err.Error())
		os.Exit(1)
	}
	defer db.Close()
	slog.Info("Database connection established successfully")
	router := gin.New()
	router.Use(gin.Recovery(), gin.Logger())

	// Добавление CORS middleware
	router.Use(cors.Default())

	// Загрузка спецификации OpenAPI
	swagger, err := openapi3.NewLoader().LoadFromFile("api/openapi.yaml")
	if err != nil {
		log.Fatalf("Failed to load OpenAPI spec: %v", err)
	}

	// OpenAPI middleware валидации запросов
	requestValidator := ginmiddleware.OapiRequestValidatorWithOptions(swagger, &ginmiddleware.Options{
		Options: openapi3filter.Options{
			AuthenticationFunc: func(_ context.Context, _ *openapi3filter.AuthenticationInput) error {
				return nil
			},
		},
	})

	// OpenAPI middleware валидации ответов
	responseValidator := ginmiddleware.OapiResponseValidator(swagger)
	hndlr := &handler.Handler{DB: db}

	// API-группа, к которой применяются валидаторы
	apiGroup := router.Group("/", requestValidator, responseValidator)
	httpserver.RegisterHandlers(apiGroup, hndlr)

	// HTML маршрутизатор для корня
	router.GET("/", httpserver.NewHtmlRouter())
	router.NoRoute(httpserver.NewHtmlRouter())

	// Запуск сервера
	addr := cfg.HTTP.Host + ":" + cfg.HTTP.Port
	slog.Info("Server running on", slog.String("address", addr))
	if err := router.Run(addr); err != nil {
		slog.Error("Server error: " + err.Error())
		os.Exit(1)
	}
}
