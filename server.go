package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http/pprof"
	"runtime"
	"sync/atomic"
	"time"
)

var (
	c, max uint64
)

func pong(resp http.ResponseWriter, req *http.Request) {
	//r := req.Method + req.URL.Path
	//log.Println(r)
	resp.Write([]byte("pong\n"))
	atomic.AddUint64(&c, 1)
}

func echo(resp http.ResponseWriter, req *http.Request) {
	buf, err := ioutil.ReadAll(req.Body)
	if err != nil {
		resp.WriteHeader(http.StatusBadRequest)
	}
	resp.Write(buf)
	atomic.AddUint64(&c, 1)
}

func metrics(resp http.ResponseWriter, req *http.Request) {
	resp.Write([]byte(fmt.Sprintf("Top performance: %d req/sec", atomic.LoadUint64(&max))))
}

func main() {

	http.HandleFunc("/ping", pong)
	http.HandleFunc("/echo", echo)
	http.HandleFunc("/debug/pprof/metrics", metrics)
	go func() {
		for {
			time.Sleep(time.Millisecond * 200)
			cnt := atomic.SwapUint64(&c, 0) * 5
			if cnt > atomic.LoadUint64(&max) {
				atomic.StoreUint64(&max, cnt)
			}
		}
	}()

	log.Printf("Ping/Pong server started at: 0.0.0.0:8080 on %d thread(s)\n", runtime.NumCPU())
	log.Println(http.ListenAndServe("0.0.0.0:8080", nil))
}
