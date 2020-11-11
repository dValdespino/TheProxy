package main

import (
	"fmt"
	// "github.com/agtorre/gocolorize"
	"flag"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var (
	colors map[string]func(...interface{}) string = make(map[string]func(...interface{}) string)
)

// func RegisterColor(name, back, front string, underline bool) {
// 	colorize := gocolorize.NewColor(front + ":" + back)

// 	if underline {
// 		colorize.ToggleUnderline()
// 	}

// 	colors[name] = colorize.Paint
// }

// func Colorize(color string, opts ...interface{}) string {
// 	return colors[color](opts)
// }

func Colorize(color string, opts ...interface{}) string {
	var ret string

	for _, opt := range opts {
		ret = fmt.Sprintf("%s%v", ret, opt)
	}

	return ret
}

var (
	proxy_url *url.URL
	proxy     *httputil.ReverseProxy
)

func serveReverseProxy(res http.ResponseWriter, req *http.Request) {
	fmt.Printf("Handling [ %s ] ...\n", Colorize("link", req.URL.String()))

	proxy.ServeHTTP(res, req)
}


func main() {
	uname := flag.String("user", "dvaldespino", "User name")
	passwd := flag.String("passwd", "A32x140p", "User password")
	proxy_addr := flag.String("proxy", "proxy.dpe.cm.rimed.cu", "Proxy addr")
	proxy_port := flag.Int("port", 3128, "Proxy port")
	protocol := flag.String("protocol", "http", "Proxy protocol(http or https)")

	flag.Usage=func(){
		fmt.Println(`
The Proxy
---------

 USAGE INFO:
  TheProxy [-user username] [-passwd password] [-proxy proxy_addr] [-port proxy_port] [-protocol protocol]

 -user User name
 -passwd Password
 -proxy The Proxy's address
 -port The Proxy's port
 -protocol The Proxy's protocol (might be http or http)
 `)		
	}
	flag.Parse()


	full_url := fmt.Sprintf("%s://%s:%s@%s:%d", *protocol, *uname, *passwd, *proxy_addr, *proxy_port)
	public_url := fmt.Sprintf("%s://%s:******@%s:%d", *protocol, *uname, *proxy_addr, *proxy_port)
	fmt.Printf("Using real proxy: '%s'\n", public_url)
	proxy_url, _ = url.Parse(full_url)

	proxy = httputil.NewSingleHostReverseProxy(proxy_url)

	// RegisterColor("link", "blue", "black", false)
	// RegisterColor("title", "blue", "black", false)

	fmt.Println("Waiting for requests...")
	http.HandleFunc("/", serveReverseProxy)

	log.Fatal(http.ListenAndServe("0.0.0.0:3128", nil))
}
