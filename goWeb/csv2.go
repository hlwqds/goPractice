package main

import (
	"fmt"
	"encoding/csv"
	"os"
	"strconv"
)

type Post struct{
	Id int
	Content string
	Author string
}

func main() {
	csvFile, err := os.Create("posts.csv")
	if err != nil{
		panic(err)
	}

	defer csvFile.Close()

	allPosts := []Post{
		Post{Id:1, Content: "Hello world!", Author: "Lin Huang"},
		Post{Id:2, Content: "Hello world1!", Author: "Lin Huanglin1"},
		Post{Id:3, Content: "Hello world2!\"", Author: "Lin Huanglin2"},
		Post{Id:4, Content: "Hello world3!", Author: "Lin Huanglin3"},
	}

	writer := csv.NewWriter(csvFile)
	for _, post := range allPosts{
		line := []string{strconv.Itoa(post.Id), post.Content, post.Author}
		fmt.Println(line)
		err := writer.Write(line)
		if err != nil{
			panic(err)
		}
	}
	writer.Flush()

	file, err := os.Open("posts.csv")
	if err != nil{
		panic(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1
	record, err := reader.ReadAll()
	if err != nil{
		panic(err)
	}

	var posts []Post
	for _, item := range record{
		id, _ := strconv.ParseInt(item[0], 0, 0)
		post := Post{Id: int(id), Content: item[1], Author: item[2]}
		posts = append(posts, post)
	}

	fmt.Println(posts[0].Id)
	fmt.Println(posts[0].Content)
	fmt.Println(posts[0].Author)

}