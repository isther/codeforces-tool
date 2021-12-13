package config

type CodeTemplate struct {
	Alias        string `json:"alias"`
	Lang         string `json:"lang"`
	Path         string `json:"path"`
	Suffix       string `json:"suffix"`
	BeforeScript string `json:"before_script"`
	Script       string `json:"script"`
	AfterScript  string `json:"after_script"`
}

type Config struct {
	Template      []CodeTemplate    `json:"template"`
	Default       int               `json:"default"`
	GenAfterParse bool              `json:"gen_after_parse"`
	Host          string            `json:"host"`
	Proxy         string            `json:"proxy"`
	FolderName    map[string]string `json:"folder_name"`
	path          string
}
