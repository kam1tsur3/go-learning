package main

import (
	"io"
	"net/http"
	//	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "http.ResponsWriter sample")
}

func main() {
	/*
		conn, err := net.Dial("tcp", "ascii.jp:80")
		req, err := http.NewRequest("GET", "http://ascii.jp", nil)
		if err != nil {
			panic(err)
		}

		//	io.WriteString(conn, "GET / HTTP/1.0\r\nHost: ascii.jp\r\n\r\n")
		req.Write(conn)
		io.Copy(os.Stdout, conn)
	*/

	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)

}
