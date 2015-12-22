package config

type AppConfig struct {
	Config1      string     `json:"config1"`
	Config2      string     `json:"config2"`
	NestedConfig MoreConfig `json:"moreConfig"`
}
type MoreConfig struct {
	more1 string `json:"more1"`
	more2 string `json:"more2"`
}
