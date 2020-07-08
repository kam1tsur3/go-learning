package main

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"net/http"
	"net/http/httputil"
	"strings"
)

var contents = []string{
	"夜の11時くらいですかね、突然堀江さんから「ひろゆきお前暇か？」って電話がかかってきたんですよね。",
	"「いや寝る所だったんですけど」って答えたら「今可愛い子とホテルに行く所なんだけどお前も来てくれ」って言うんですよ。",
	"それで、ちょっと期待しながらホテル行ってみたら、結構可愛い子で僕もちょっとやる気になっちゃったんですよね。",
	"でもシャワー浴びて準備してベッドに行ってみたらなんとその子チンコ付いてたんですよ！（笑）",
	"「堀江さん、この子男じゃないですか！」って文句言ったら「可愛ければいいじゃねぇか！」って怒るんですよ。",
	"でもまあ堀江さんの言うことも一理あるなって。ニューハーフのお尻に入れるのもまた経験だなと思ってたら、堀江さんが「なあひろゆき、俺に入れて貰うことできるか」って言うんですよ！できる訳ないじゃないですか！",
	"「前と後ろ両方から責めて欲しい。前を触るのは流石にお前でも嫌がると思った」って。",
	"なんなんだよその心遣い。（笑）",
	"普通は中年のおっさんのチンコ触りたくないしお尻にも入れたくないですよ。",
	"結局僕はニューハーフの子にフェラチオだけしてもらって帰ってきました。",
	"もう二度と行きません。（笑）",
}

func isGZipAcceptable(request *http.Request) bool {
	return strings.Index(
		strings.Join(request.Header["Accept-Encoding"], ","), "gzip") != -1
}

func processSession(conn net.Conn) {
	fmt.Printf("Accept %v\n", conn.RemoteAddr())
	defer conn.Close()
	for {
		request, err := http.ReadRequest(bufio.NewReader(conn))
		if err != nil {
			if err == io.EOF {
				break
			}
			panic(err)
		}
		dump, err := httputil.DumpRequest(request, true)
		if err != nil {
			panic(err)
		}
		fmt.Println(string(dump))

		fmt.Fprintf(conn, strings.Join([]string{
			"HTTP/1.1 200 OK",
			"Content-Type: text/plain",
			"Transfer-Encoding: chunked",
			"", "",
		}, "\r\n"))
		for _, content := range contents {
			bytes := []byte(content)
			fmt.Fprintf(conn, "%x\r\n%s\r\n", len(bytes), content)
		}
		fmt.Fprintf(conn, "0\r\n\r\n")
	}
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8888")
	if err != nil {
		panic(err)
	}
	fmt.Println("Server is running at localhost:8888")
	for {
		conn, err := listener.Accept()
		if err != nil {
			panic(err)
		}
		go processSession(conn)
	}
}
