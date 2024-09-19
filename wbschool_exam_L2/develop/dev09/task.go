package main

/*
=== Утилита wget ===

Реализовать утилиту wget с возможностью скачивать сайты целиком

Программа должна проходить все тесты. Код должен проходить проверки go vet и golint.
*/

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"path"
	"strings"

	"golang.org/x/net/html"
)

// downloadPage загружает страницу и возвращает содержимое.
func downloadPage(url string) (string, error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return "", fmt.Errorf("не удалось загрузить страницу: %s", resp.Status)
	}

	bodyBytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}

	// Сохраняем страницу
	fileName := "index.html"
	err = os.WriteFile(fileName, bodyBytes, 0644)
	if err != nil {
		return "", err
	}

	return string(bodyBytes), nil
}

// parseHTML загружает все ресурсы и ссылки со страницы.
func parseHTML(baseURL, body string) error {
	tokenizer := html.NewTokenizer(strings.NewReader(body))
	for {
		tt := tokenizer.Next()
		switch tt {
		case html.ErrorToken:
			return nil
		case html.StartTagToken, html.SelfClosingTagToken:
			tagName, _ := tokenizer.TagName()
			if string(tagName) == "a" || string(tagName) == "img" {
				for {
					attrName, attrValue, moreAttr := tokenizer.TagAttr()
					if string(attrName) == "href" || string(attrName) == "src" {
						resourceURL := string(attrValue)
						absoluteURL := resolveURL(baseURL, resourceURL)
						downloadResource(absoluteURL)
					}
					if !moreAttr {
						break
					}
				}
			}
		}
	}
}

// resolveURL преобразует относительные пути в абсолютные.
func resolveURL(baseURL, resourceURL string) string {
	if strings.HasPrefix(resourceURL, "http") {
		return resourceURL
	}
	return fmt.Sprintf("%s/%s", strings.TrimRight(baseURL, "/"), strings.TrimLeft(resourceURL, "/"))
}

// downloadResource загружает ресурс и сохраняет его локально.
func downloadResource(url string) {
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("Ошибка загрузки ресурса:", err)
		return
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		fmt.Printf("Не удалось загрузить ресурс: %s (%s)\n", url, resp.Status)
		return
	}

	fileName := path.Base(resp.Request.URL.Path)
	if fileName == "/" || fileName == "" {
		fileName = "index.html"
	}

	file, err := os.Create(fileName)
	if err != nil {
		fmt.Println("Ошибка создания файла:", err)
		return
	}
	defer file.Close()

	_, err = io.Copy(file, resp.Body)
	if err != nil {
		fmt.Println("Ошибка записи файла:", err)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Использование: wget <URL>")
		return
	}

	url := os.Args[1]

	fmt.Println("Загрузка страницы:", url)

	// Шаг 1: Загрузка основной страницы
	body, err := downloadPage(url)
	if err != nil {
		fmt.Println("Ошибка загрузки страницы:", err)
		return
	}

	// Шаг 2: Парсинг HTML и загрузка ресурсов
	err = parseHTML(url, body)
	if err != nil {
		fmt.Println("Ошибка парсинга HTML:", err)
	}
}
