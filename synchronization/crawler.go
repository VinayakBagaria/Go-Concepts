package synchronization

import (
	"fmt"
	"io"
	"net/http"
	"strings"
)

const word = "the"

var (
	urls = []string{
		"https://code-pilot.me/how-ive-made-this-platform",
		"https://code-pilot.me/making-a-beautiful-error-handler-in-go",
		"https://code-pilot.me/mastering-goroutines-and-channels",
		"https://code-pilot.me/why-should-you-curry",
		"https://code-pilot.me/not-a-real-page",
		"not@validURL",
	}
	resChan  = make(chan int)
	errChan  = make(chan error)
	doneChan = make(chan int, 1)
)

func get(url string) (string, error) {
	res, err := http.Get(url)
	if err != nil {
		return "", err
	}

	if res.StatusCode == 404 {
		return "", fmt.Errorf("not found: %s", url)
	}

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(body), nil
}

func CrawlerImplementation() {
	for _, url := range urls {
		go func(url string) {
			body, err := get(url)
			if err != nil {
				errChan <- err
				return
			}

			y := strings.Count(strings.ToLower(body), word)
			fmt.Printf("found %d occurrences of the word \"%s\" in %s\n", y, word, url)
			resChan <- y
		}(url)
	}

	sum := 0
	ans := 0

	for {
		select {
		case x := <-resChan:
			sum += x
			doneChan <- 1
		case err := <-errChan:
			fmt.Println(err.Error())
			doneChan <- 1
		case <-doneChan:
			ans++
			if ans >= len(urls) {
				fmt.Printf("\nfound %d occurrences of the word \"%s\" in total\n", sum, word)
				close(resChan)
				close(errChan)
			}
		}
	}

}
