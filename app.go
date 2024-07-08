package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	dict_repo "zap/internal/dict/data/repo"
	dict_db "zap/internal/dict/database"
	dict_handlers "zap/internal/dict/handlers"
	zap_db "zap/internal/zaps/data/database"
	zaps_repo "zap/internal/zaps/data/repo"
	zap_handlers "zap/internal/zaps/handlers"
)

var dictDB *sql.DB

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	dictDB = dict_db.DictDB()
	defer dictDB.Close()

	dictHandlers := dict_handlers.Handlers{
		Router: v1.Group("/dict"),
		Repo:   &dict_repo.Repo{DB: dictDB},
	}
	dictHandlers.Init()

	zapDB := zap_db.ZapDB()
	defer zapDB.Close()
	zapHandlers := zap_handlers.Handlers{
		Router: v1.Group("/zaps"),
		Repo:   &zaps_repo.ZapsRepo{DB: zapDB},
	}
	zapHandlers.Init()

	router.Run("localhost:8080")

}
