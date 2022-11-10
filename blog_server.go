package main

import (
	"fmt"
	"net/http"
	"text/template"
)

func home(w http.ResponseWriter, r *http.Request) {

	t, _ := template.ParseFiles("./blog/home.html")

	data := map[string]string{}

	t.Execute(w, data)

}

func main() {

	http.HandleFunc("/", home)

	http.Handle("/blog/", http.StripPrefix("/blog/", http.FileServer(http.Dir("./blog"))))

	fmt.Println("服务器即将开启")

	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Println("服务器开启错误: ", err)
	}
	fmt.Println("webserver")
}
