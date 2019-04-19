package config

type Proxy struct {
	Context     string            `json:"context"`
	Target      string            `json:"target"`
	PathRewrite map[string]string `json:"pathRewrite"`
}

type Configuration struct {
	Port     int     `json:"port"`
	BaseDir  string  `json:"baseDir"`
	Compress bool    `json:"compress"`
	Proxy    []Proxy `json:"proxy"`
}
