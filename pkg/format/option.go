package format

type Option struct {
	Debug  bool
	Lint   bool
	Path   string
	Match  string
	Ignore []string
}
