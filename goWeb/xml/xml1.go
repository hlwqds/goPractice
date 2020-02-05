package main
import (
	"fmt"
	"encoding/xml"
	"io/ioutil"
	"os"
)

type Post struct {
	XMLName xml.Name `xml:"post"`
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author `xml:"author"`
	Comments []Comment `xml:"comments>comment"`
	Xml string `xml:",innerxml"`
}

type Comment struct {
	Id string `xml:"id,attr"`
	Content string `xml:"content"`
	Author Author `xml:"author"`
}

type Author struct {
	Id string `xml:"id,attr"`
	Name string `xml:",chardata"`
}

func main(){
	xmlFile, err := os.Open("post.xml")
	if err != nil{
		fmt.Println("Error opening XML file:", err)
		return
	}
	defer xmlFile.Close()

	xmlData, err := ioutil.ReadAll(xmlFile)
	if err != nil{
		fmt.Println("Error reading XML data:", err)
		return
	}

	var post Post
	xml.Unmarshal(xmlData, &post)
	fmt.Println(post)
}