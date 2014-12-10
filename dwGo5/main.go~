package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"runtime"
	"strings"
	"sync"
)

func download(url string, wg *sync.WaitGroup) {
	defer wg.Done()
	tmp := strings.Split(url, "/")
	fileName := tmp[len(tmp)-1]
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	ioutil.WriteFile(fileName, body, 0644)
}
func main() {
	wg := new(sync.WaitGroup)
	runtime.GOMAXPROCS(runtime.NumCPU())
	for _, url := range os.Args[1:] {
		wg.Add(1)
		go download(url, wg)
	}
	wg.Wait()
}
