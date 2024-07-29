package main

import (
	"context"
	"database/sql"

	"github.com/gin-gonic/gin"

	"zap/internal/auth/auth_endpoints"
	"zap/internal/database/auth_db"
	"zap/internal/database/dict_db"
	"zap/internal/database/zap_db"
	"zap/internal/dict/dict_endpoints"
	"zap/internal/dict/dict_repo"
	"zap/internal/zap/zap_endpoints"
)

var dictDB *sql.DB

func main() {
	router := gin.Default()
	v1 := router.Group("/api/v1")

	dictDB = dict_db.New()
	defer dictDB.Close()
	dictEndpoints := dict_endpoints.New(
		v1.Group("/dict"),
		&dict_repo.Repo{DB: dictDB},
	)
	dictEndpoints.Init()

	authDB := auth_db.New()
	defer authDB.Close(context.Background())
	authHandlers := auth_endpoints.New(
		v1.Group("/auth"),
		authDB,
	)
	authHandlers.Init()

	zapDB := zap_db.New()
	defer zapDB.Close(context.Background())
	zapHandlers := zap_endpoints.New(
		v1.Group("/zaps"),
		zapDB,
	)
	zapHandlers.Init()

	router.Run("localhost:8080")

}
