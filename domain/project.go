package domain

type Project struct {
	FilePath string   `json:"filePath"`
	GitStats GitStats `json:"gitStats"`
}

type GitHead struct {
	Text string `json:"text"`
}

type GitStatus struct {
	Clean bool `json:"clean"`
	// Text  string // can take a long time to return if lots of changes
}

type GitRemote struct {
	Name   string   `json:"name"`
	URLs   []string `json:"urls" nullable:"false"`
	Mirror bool     `json:"mirror"`
}

type GitStats struct {
	Head    GitHead     `json:"head"`
	Status  GitStatus   `json:"status"`
	Remotes []GitRemote `json:"remotes" nullable:"false"`
}
