package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// Routes is a structure that handles routes and Gin router configuration.
type Routes struct {
	Routes *gin.Engine
}

// SetUpRouter sets up a new Gin router and returns it.
func (r *Routes) SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

// CreateRoutes configures the routes for the application and assigns the router to the Routes structure.
func (r *Routes) CreateRoutes(db interface{}) {
	router := r.SetUpRouter() // Set up a new Gin router

	// Define a route for the root of the server
	router.GET("/", r.Print)

	//add swagger
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Configure specific tasks routes using the TasksRoutes method
	r.TasksRoutes(router, db)

	// Assign the configured router to the Routes structure
	r.Routes = router
}

// Run starts the HTTP server on port 8080 using the configured router.
func (r *Routes) Run() {
	r.Routes.Run(":8080")
}

// Print is a handler for the root route that responds with a JSON message.
func (r *Routes) Print(gc *gin.Context) {
	gc.JSON(http.StatusOK, gin.H{"message": "Welcome to my API with Golang"})
}
