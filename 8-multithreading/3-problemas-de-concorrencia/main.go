package main

import (
	"fmt"
	"net/http"
	"sync/atomic"
	"time"
)

var number int64

func main() {
	//m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		//m.Lock()
		atomic.AddInt64(&number, 1)
		//m.Unlock()
		time.Sleep(300 * time.Millisecond)
		w.Write([]byte(fmt.Sprintf("Você é o visitante numero: %d\n", number)))
	})
	http.ListenAndServe(":3000", nil)
}
