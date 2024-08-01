package config

type RootFlags struct {
	Quiet          bool
	Verbose        bool
	Context        string
	Path           string
	URL            string
	ServiceAccount bool
	Token          string
	CACert         string
	APIServerURL   string
	Output         string
}
