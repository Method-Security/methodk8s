package config

type ServiceAccountConfig struct {
	ServiceAccount bool
	Token          string
	CACert         string
}

type KubeConfig struct {
	Context string
	Path    string
	URL     string
}

type RootFlags struct {
	Quiet                bool
	Verbose              bool
	ServiceAccountConfig ServiceAccountConfig
	KubeConfig           KubeConfig
}
