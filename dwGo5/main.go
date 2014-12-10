package main

import (
	"fmt"
	"github.com/codegangsta/cli"
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
	app := cli.NewApp()
	app.Name = "download"
	app.Usage = "download"
	app.Action = func(c *cli.Context) {
		wg := new(sync.WaitGroup)
		runtime.GOMAXPROCS(runtime.NumCPU())
		for _, url := range os.Args[1:] {
			wg.Add(1)
			go download(url, wg)
		}
		wg.Wait()
	}
	app.Run(os.Args)
}
