// router.go
//
// Author: Lin Jun
// Created: 2025-06-18
//
// License: MIT License
//
// Description:
// This file initializes the Gin router, sets up API route groups,
// and registers both API routes and static file routes.
//
// ------------------------------------------------------------

package router

import (
	"github.com/0x716/Turnipin/internal/config"
	"github.com/0x716/Turnipin/internal/router/routes"
	"github.com/0x716/Turnipin/web"
	"github.com/gin-gonic/gin"
	"log"
)

// InitRouter creates a Gin engine, configures middleware,
// sets up API and static routes, then returns the engine instance.
func InitRouter() *gin.Engine {
	// Set Gin Mode
	switch config.GlobalConfig.Development.Mode {
	case "debug":
		gin.SetMode(gin.DebugMode)
	case "release":
		gin.SetMode(gin.ReleaseMode)
	case "test":
		gin.SetMode(gin.TestMode)
	default:
		log.Fatal("Development mode not supported")
	}

	// Create default Gin engine with logger and recovery middleware
	engine := gin.Default()

	// Create a route group for /api endpoints
	apiGroup := engine.Group("/api")
	// Register API routes to the /api group
	routes.RegisterRoutes(apiGroup)

	// Register static file routes to serve frontend assets
	routes.RegisterStaticRoutes(web.DistDirFS, engine)

	return engine
}
