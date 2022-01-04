package main

import (
	"fmt"
	"os"

	"github.com/MatthewCAlbert/devcode-backend-todo-challenge-21/internal/config"
	"github.com/MatthewCAlbert/devcode-backend-todo-challenge-21/internal/middleware"
	"github.com/MatthewCAlbert/devcode-backend-todo-challenge-21/internal/model"
	"github.com/MatthewCAlbert/devcode-backend-todo-challenge-21/internal/routes"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cache"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	// load application configurations
	cfg, err := config.Load()
	if err != nil {
		fmt.Println("Config failed")
		os.Exit(-1)
	}

	// setup and connect db
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.MySqlUser, cfg.MySqlPassword, cfg.MySqlHost, cfg.MySqlPort, cfg.MySqlDatabase,
	)
	db, dberr := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if dberr != nil {
		fmt.Println("DB Connection failed")
		os.Exit(-1)
	}

	// Auto Migrate
	migrateerr := db.AutoMigrate(&model.Activity{}, &model.Todo{})
	if migrateerr != nil {
		fmt.Println("Migration failed")
		os.Exit(-1)
	}

	// Random Insert

	// Init App
	app := fiber.New(fiber.Config{
		// Prefork:       true,
		CaseSensitive:         true,
		StrictRouting:         true,
		ServerHeader:          "Fiber",
		AppName:               "TODO Matthew DevCode",
		ErrorHandler:          middleware.ErrorHandler{}.Default,
		DisableStartupMessage: true,
	})
	app.Use(cache.New())

	// Init Routes
	routes := routes.Routes{
		Router: app, DB: db,
	}

	routes.Init()

	// Listen
	fmt.Println("Listening on port " + cfg.ServerPort)
	app.Listen(":" + cfg.ServerPort)
}
