package gin

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/rusvlp/JWT/internal/delivery/gin/handlers"
	"github.com/rusvlp/JWT/internal/repository"
	"github.com/rusvlp/JWT/internal/usecase"
)

func SetupRouters(db *sql.DB) *gin.Engine {
	router := gin.Default()

	rep := repository.NewUserRepository(db)
	useCase := *usecase.NewUserUseCase(rep)
	handler := handlers.UserHandler{UseCase: useCase}

	api := router.Group("/v1")
	{
		api.POST("/reg", handler.Register)
		api.POST("/login", handler.Login)

		api.GET("/users", handler.GetAll)
		api.GET("/user/email/:email", handler.GetUserByEmail)
		api.GET("/user/:id", handler.GetUserByID)

		api.DELETE("/user/:id", handler.DeleteUser)
	}

	auth := router.Group("/profile")
	auth.Use(handlers.Authorization())
	{

	}
	return router
}
