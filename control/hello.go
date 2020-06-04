package control

import "net/http"

func HelloWorld(w http.ResponseWriter, r *http.Request) {

	w.Write([]byte("Hello, welcome to this aplication"))

}
