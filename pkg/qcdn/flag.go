package qcdn

type Flag struct {
	SiteMap       string   `yaml:"sitemap" flag:"sitemap" usage:"SiteMap 文件"`
	Purge         bool     `yaml:"purge" flag:"purge" usage:"刷新URL"`
	Push          bool     `yaml:"push" flag:"push" usage:"预热URL"`
	Include       []string `yaml:"include" flag:"include" usage:"包含"`
	Exclude       []string `yaml:"exclude" flag:"exclude" usage:"排除"`
	LastModInDays int      `yaml:"lastModInDays" flag:"lastModInDays" usage:"几天之内修改的"`

	Config string `yaml:"config" flag:"config" shorthand:"c" usage:"config file, 最高优先级， 可能覆盖其他参数"`
}
