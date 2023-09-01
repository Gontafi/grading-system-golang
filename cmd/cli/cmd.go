package cli

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jackc/pgx/v5"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"

	"github.com/grading-system-golang/internal/config"
	"github.com/grading-system-golang/internal/handler"
	"github.com/grading-system-golang/internal/repositories"
	"github.com/grading-system-golang/internal/services"

	"context"
	"fmt"
	"log"
	"os"
	"time"
)

var (
	db          *pgx.Conn
	app         *fiber.App
	redisClient *redis.Client
)

var rootCmd = &cobra.Command{
	Use:   "grading-system",
	Short: "A grading system CLI",
}

var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Start the grading system",
	Run: func(cmd *cobra.Command, args []string) {
		startGradingSystem()
	},
}

var stopCmd = &cobra.Command{
	Use:   "stop",
	Short: "Stop the grading system",
	Run: func(cmd *cobra.Command, args []string) {
		stopGradingSystem()
	},
}

func init() {
	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(stopCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func startGradingSystem() {

	cfg, err := config.LoadConfig(os.Getenv("CONFIG_PATH"))
	if err != nil {
		log.Println("Failed to load config:", err)
		return
	}
	ctx := context.Background()

	log.Println("connecting to the database")
	databaseURL := fmt.Sprintf("postgres://%s:%s@%s:%d/%s?sslmode=disable",
		cfg.Database.User,
		cfg.Database.Password,
		cfg.Database.Host,
		cfg.Database.Port,
		cfg.Database.DBName,
	)
	db, err = pgx.Connect(ctx, databaseURL)
	if err != nil {
		log.Println("failed connection to the database", err)
		return
	}

	log.Println("connecting to the redis")
	redisClient = redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	_, err = redisClient.Ping(ctx).Result()
	if err != nil {
		log.Println("failed connection to the redis", err)
		return
	}

	repos := repositories.NewRepository(ctx, db)
	service := services.NewService(repos, redisClient, ctx, time.Second*30)
	handlers := handler.NewHandler(service)

	log.Println("starting server")
	app = fiber.New(fiber.Config{
		ReadTimeout:  cfg.Server.ReadTimeout,
		WriteTimeout: cfg.Server.WriteTimeout,
		IdleTimeout:  cfg.Server.IdleTimeout,
	})

	handlers.InitRoutes(app)

	err = app.Listen(":" + cfg.Server.Port)
	if err != nil {
		log.Println("failed to start server", err)
		return
	}
}

func stopGradingSystem() {
	if db != nil {
		err := db.Close(context.Background())
		if err != nil {
			log.Println("failed to close database postgres", err)
		}
	}

	if redisClient != nil {
		err := redisClient.Close()
		if err != nil {
			log.Println("failed to close redis", err)
		}
	}

	if app != nil {
		err := app.Shutdown()
		if err != nil {
			log.Println("failed to close server", err)
		}
	}
}
