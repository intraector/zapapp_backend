package main

import (
	"database/sql"

	"github.com/gin-gonic/gin"

	bodyTypes "zap/internal/dict/body_types/search"
	dict_repo "zap/internal/dict/data/repo"
	dict_db "zap/internal/dict/database"
	generations "zap/internal/dict/generations/search"
	dict_handlers "zap/internal/dict/handlers"
	models "zap/internal/dict/models/search"
	modifications "zap/internal/dict/modifications/search"
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

	dict := v1.Group("/dict")
	{
		dict.GET("/models", models.Search(dictDB))
		dict.GET("/generations", generations.Search(dictDB))
		dict.GET("/body_types", bodyTypes.Search(dictDB))
		dict.GET("/modifications", modifications.Search(dictDB))
	}

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
