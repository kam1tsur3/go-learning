package main

import (
	"compress/gzip"
	"encoding/json"
	"io"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Encoding", "gzip")
	w.Header().Set("Content-Type", "application/json")
	source := map[string]string{
		"Hello": "World",
	}
	g_writer := gzip.NewWriter(w)
	g_writer.Header.Name = "test.json"
	m_writer := io.MultiWriter(os.Stdout, g_writer)
	encoder := json.NewEncoder(m_writer)
	encoder.SetIndent("", "    ")
	encoder.Encode(source)
	g_writer.Close()
}

func main() {
	/* Q 2.1
	a := 3
	b := "KM2"
	file, err := os.Create("Q2.1.txt")
	if err != nil {
		panic(err)
	}n
	fmt.Fprintf(file, "%s, %d", b, a)
	fmt.Fprintf(os.Stdout, "%s, %d", b, a)
	*/
	/* Q 2.2

	records := [][]string{
		{"pokemon", "fuck", "kame"},
		{"fuck", "unko", "gep"},
	}

	file, e := os.Create("Q2-2.csv")
	if e != nil {
		panic(e)
	}

	// w := csv.NewWriter(os.Stdout)
	w := csv.NewWriter(file)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			panic(err)
		}
	}
	w.Flush()
	*/

	/* Q2-3 */
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
