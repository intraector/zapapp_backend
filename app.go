package main

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"

	dict_repo "zap/internal/dict/data/repo"
	dict_db "zap/internal/dict/database"
	dict_handlers "zap/internal/dict/handlers"
	zap_db "zap/internal/zaps/data/database"
	zap "zap/internal/zaps/handlers"
)

var dictDB *sql.DB

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	dictDB = dict_db.DictDB()
	defer dictDB.Close()

	dictHandlers := dict_handlers.New(
		v1.Group("/dict"),
		&dict_repo.Repo{DB: dictDB},
	)

	dictHandlers.Init()

	zapDB := zap_db.New()
	defer zapDB.Close(context.Background())

	zapHandlers := zap.New(v1.Group("/zaps"), zapDB)
	zapHandlers.Init()

	router.Run("localhost:8080")

}
