package main

import (
	"io"
	"log"
	"net"
	"os"
	"strings"
)

type multiTimezoneReader struct {
	connections []net.Conn
	names       []string
}

func (r multiTimezoneReader) Read(p []byte) (n int, err error) {
	slices := make([][]byte, len(r.connections))
	for i := 0; i < len(r.connections); i++ {
		prefix := []byte("\t\t" + r.names[i] + ": ")
		slices[i] = make([]byte, 20) // 適当なサイズでバッファを作成
		nbytes, err := r.connections[i].Read(slices[i])
		if err != nil {
			return 0, err
		}
		// 新しくConnから時刻を受け取った時のみprefixをつける
		if nbytes > 0 {
			slices[i] = append(prefix, []byte(strings.Replace(string(slices[i]), "\n", "", 1))...)
		}
	}
	src := []byte("\r")
	for _, slice := range slices {
		src = append(src, slice...)
	}
	n = copy(p, src)
	return
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}

func main() {
	tzWithPorts := os.Args[1:]
	timezones := make([]string, 0)
	ports := make([]string, 0)

	for _, arg := range tzWithPorts {
		splitted := strings.Split(arg, "=")
		timezones = append(timezones, splitted[0])
		ports = append(ports, splitted[1])
	}

	connections := make([]net.Conn, 0)
	for _, port := range ports {
		conn, err := net.Dial("tcp", port)
		if err != nil {
			log.Fatal(err)
		}
		connections = append(connections, conn)
	}
	r := multiTimezoneReader{connections, timezones}
	mustCopy(os.Stdout, r)
}
