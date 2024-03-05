package main

import (
	"encoding/base64"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/gocolly/colly"
)

type Downloader struct {
	CookieJarPath string
	CookieJar     http.CookieJar
	AsfUrs4       map[string]string
	Context       *colly.Context
}

func NewDownloader() *Downloader {
	cookieJarPath := filepath.Join(os.ExpandUser("~"), ".bulk_download_cookiejar.txt")
	return &Downloader{
		CookieJarPath: cookieJarPath,
		CookieJar:     nil,
		AsfUrs4: map[string]string{
			"url":    "https://urs.earthdata.nasa.gov/oauth/authorize",
			"client": "BO_n7nTIlMljdvU6kRRB3g",
			"redir":  "https://auth.asf.alaska.edu/login",
		},
		Context: colly.NewContext(),
	}
}

func (d *Downloader) GetCookie() {
	parsedURL, _ := url.Parse(d.AsfUrs4["url"])
	for !d.CheckCookie(parsedURL) {
		d.GetNewCookie(parsedURL)
	}
}

func (d *Downloader) CheckCookie(parsedURL *url.URL) bool {
	if d.CookieJar == nil {
		fmt.Printf(" > Cookiejar is bunk: %v\n", d.CookieJar)
		return false
	}

	fileCheck := "https://urs.earthdata.nasa.gov/profile"
	c := colly.NewCollector()
	c.WithTransport(&http.Transport{TLSClientConfig: d.Context.TLSClientConfig})
	c.SetCookies(parsedURL, d.CookieJar.Cookies(parsedURL))

	request, err := http.NewRequest("HEAD", fileCheck, nil)
	if err != nil {
		log.Fatal(err)
	}
	timeout := time.Duration(30 * time.Second)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: d.Context.TLSClientConfig,
		},
		Timeout: timeout,
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	respCode := response.StatusCode

	if !d.CheckCookieIsLoggedIn() {
		return false
	}

	err = d.CookieJar.SaveCookies(parsedURL, d.CookieJar.Cookies(parsedURL))
	if err != nil {
		log.Fatal(err)
	}

	if respCode >= 300 && respCode <= 303 {
		redirURL, _ := response.Location()
		if strings.Contains(redirURL.String(), "vertex-retired.daac.asf.alaska.edu") && strings.Contains(d.AsfUrs4["redir"], "test") {
			fmt.Println("Cough, cough. It's dusty in this test env!")
			return true
		}
		return false
	}

	if respCode == 200 || respCode == 307 {
		return true
	}

	return false
}

func (d *Downloader) GetNewCookie(parsedURL *url.URL) {
	authCookieURL := d.AsfUrs4["url"] + "?client_id=" + d.AsfUrs4["client"] + "&redirect_uri=" + d.AsfUrs4["redir"] + "&response_type=code&state="
	userPass := base64.StdEncoding.EncodeToString([]byte("mocheer:gulaxy@1107G"))

	d.CookieJar, _ = cookiejar.New(nil)
	c := colly.NewCollector()
	c.WithTransport(&http.Transport{TLSClientConfig: d.Context.TLSClientConfig})

	request, err := http.NewRequest("GET", authCookieURL, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Set("Authorization", "Basic "+userPass)

	timeout := time.Duration(30 * time.Second)
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: d.Context.TLSClientConfig,
		},
		Timeout: timeout,
	}
	response, err := client.Do(request)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()

	if d.CheckCookieIsLoggedIn() {
		err = d.CookieJar.SaveCookies(parsedURL, d.CookieJar.Cookies(parsedURL))
		if err != nil {
			log.Fatal(err)
		}
		return
	log.Fatal("登录失败")
}
}
func (d *Downloader) CheckCookieLoggedIn() bool {
	if d.CookieJar == nil {
			return false
	}
	parsedURL, _ := url.Parse("https://urs.earthdata.nasa/profile")
	cookies := d.CookieJar.Cookies(parsedURL)
 _, cookie := range cookies {
		if cookie.Name == "urs_user_already_logged" {
			return true
		}
	}
	return false
}

func main() {
	downloader := NewDownloader()
	downloader.GetCookie()
}