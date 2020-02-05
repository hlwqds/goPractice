package main
import (
//	"fmt"
	"net/http"
	"html/template"
	"math/rand"
	"time"
)

func process(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("tmpl.html")
	rand.Seed(time.Now().Unix())
	t.Execute(w, rand.Intn(10) > 5)
}

func process2(w http.ResponseWriter, r *http.Request){
	t, _ := template.ParseFiles("tmpl2.html")
	daysOfWeek := []string{"Mon", "Tue", "Wed", "Thu", "Fri", "Sat", "Sun"}
	t.Execute(w, daysOfWeek)
}

func formatDate(t time.Time) string{
	layout := "2006-01-02"
	return t.Format(layout)
}

func process3(w http.ResponseWriter, r *http.Request){
	funcMap := template.FuncMap{
		"fdate": formatDate,
	}

	t := template.New("tmpl3.html").Funcs(funcMap)
	t, _ = t.ParseFiles("tmpl3.html")
	t.Execute(w, time.Now())
}

func main(){
	server := http.Server{
		Addr: "127.0.0.1:8080",
	}
	http.HandleFunc("/process", process)
	http.HandleFunc("/process2", process2)
	http.HandleFunc("/process3", process3)

	server.ListenAndServe()
}