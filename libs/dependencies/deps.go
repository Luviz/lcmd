package dependencies

var deps = []Dependency{
	{
		CmdName:     "eza",
		Description: "used instead of ls",
	},
	{
		CmdName:     "fd",
		Description: "used instead of find",
	},
	{
		CmdName:     "jq",
		Description: "json parser",
	},
	{
		CmdName:     "yq",
		Description: "yaml parser",
	},
	{
		CmdName:     "fzf",
		Description: "used for fuzzy search in large data sets",
	},
	{
		CmdName:     "kubectl",
		Description: "k8s cli tool",
	},
	{
		CmdName:     "terraform",
		Description: "terraform cli",
	},
}
