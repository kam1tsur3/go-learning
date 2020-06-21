package main

import (
	"archive/zip"
	"io"
	"os"
	"strings"
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

	/* q3.2
	file, _ := os.Create("q32.txt")
	io.CopyN(file, rand.Reader, 1024)
	*/

	/* q3.3 */
	file, _ := os.Create("q33.zip")
	zipWriter := zip.NewWriter(file)
	f1, _ := zipWriter.Create("file1.txt")
	io.WriteString(f1, "in file1\n")
	f2, _ := zipWriter.Create("file2.txt")
	s_rdr := strings.NewReader("in file2\n")
	io.Copy(f2, s_rdr)

	zipWriter.Close()

}
