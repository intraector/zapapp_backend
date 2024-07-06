package main

import (
	"database/sql"
	"log"

	"github.com/gin-gonic/gin"

	"github.com/go-sql-driver/mysql"

	generation "zap/internal/dict/generation/search"
	mark "zap/internal/dict/mark/search"
	model "zap/internal/dict/model/search"
)

var db *sql.DB
var dictDB *sql.DB

func main() {

	// cfg := mysql.Config{
	// 	User:                 "root",
	// 	Passwd:               "",
	// 	Net:                  "tcp",
	// 	Addr:                 "localhost:3306",
	// 	DBName:               "zapdb",
	// 	AllowNativePasswords: true,
	// }
	dictConfig := mysql.Config{
		User:                 "root",
		Passwd:               "",
		Net:                  "tcp",
		Addr:                 "localhost:3306",
		DBName:               "dictdb",
		AllowNativePasswords: true,
	}

	var err error

	dictDB, err = sql.Open("mysql", dictConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	// db, err = sql.Open("mysql", cfg.FormatDSN())
	// if err != nil {
	// 	log.Fatal(err)
	// }

	// res, err := dictDB.Query("SHOW TABLES")

	// if err != nil {
	// 	log.Fatal(err)
	// }

	// albums, err := albumsByArtist("John Coltrane")
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Printf("Albums found: %v\n", albums)

	// var table string

	// for res.Next() {
	// 	res.Scan(&table)
	// 	fmt.Println(table)
	// }

	router := gin.Default()
	router.GET("/marks", mark.Search(dictDB))
	router.GET("/models", model.Search(dictDB))
	router.GET("/generations", generation.Search(dictDB))
	// router.GET("/albums", getAlbums)
	// router.GET("/albums/:id", getAlbumByID)
	// router.POST("/albums", postAlbums)

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
