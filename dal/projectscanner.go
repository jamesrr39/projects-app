package dal

import (
	"os"
	"path/filepath"

	"github.com/jamesrr39/go-errorsx"

	"github.com/go-git/go-billy/v6/osfs"
	"github.com/go-git/go-git/v6"
	"github.com/go-git/go-git/v6/plumbing/cache"
	"github.com/go-git/go-git/v6/storage/filesystem"

	"github.com/jamesrr39/projects-app/domain"
)

type ProjectScanner struct {
	Projects []domain.Project `json:"projects"`
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

		remotes := []domain.GitRemote{}
		for _, rawRemote := range rawRemotes {
			remotes = append(remotes, domain.GitRemote{
				Name:   rawRemote.Config().Name,
				URLs:   rawRemote.Config().URLs,
				Mirror: rawRemote.Config().Mirror,
			})
		}

		ps.Projects = append(ps.Projects, domain.Project{
			FilePath: baseDir,
			GitStats: domain.GitStats{
				Head: domain.GitHead{
					Text: head.String(),
				},
				Status: domain.GitStatus{
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
