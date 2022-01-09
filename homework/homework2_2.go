package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func index(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("<h1>Welcome to cloud native</h1>"))

	for k, v := range r.Header {
		//fmt.Println(k, v)
		for _, vv := range v {
			//写入
			w.Header().Set(k, vv)
		}
	}

	os.Setenv("VERSION", "0.0.1")
	version := os.Getenv("VERSION")
	fmt.Println(version)
	w.Header().Set("VERSION", version)

	//取Client ip
	// clientIP := r.RemoteAddr
	// fmt.Println(clientIP)

	//如果经过负载均衡器，proxy，remoteAddr取到的clientIP是负载均衡器，proxy的地址
	//不是真实用户地址
	//X-REAL-IP 真实用户地址
	//X-FORWORD-FOR 真实用户地址
	clientIP := getCurrentIP(r)
	httpCode := http.StatusOK
	log.Printf("clientIP: %s, status code: %d\n", clientIP, httpCode)

}

func getCurrentIP(r *http.Request) string {
	ip := r.Header.Get("X-REAL-IP")
	if ip == "" {
		//remoteAddr IP:PORT
		//ip = strings.Split(r.RemoteAddr, ":")[0]
		ip = r.RemoteAddr
	}
	return ip
}

func healthz(w http.ResponseWriter, r *http.Request) {
	fmt.Println("I am healthz")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", index)
	mux.HandleFunc("/healthz", healthz)
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal("start server failed, %s\n", err.Error())
	}
}
