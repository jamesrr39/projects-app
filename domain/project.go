package domain

type Project struct {
	FilePath string   `json:"filePath"`
	GitStats GitStats `json:"gitStats" required:"true"`
}

type GitHead struct {
	Text string `json:"text"`
}

type GitStatus struct {
	Clean bool `json:"clean"`
	// Text  string // can take a long time to return if lots of changes
}

type GitRemote struct {
	Name     string   `json:"name"`
	GitStats GitStats `json:"gitStats" required:"true"`
	URLs     []string `json:"urls" nullable:"false" required:"true"`
	Mirror   bool     `json:"mirror"`
}

type GitStats struct {
	Head    GitHead     `json:"head" required:"true"`
	Status  GitStatus   `json:"status" required:"true"`
	Remotes []GitRemote `json:"remotes" nullable:"false" required:"true"`
}
