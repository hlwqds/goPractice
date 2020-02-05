package main

import (
	"database/sql"
	"fmt"
	"errors"
	_ "github.com/lib/pq"
)

type Post struct {
	Id int
	Content string
	Author string
	Comments []Comment
}

type Comment struct {
	Id int
	Content string
	Author string
	Post *Post
}

var Db *sql.DB
func init(){
	var err error
	Db, err = sql.Open("postgres", "user=hlwqds dbname=hlwqds password=printf`1 sslmode=disable")
	if err != nil{
		panic(err)
	}

	fmt.Println(Db)
}

func (comment *Comment) Create() (err error){
	if comment.Post == nil{
		err = errors.New("Post not found")
		return
	}
	err = Db.QueryRow("insert into comments (content, author, post_id) values",
	"($1, $2, $3) returning id", comment.Content, comment.Author, comment.Post.Id).Scan(&comment.Id)
	
	return
}

func Posts(limit int) (posts []Post, err error){
	rows, err := Db.Query("select id, content, author from posts limit $1", limit)
	if err != nil{
		panic(err)
	}
	defer rows.Close()

	for rows.Next(){
		post := Post{}
		err = rows.Scan(&post.Id, &post.Content, &post.Author)
		if err != nil{
			return
		}
		posts = append(posts, post)
	}

	return
}

func GetPost(id int) (post Post, err error){
	post = Post{}
	post.Comments = []Comment{}

	row := Db.QueryRow("select id, content, author from posts where id=$1", id)
	if err != nil{
		panic(err)
	}

	err = row.Scan(&post.Id, &post.Content, &post.Author)
	rows, err := Db.Query("select id, content, author from posts where post_id=$1", post.Id)
	if err != nil{
		return
	}
	defer rows.Close()

	for rows.Next(){
		comment := Comment{Post: &post}
		err = rows.Scan(&comment.Id, &comment.Content, &comment.Author)
		if err != nil{
			return
		}
		post.Comments = append(post.Comments, comment)
	}
	return
}

func (post *Post) Create() (err error){
	statement := "insert into posts (content, author) values ($1, $2) returning id"
	stmt, err := Db.Prepare(statement)
	if err != nil{
		panic(err)
	}
	defer stmt.Close()
	err = stmt.QueryRow(post.Content, post.Author).Scan(&post.Id)
	return
}

func (post *Post) Update() (err error){
	_, err = Db.Exec("update posts set content=$2, author=$3 where id=$1", post.Id,
	post.Content, post.Author)
	return
}

func (post *Post) Delete() (err error){
	_, err = Db.Exec("delete from posts where id=$1", post.Id)
	return
}

func main(){
	post := Post{Content:"Hello World", Author: "Lin Huang"}

	fmt.Println(post)
	post.Create()
	fmt.Println(post)
	readPost, _ := GetPost(post.Id)
	fmt.Println(readPost)

	readPost.Content = "Bonjour Monde!"
	readPost.Author = "Pierre"
	readPost.Update()
	fmt.Println(readPost)

	posts, _ := Posts(10)
	fmt.Println(posts)

	readPost.Delete()
}