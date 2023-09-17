package main

import (
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"

	"golang.org/x/net/html"
)

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

type flagSet struct {
	depth int
	url   string
}

var outDir string = "downloaded"

var fSet flagSet

func init() {
	flag.IntVar(&fSet.depth, "d", 1, "depth. if file, then 1")
	flag.StringVar(&fSet.url, "u", "https://example.com/", "url. must be set")
}
func resolveURL(baseURL string, href string) (string, error) {
	baseUrl, _ := url.Parse(baseURL)
	hrefUrl, err := url.Parse(href)
	if err != nil {
		return "", err
	}
	baseUrl.ResolveReference(hrefUrl)
	return baseUrl.Path, nil
}
func linksFromBody(baseURL string, body io.Reader) ([]string, error) {
	var links []string

	tokenizer := html.NewTokenizer(body)
	for {
		tokenType := tokenizer.Next()
		if tokenType == html.ErrorToken {
			err := tokenizer.Err()
			if err != nil && err != io.EOF {
				return nil, err
			}
			break
		}

		token := tokenizer.Token()
		if tokenType == html.StartTagToken && token.Data == "a" {
			for _, attr := range token.Attr {
				if attr.Key == "href" {
					linkURL, err := resolveURL(baseURL, attr.Val)
					if err != nil {
						return nil, err
					}
					links = append(links, linkURL)
				}
			}
		}
	}

	return links, nil
}

func recUrlHandler(curUrl string, dep int) error {
	if dep <= 0 {
		return nil
	}
	urlParsed, err := url.Parse(curUrl)
	if err != nil {
		return err
	}
	if urlParsed.Scheme != "http" && urlParsed.Scheme != "https" {
		return fmt.Errorf("Unsupported scheme")
	}
	filePath := filepath.Join(outDir, urlParsed.Host, urlParsed.Path)
	if strings.HasSuffix(urlParsed.Path, "/") {
		filePath = filepath.Join(filePath, "index.html")
	}
	err = os.MkdirAll(filepath.Dir(filePath), os.ModePerm)
	if err != nil {
		return err
	}
	resp, err := http.Get(curUrl)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	curFile, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer curFile.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	_, err = curFile.Write(bytes)
	if err != nil {
		return err
	}
	baseUrl := urlParsed.Scheme + "://" + urlParsed.Host

	resp, err = http.Get(curUrl)
	if err != nil {
		return err
	}
	links, err := linksFromBody(baseUrl, resp.Body)
	if err != nil {
		return err
	}
	for _, link := range links {
		err := recUrlHandler(link, dep-1)
		if err != nil {
			return err
		}
	}
	return nil

}

func myWget() error {
	err := os.MkdirAll(outDir, os.ModePerm)
	if err != nil {
		return fmt.Errorf("Cant create outDir: %s", err)
	}
	return recUrlHandler(fSet.url, fSet.depth)
}

func main() {
	flag.Parse()
	err := myWget()
	if err != nil {
		fmt.Println(err)
	}
}
