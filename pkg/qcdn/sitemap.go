package qcdn

import (
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"strings"
)

type SiteMap struct {
	XMLName xml.Name `xml:"urlset"`
	URLs    []URL    `xml:"url"`
}

type URL struct {
	XMLName xml.Name `xml:"url"`
	Loc     string   `xml:"loc"`
	LastMod string   `xml:"lastmod,omitempty"`
}

func mustURLs(flag *Flag) []string {

	data, err := readSitemap(flag.SiteMap)
	if err != nil {
		panic(err)
	}

	sitemap := &SiteMap{}
	err = xml.Unmarshal(data, sitemap)
	if err != nil {
		panic(err)
	}

	urls := make([]string, len(sitemap.URLs))
	for i, u := range sitemap.URLs {
		urls[i] = u.Loc
	}

	return urls
}

func readSitemap(name string) ([]byte, error) {
	if strings.HasPrefix(name, "http") {
		return readlink(name)
	}

	return readfile(name)
}

func readfile(name string) ([]byte, error) {
	return os.ReadFile(name)
}

func readlink(name string) ([]byte, error) {

	resp, err := http.Get(name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
