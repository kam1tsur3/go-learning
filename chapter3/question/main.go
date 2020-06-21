package main

import (
	"crypto/rand"
	"io"
	"os"
)

func main() {
	/* q3.1
	def_s := "default"
	newf := flag.String("filename", def_s, "filename flag")

	flag.Parse()

	if *newf == def_s {
		fmt.Println("Usage: $./q31 --filename=<filename>\n")
		os.Exit(1)
	}

	oldf, _ := os.Open("old.txt")
	writer, _ := os.Create(*newf)
	io.Copy(writer, oldf)
	*/

	/* q3.2 */
	file, _ := os.Create("q32.txt")
	io.CopyN(file, rand.Reader, 1024)
}
