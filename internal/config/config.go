package config

type ServiceAccountConfig struct {
	ServiceAccount bool
	Token          string
	CACert         string
	APIServerURL   string
}

type OtherConfig struct {
	Context string
	Path    string
	URL     string
}

type RootFlags struct {
	Quiet                bool
	Verbose              bool
	ServiceAccountConfig ServiceAccountConfig
	OtherConfig          OtherConfig
}
