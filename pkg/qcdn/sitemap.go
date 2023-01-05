package qcdn

import (
	"context"
	"encoding/xml"
	"fmt"
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

	// 如果 LastModInDays 为 0， 则全部添加
	if flag.LastModInDays == 0 {
		urls := make([]string, len(sitemap.URLs))
		for i, u := range sitemap.URLs {
			urls[i] = u.Loc
		}

		return urls
	}

	// 否则按照最近日期添加
	// format := `2023-01-03T18:15:50+08:00`
	// urls := make([]string, len(sitemap.URLs))
	urls := []string{}

	now := time.Now()
	for _, u := range sitemap.URLs {

		// N 天内
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

	// 跳过缓存
	name = fmt.Sprintf("%s?%d", name, time.Now().Unix())

	resp, err := http.Get(name)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	return io.ReadAll(resp.Body)
}
