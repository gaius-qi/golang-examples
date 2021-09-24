/*
   File that demostrates file descriptor leak
   in case the client forgets to close response body
*/

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"path"
	"sync"
)

const CLOSE = false

func main() {
	wg := &sync.WaitGroup{}

	for i := 0; i < 100; i++ {
		wg.Add(1)

		go func() {
			defer wg.Done()

			resp, err := http.Get("https://github.com/")
			if err != nil {
				fmt.Println(err)
				return
			}

			if CLOSE {
				resp.Body.Close()
			}
		}()
	}

	wg.Wait()

	pid := fmt.Sprint(os.Getpid())
	if files, err := ioutil.ReadDir(path.Join("/proc", pid, "/fd")); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(len(files))
	}
}
