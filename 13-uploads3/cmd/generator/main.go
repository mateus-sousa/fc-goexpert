package main

import (
	"fmt"
	"os"
	"time"
)

func main() {
	i := 0
	for {
		time.Sleep(300 * time.Millisecond)
		f, err := os.Create(fmt.Sprintf("./tmp/file-%d.txt", i))
		if err != nil {
			panic(err)
		}
		defer f.Close()
		f.WriteString("HELLO WORLD!")
		i++
	}
}
