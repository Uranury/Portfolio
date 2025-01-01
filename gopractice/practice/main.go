package practice

import (
	"log"
	"net/http"
	"sync"
)

type URLresult struct {
	URL    string `json:"url"`
	Status string `json:"status"`
	Error  string `json:"error"`
}

var Wg sync.WaitGroup

func checkRealURL(url string) (*http.Response, error) {
	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	return resp, nil
}

func Example1() {
	URLs := [...]string{"https://google.com", "https://wsp.kbtu.kz", "https://amazon.com", "https://nhentai.net", "https://amirlox.com"}

	for _, url := range URLs {
		Wg.Add(1)
		go func(url string) {
			defer Wg.Done()
			resp, err := checkRealURL(url)
			if err != nil {
				log.Println(url + " failed with error: " + err.Error())
			} else {
				log.Println(url + " " + resp.Status)
				resp.Body.Close() // Ensure response body is closed
			}
		}(url)
	}

	Wg.Wait()

	// #OUTPUT
	/*
		2024/12/06 11:37:32 https://amirlox.com failed with error: Get "https://amirlox.com": dial tcp: lookup amirlox.com: no such host
		2024/12/06 11:37:32 https://wsp.kbtu.kz 200
		2024/12/06 11:37:32 https://nhentai.net 200 OK
		2024/12/06 11:37:33 https://google.com 200 OK
		2024/12/06 11:37:33 https://amazon.com 200 OK
	*/
}
