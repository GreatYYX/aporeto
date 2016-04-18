package main

import (
	"bytes"
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"regexp"
	"strings"
)

// the finish signal of goroutine
type FinishSingal interface{}

var (
	re           *regexp.Regexp
	finishSignal FinishSingal
	finishChan   chan FinishSingal
)

func init() {
	re = regexp.MustCompile(`([A-Za-z0-9]+)`)
}

func main() {
	// parse arugments
	flag.Usage = func() {
		fmt.Println("Usage:")
		fmt.Println("gosample [-help|-h]")
		fmt.Println("gosample -urls=<url1>[,<url2>[,<url3>[,<url4>]]]")
		// flag.PrintDefaults()
	}
	optUrls := flag.String("urls", "url", "url string")
	flag.Parse()

	// compute word count & write to file with goroutines
	urls := strings.Split(*optUrls, ",")
	workerCount := len(urls)
	finishChan = make(chan FinishSingal, workerCount)
	for idx, url := range urls {
		go worker(idx, url)
	}
	for i := 0; i < workerCount; i++ {
		<- finishChan
	}
}

// get content from url
func getContent(url string) string {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	defer resp.Body.Close()

	contents, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Printf("%s", err)
		os.Exit(1)
	}
	return string(contents)
}

// tokenize sentences and compute word count
func wordCount(s string) map[string]int {
	t := re.FindAllStringSubmatch(s, -1)
	if t == nil {
		return nil
	}

	hash := make(map[string]int)
	for i := 0; i < len(t); i++ {
		word := t[i][1]
		if _, ok := hash[word]; ok {
			hash[word]++
		} else {
			hash[word] = 1
		}
	}

	return hash
}

// write byte array to file
func writeToFile(name string, data []byte) {
	f, err := os.Create(name)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	_, err = f.Write(data)
	if err != nil {
		panic(err)
	}
}

// goroutine (finishSignal will be emitted once finished)
func worker(idx int, url string) {
	var buffer bytes.Buffer
	buffer.WriteString(fmt.Sprintf("url: %[1]s\n", url))
	res := wordCount(getContent(url))
	for k, v := range res {
		buffer.WriteString(fmt.Sprintf("\t%[1]s: %[2]d\n", k, v))
	}
	writeToFile(fmt.Sprintf("url%[1]d.txt", idx + 1), buffer.Bytes())
	finishChan <- finishSignal
}
