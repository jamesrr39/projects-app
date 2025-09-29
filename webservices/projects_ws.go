package webservices

import (
	"context"

	"github.com/jamesrr39/go-openapix"

	"github.com/jamesrr39/projects-app/dal"
	"github.com/jamesrr39/projects-app/domain"

	"github.com/swaggest/rest/nethttp"
)

type EmptyStruct struct{}

type ListProjectsResponse struct {
	Projects []domain.Project `json:"projects" nullable:"false" required:"true"`
}

func GetAllProjects(d *dal.ProjectScanner, baseDir string) *nethttp.Handler {
	return openapix.MustCreateOpenapiEndpoint(
		"Get projects listing",
		&openapix.HandlerOptions{},
		func(ctx context.Context, input *EmptyStruct, output *ListProjectsResponse) error {

			err := d.ScanForProjects(baseDir)
			if err != nil {
				return err
			}

			output.Projects = d.Projects

			return nil
		},
	)
}
