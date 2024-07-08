package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/go-sql-driver/mysql"

	bodyTypes "zap/internal/dict/body_types/search"
	brands "zap/internal/dict/brands/search"
	generations "zap/internal/dict/generations/search"
	models "zap/internal/dict/models/search"
	modifications "zap/internal/dict/modifications/search"
	years "zap/internal/dict/years/search"
	zapsEndpoints "zap/internal/zaps/data/endpoints"
	zapsRepo "zap/internal/zaps/data/repo"
)

var zapDB *sql.DB
var dictDB *sql.DB

func main() {
	var err error

	dictConfig := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "dictdb",
		AllowNativePasswords: true,
	}
	dictDB, err = sql.Open("mysql", dictConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer dictDB.Close()

	zapConfig := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "zapdb",
		AllowNativePasswords: true,
	}
	zapDB, err = sql.Open("mysql", zapConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}
	defer zapDB.Close()

	repo := zapsRepo.ZapsRepo{DB: zapDB}
	zapsEndpoints := zapsEndpoints.ZapsEndpoints{}
	zapsEndpoints.Repo = &repo
	router := gin.Default()
	v1 := router.Group("/api/v1")

	dict := v1.Group("/dict")
	{
		dict.GET("/brands", brands.Search(dictDB))
		dict.GET("/models", models.Search(dictDB))
		dict.GET("/generations", generations.Search(dictDB))
		dict.GET("/body_types", bodyTypes.Search(dictDB))
		dict.GET("/modifications", modifications.Search(dictDB))
		dict.GET("/years", years.Search(dictDB))
	}
	zaps := v1.Group("/zaps")
	{
		zaps.POST("/create", zapsEndpoints.Create())
	}

	router.Run("localhost:8080")

}

// type Album struct {
// 	ID     string  `json:"id"`
// 	Title  string  `json:"title"`
// 	Artist string  `json:"artist"`
// 	Price  float64 `json:"price"`
// }

// var albums = []Album{
// 	{ID: "1", Title: "Blue Train", Artist: "John Coltrane", Price: 56.99},
// 	{ID: "2", Title: "Jeru", Artist: "Gerry Mulligan", Price: 17.99},
// 	{ID: "3", Title: "Sarah Vaughan and Clifford Brown", Artist: "Sarah Vaughan", Price: 39.99},
// }

// func getAlbums(c *gin.Context) {
// 	c.IndentedJSON(http.StatusOK, albums)
// }

// func postAlbums(c *gin.Context) {
// 	var newAlbum Album

// 	// Call BindJSON to bind the received JSON to
// 	// newAlbum.
// 	if err := c.BindJSON(&newAlbum); err != nil {
// 		return
// 	}

// 	// Add the new album to the slice.
// 	albums = append(albums, newAlbum)
// 	c.IndentedJSON(http.StatusCreated, newAlbum)
// }

// func getAlbumByID(c *gin.Context) {
// 	id := c.Param("id")

// 	var alb Album

// 	row := db.QueryRow("SELECT * FROM album WHERE id = ?", id)
// 	if err := row.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
// 		if err == sql.ErrNoRows {
// 			c.IndentedJSON(http.StatusNotFound, gin.H{"message": "albumsById: no such album"})
// 		}
// 	}
// 	c.IndentedJSON(http.StatusOK, alb)
// }

// func albumsByArtist(name string) ([]Album, error) {
// 	// An albums slice to hold data from returned rows.
// 	var albums []Album

// 	rows, err := db.Query("SELECT * FROM album WHERE artist = ?", name)
// 	if err != nil {
// 		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
// 	}
// 	defer rows.Close()
// 	// Loop through rows, using Scan to assign column data to struct fields.
// 	for rows.Next() {
// 		var alb Album
// 		if err := rows.Scan(&alb.ID, &alb.Title, &alb.Artist, &alb.Price); err != nil {
// 			return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
// 		}
// 		albums = append(albums, alb)
// 	}
// 	if err := rows.Err(); err != nil {
// 		return nil, fmt.Errorf("albumsByArtist %q: %v", name, err)
// 	}
// 	return albums, nil
// }
