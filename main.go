package main

import (
	"fmt"
	"strings"

	"github.com/go-resty/resty/v2"
	"github.com/inancgumus/screen"
)

var client = resty.New()

var xssPayload = "<script>alert(1)</script>"

var sqliPayloads = []string{
	"'",
	"' OR '1'='1",
	"'--",
}

var sqliErrors = []string{
	"SQL syntax",
	"mysql_fetch",
	"ORA-01756",
	"syntax error",
}

func main() {
	url := MenuStart()
	clear()

	Run(url)
}

func CheckHeaders(url string) {
	resp, err := client.R().Get(url)

	if err != nil {
		return
	}

	headers := resp.Header()

	if headers.Get("X-Frame-Options") == "" {
		println("[!] Falta X-Frame-Options")
	}

	if headers.Get("Content-Security-Policy") == "" {
		println("[!] Falta CSP")
	}

	if headers.Get("X-XSS-Protection") == "" {
		println("[!] Falta X-XSS-Protection")
	}
}

func TestSql(baseURL string) {
	for _, payload := range sqliPayloads {
		target := baseURL + payload

		body, _, err := MakeRequest(target)
		if err != nil {
			continue
		}

		for _, errMsg := range sqliErrors {
			if strings.Contains(strings.ToLower(body), strings.ToLower(errMsg)) {
				println("[!] Possivel Sql Injection")
			}
		}
	}
}

func TestXSS(baseURL string) {
	target := baseURL + xssPayload

	body, _, err := MakeRequest(target)
	if err != nil {
		return
	}

	if strings.Contains(body, xssPayload) {
		println("[! Possivel XSS:", target)
	}
}

func Run(url string) {
	println("[*] Iniciando scan em:", url)

	TestSql(url)
	TestXSS(url)
	CheckHeaders(url)
}

func clear() {
	screen.Clear()
	screen.MoveTopLeft()
}

func MakeRequest(url string) (string, int, error) {

	resp, err := client.R().Get(url)
	if err != nil {
		return "", 0, err
	}

	body := string(resp.Body())

	return body, resp.StatusCode(), nil
}

func MenuStart() string {
	println("Digite a url\nExemplo: http://site.com/page?id=")
	var url string
	fmt.Scanln(&url)

	return url
}
