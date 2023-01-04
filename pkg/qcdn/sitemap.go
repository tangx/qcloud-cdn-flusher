package qcdn

import (
	"context"
	"encoding/xml"
	"io"
	"net/http"
	"os"
	"strings"
	"time"

	"github.com/go-jarvis/logr"
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

func mustParseURLs(ctx context.Context, flag *Flag) []string {
	log := logr.FromContext(ctx)

	data, err := readSitemap(flag.SiteMap)
	if err != nil {
		panic(err)
	}

	sitemap := &SiteMap{}
	err = xml.Unmarshal(data, sitemap)
	if err != nil {
		panic(err)
	}

	// format := `2023-01-03T18:15:50+08:00`
	// urls := make([]string, len(sitemap.URLs))
	urls := []string{}

	now := time.Now()
	for _, u := range sitemap.URLs {
		if u.LastMod == "" {
			continue
		}

		lastmod, err := time.Parse(time.RFC3339, u.LastMod)
		if err != nil {
			panic(err)
		}

		tt := lastmod.Add(time.Duration(flag.LastModInDays) * 24 * time.Hour)
		if tt.After(now) {
			log.Debug("Add %s", u.Loc)
			urls = append(urls, u.Loc)
		}
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
