package qcdn

type Flag struct {
	SiteMap string   `flag:"sitemap" usage:"SiteMap 文件"`
	Purge   bool     `flag:"purge" usage:"刷新URL"`
	Push    bool     `flag:"push" usage:"预热URL"`
	Include []string `flag:"include" usage:"包含"`
	Exclude []string `flag:"exclude" usage:"排除"`

	Config string `flag:"config" shorthand:"c" usage:"config file, 最高优先级， 可能覆盖其他参数"`
}
