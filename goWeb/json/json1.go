package main

import (
	"encoding/json"
	"fmt"
	"os"
	"io/ioutil"
)

type Post struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author Author `json:"author"`
	Comments []Comment `json:"comments"`
}

type Comment struct {
	Id int `json:"id"`
	Content string `json:"content"`
	Author string `json:"author"`
}

type Author struct {
	Id int `json:"id"`
	Name string `json:"name"`
}

func decode(fileName string)(post Post, err error){
	jsonFile, err := os.Open(fileName)
	if err != nil{
		fmt.Println("Open err", err)
		return
	}
	defer jsonFile.Close()

	jsonData, err := ioutil.ReadAll(jsonFile)
	if err != nil{
		fmt.Println("ReadAll err", err)
		return
	}

	err = json.Unmarshal(jsonData, &post)
	if err != nil{
		fmt.Println("Unmarshal err", err)
		return
	}

	fmt.Println(post)
	return
}

func main(){
	_, err := decode("post.json")
	if err != nil{
		fmt.Println("decode error!", err)
	}
}