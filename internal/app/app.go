package app

import (
	"github.com/rusvlp/JWT/internal/delivery/gin"
	"github.com/rusvlp/JWT/pkg/database"
)

func Run() {
	db := database.SQLite()
	eng := gin.SetupRouters(db)
	eng.Run(":7328")
}
