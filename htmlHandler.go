package main

import (
	"fmt"
	"html/template"
	"os"
	//"strings"
	//	"io/ioutil"
	//"encoding/json"
	"log"
	"net/http"
	"net/url"
	//	"github.com/PuerkitoBio/goquery"
)

func root_begin(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.String())
	url, _ := url.QueryUnescape(r.URL.String())
	//fmt.Println("root_begin url:" + url)

	if len(url) == 0 || url == "/" {
		t, _ := template.ParseFiles("./bootstrap-3.3.7/" + "offcanvas.html")
		t.Execute(w, nil)
		return
	} else {
		fmt.Println("else url:" + url)
	}

}

func router() {
	http.HandleFunc("/", root_begin)
	http.Handle("/offcanvas.js", http.FileServer(http.Dir("./bootstrap-3.3.7/")))
	http.Handle("/offcanvas.css", http.FileServer(http.Dir("./bootstrap-3.3.7/")))
	http.Handle("/dist/js/", http.FileServer(http.Dir("./bootstrap-3.3.7/docs/")))
	http.Handle("/dist/css/", http.FileServer(http.Dir("./bootstrap-3.3.7/docs/")))
	http.Handle("/dist/fonts/", http.FileServer(http.Dir("./bootstrap-3.3.7/docs/")))
	http.Handle("/assets/js/", http.FileServer(http.Dir("./bootstrap-3.3.7/docs/")))
	http.Handle("/assets/css/", http.FileServer(http.Dir("./bootstrap-3.3.7/docs/")))
	http.Handle("/assets/fonts/", http.FileServer(http.Dir("./bootstrap-3.3.7/docs/")))
}

func main() {
	fmt.Println("Server begin")

	port := os.Getenv("PORT")

	if port == "" {
		port = "6060"
	}
	fmt.Println("port: " + port)
	router()

	err := http.ListenAndServe(":"+port, nil) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
