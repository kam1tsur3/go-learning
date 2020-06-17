package main

import (
	"fmt"
	"net/http"
	"os"
	"time"
)

func main() {
	fmt.Fprintf(os.Stdout, "Write with at %v", time.Now())
	/*
		encoder := json.NewEncoder(os.Stdout)
		encoder.SetIndent("", "	")
		encoder.Encode(map[string]string{
			"example": "encoding/json",
			"hello": "world",
		})
	*/
	request, err := http.NewRequest("GET", "http://ascii.jp", nil)
	if err != nil {
		panic(err)
	}
	request.Header.Set("X-TEST", "you can also add header")
	request.Write(os.Stdout)
}
