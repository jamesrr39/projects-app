package dal

import (
	"os"
	"path/filepath"

	"github.com/jamesrr39/go-errorsx"

	"github.com/go-git/go-billy/v6/osfs"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/cache"
	"github.com/go-git/go-git/v6/storage/filesystem"
)

type Project struct {
	FilePath string   `json:"filePath"`
	GitStats GitStats `json:"gitStats"`
}

type ProjectScanner struct {
	Projects []Project `json:"projects"`
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
	URLs   []string `json:"urls"`
	Mirror bool     `json:"mirror"`
}

type GitStats struct {
	Head    GitHead     `json:"head"`
	Status  GitStatus   `json:"status"`
	Remotes []GitRemote `json:"remotes"`
}

func (ps *ProjectScanner) ScanForProjects(baseDir string) errorsx.Error {
	gitDir := filepath.Join(baseDir, git.GitDirName)
	_, err := os.Stat(gitDir)
	if err != nil && !os.IsNotExist(err) {
		return errorsx.Wrap(err, "gitDir", gitDir)
	}

	if err == nil {
		// .git dir found

		fs := osfs.New(gitDir)
		s := filesystem.NewStorageWithOptions(fs, cache.NewObjectLRUDefault(), filesystem.Options{KeepDescriptors: true})
		repo, err := git.Open(s, fs)
		if err != nil {
			return errorsx.Wrap(err, "dir", baseDir, "gitDir", gitDir)
		}

		workTree, err := repo.Worktree()
		if err != nil {
			return errorsx.Wrap(err, "dir", baseDir, "gitDir", gitDir)
		}

		status, err := workTree.Status()
		if err != nil {
			return errorsx.Wrap(err, "dir", baseDir, "gitDir", gitDir)
		}

		head, err := repo.Head()
		if err != nil {
			return errorsx.Wrap(err, "dir", baseDir, "gitDir", gitDir)
		}

		rawRemotes, err := repo.Remotes()
		if err != nil {
			return errorsx.Wrap(err, "dir", baseDir, "gitDir", gitDir)
		}

		remotes := []GitRemote{}
		for _, rawRemote := range rawRemotes {
			remotes = append(remotes, GitRemote{
				Name:   rawRemote.Config().Name,
				URLs:   rawRemote.Config().URLs,
				Mirror: rawRemote.Config().Mirror,
			})
		}

		ps.Projects = append(ps.Projects, Project{
			FilePath: baseDir,
			GitStats: GitStats{
				Head: GitHead{
					Text: head.String(),
				},
				Status: GitStatus{
					Clean: status.IsClean(),
					// Text:  status.String(),
				},
				Remotes: remotes,
			},
		})

		return nil
	}

	entries, err := os.ReadDir(baseDir)
	if err != nil {
		return errorsx.Wrap(err, "dir", baseDir)
	}

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}

		err := ps.ScanForProjects(filepath.Join(baseDir, entry.Name()))
		if err != nil {
			return errorsx.Wrap(err, "dir", baseDir)
		}
	}

	return nil
}
