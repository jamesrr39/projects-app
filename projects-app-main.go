package main

import (
	"encoding/json"
	"os"

	"github.com/jamesrr39/go-errorsx"
	"github.com/jamesrr39/projects-app/dal"
	"gopkg.in/alecthomas/kingpin.v2"
)

func main() {
	filePath := kingpin.Arg("filepath", "").Required().String()
	kingpin.Parse()

	err := run(*filePath)
	errorsx.ExitIfErr(err)
}

func run(filePath string) errorsx.Error {
	var err error

	projectScanner := dal.ProjectScanner{Projects: []dal.Project{}}
	err = projectScanner.ScanForProjects(filePath)
	if err != nil {
		return errorsx.Wrap(err)
	}

	b, err := json.MarshalIndent(projectScanner.Projects, "", "\t")
	if err != nil {
		return errorsx.Wrap(err)
	}

	os.Stdout.Write(b)
	return nil
}
