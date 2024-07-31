package config

type RootFlags struct {
	Quiet   bool
	Verbose bool
	Context string
	Path    string
	URL     string
	Secret  bool
	Output  string
}
