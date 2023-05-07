package mediumremover

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/exec"
	"regexp"
	"strings"
)

func fetchHtml(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		log.Fatalf("Some error occured in fetching: %v", err)
	}

	body, err := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	if err != nil {
		log.Fatalf("Some error occured in decoding: %v", err)
	}

	return string(body)
}

func sanitizeHtml(htmlBody string) string {
	scriptRegex := regexp.MustCompile(`<script src="https://cdn-client.medium.com(.*)script>`)

	sanitized := scriptRegex.ReplaceAllString(htmlBody, "")
	sanitized = strings.Replace(sanitized, "\n", "", -1)
	return strings.TrimSpace(sanitized)
}

func getTitle(htmlBody string) string {
	titleRegex := regexp.MustCompile(`<title(.*)</title>`)
	found := titleRegex.FindAllString(htmlBody, -1)
	if len(found) == 0 {
		return "Some title"
	}

	title := strings.Replace(found[0], "</title>", "", -1)
	splitted := strings.Split(title, ">")[1]
	return splitted
}

func writeToFile(filePrefix string, contents string) string {
	fileName := fmt.Sprintf("%s.html", strings.Replace(filePrefix, " ", "-", -1))
	f, err := os.Create(fileName)
	if err != nil {
		log.Fatalf("cant create file: %v", err)
	}
	defer f.Close()

	_, err = f.WriteString(contents)
	if err != nil {
		log.Fatalf("unable to save to file: %v", err)
	}

	return fileName
}

func openBrowser(url string) {
	err := exec.Command("open", url).Start()
	if err != nil {
		log.Fatalf("unable to open browser: %v", err)
	}
}

func removeFile(filePath string) {
	err := exec.Command("rm", filePath).Start()
	if err != nil {
		log.Fatalf("unable to remove file: %v", err)
	}
}

func DoWork() {
	if len(os.Args) == 1 {
		log.Fatal("Provide url")
	}

	url := os.Args[1]

	htmlBody := fetchHtml(url)

	sanitized := sanitizeHtml(htmlBody)
	title := getTitle(sanitized)
	fileLocation := writeToFile(title, sanitized)

	openBrowser(fileLocation)
}
