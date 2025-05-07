package wgexample

import (
	"fmt"
	"net/http"
	"sync"
)

var urls []string = []string{
	"https://id.wikipedia.org/wiki/Halaman_Utama",
	"https://id.wikipedia.org/wiki/Halaman_Utama",
	"https://id.wikipedia.org/wiki/Halaman_Utama",
	"https://id.wikipedia.org/wiki/Halaman_Utama",
}

func fetch(client *http.Client, url string) error {
	resp, err := client.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()
	fmt.Println("FINISHED")

	return nil
}

func Run() {
	client := http.Client{}
	var wg sync.WaitGroup

	for _, url := range urls {
		wg.Add(1)
		go func() {
			defer wg.Done()
			fetch(&client, url)
		}()
	}
	wg.Wait()
	fmt.Println("Goroutine Finished")
}
