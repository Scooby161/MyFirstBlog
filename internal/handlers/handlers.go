package handlers

import "net/http"

func Home(w http.ResponseWriter, r *http.Request) {

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	w.Write([]byte("Hello World"))

}

func ShowSnippet(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World 2"))
}

func CreateSnippet(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		w.Header().Set("Allow", http.MethodPost)
		w.WriteHeader(405)
		w.Write([]byte("close"))
		return
	}
	w.Write([]byte("Hello World 3"))
}
