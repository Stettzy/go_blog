package main

import (
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	ID      int    `json:"id"`
	Title   string `json:"title"`
	Content string `json:"content"`
}

type Blog struct {
	Posts []Post
}

var blogPosts = []Blog{
	{
		Posts: []Post{
			{ID: 1, Title: "Title 1", Content: "Content 1"},
			{ID: 2, Title: "Title 2", Content: "Content 2"},
			{ID: 3, Title: "Title 3", Content: "Content 3"},
		},
	},
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	fmt.Print(r.URL.Query().Get("id"))

	for i := 0; i < len(blogPosts); i++ {
		for j := 0; j < len(blogPosts[i].Posts); j++ {
			fmt.Fprintf(w, blogPosts[i].Posts[j].Title)
		}
	}
}

func main() {
	http.HandleFunc("/", handleHome)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
