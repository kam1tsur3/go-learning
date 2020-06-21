package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	/* io.Multixxxx
	header := bytes.NewBufferString("---header---\n")
	content := bytes.NewBufferString("fuck fuck")
	footer := bytes.NewBufferString("---footer---\n")

	reader := io.MultiReader(header, content, footer)
	io.Copy(os.Stdout, reader)
	*/

	/* io.Teexxx
	var buffer bytes.Buffer
	reader := bytes.NewBufferString("Example of io.Tee\n")
	teeReader := io.TeeReader(reader, &buffer)
	_, _ = ioutil.ReadAll(teeReader)
	fmt.Println(buffer.String())
	*/

	/* io.Pipe() ??? */
	r, w := io.Pipe()
	fmt.Fprintf(w, "fuck and fuck\n")
	w.Close()
	file, _ := os.Create("pipetest.txt")
	io.Copy(file, r)
}
