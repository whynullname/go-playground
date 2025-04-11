package theory

import (
	"log"
	"net/http"
)

type MyHandler struct{}

func (m MyHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello world"))
}

func CreateBaseNetServer() {
	var h MyHandler
	err := http.ListenAndServe(":8080", h)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's a test handler"))
}

func CreateBaseServerWithDefaultMux() {
	http.HandleFunc("/test", testHandler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalln(err)
		return
	}
}

func mainHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("It's a main handler"))
}

func CreateBaseNetServerWithMyMux() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", mainHandler)
	mux.HandleFunc("/test", testHandler)
	err := http.ListenAndServe(":8080", mux)
	if err != nil {
		log.Fatalln(err)
		return
	}
}
