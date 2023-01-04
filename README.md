# qcloud cdn flusher

Readme `sitemap.xml` from a website and flush cdn cache in `N` days.


## Usage

1. Config: go to [config.yml](config.yml)

```yaml
# sample.yml

sitemap: https://tangx.in/sitemap.xml
purge: true
push: false
lastModInDays: 7
```

2. Run 

```bash
export CDN_ACCESS_ID=akid_12123
export CDN_ACCESS_KEY=akey_asdfasdf

./qcloud-cdn-flusher -c config.yml
```


## Flags

```go
type Flag struct {
	SiteMap       string   `flag:"sitemap" usage:"SiteMap 文件"`
	Purge         bool     `flag:"purge" usage:"刷新URL"`
	Push          bool     `flag:"push" usage:"预热URL"`
	Include       []string `flag:"include" usage:"包含"`
	Exclude       []string `flag:"exclude" usage:"排除"`
	LastModInDays int      `flag:"lastModInDays" usage:"几天之内修改的"`

	Config string `flag:"config" shorthand:"c" usage:"config file, 最高优先级， 可能覆盖其他参数"`
}
```


## Github Action

1. Snippet from Github Action workflow 

```yaml

      - name: Push
        run: |
          make push
        env:
          CDN_ACCESS_ID: ${{ secrets.CDN_ACCESS_ID }}
          CDN_ACCESS_KEY: ${{ secrets.CDN_ACCESS_KEY }}
```

## Release

Download [Latest Version](https://github.com/tangx/qcloud-cdn-flusher/releases/latest/download/qcloud-cdn-flusher-linux-amd64)

