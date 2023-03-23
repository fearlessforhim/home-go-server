package api

import (
	"net/http"
	"encoding/json"
	"log"
	"database/sql"
	_ "modernc.org/sqlite"
)

type BlogPost struct {
	Id int		`json:"id"`
	Title string	`json:"title"`
	Content string 	`json:"content"`
	Timestamp int64 `json:"createdTime"`
	Enabled bool `json:"enabled"`
}

func FetchPosts(w http.ResponseWriter, r *http.Request) {
	var posts []BlogPost
	
	db, err := sql.Open("sqlite", "./blog.db")
	defer db.Close()

	rows, err := db.Query("SELECT * FROM BlogPost WHERE enabled = 1")
	defer rows.Close()

	for rows.Next() {
		var id int
		var title string
		var content string
		var createdTime int64
		var enabled int

		err := rows.Scan(&id, &title, &content, &createdTime, &enabled)
		if err != nil {
			log.Fatal(err)
		}

		post := BlogPost{Id: id, Title: title, Content: content, Timestamp: createdTime}
	        posts = append(posts, post)

	}

	j, err := json.Marshal(posts)

	if err != nil {
                w.Write([]byte("An error occurred " + err.Error()))
                return
        }

	w.Write([]byte(string(j)))
}
