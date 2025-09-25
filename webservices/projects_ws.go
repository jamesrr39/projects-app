package webservices

import (
	"context"

	"github.com/jamesrr39/go-openapix"

	"github.com/jamesrr39/projects-app/dal"
	"github.com/jamesrr39/projects-app/domain"

	"github.com/go-chi/chi/v5"
	"github.com/swaggest/openapi-go/openapi3"
	"github.com/swaggest/rest"
	"github.com/swaggest/rest/chirouter"
	"github.com/swaggest/rest/jsonschema"
	"github.com/swaggest/rest/nethttp"
	"github.com/swaggest/rest/openapi"
	"github.com/swaggest/rest/request"
	"github.com/swaggest/rest/response"
)

func CreateApiRouter(d *dal.ProjectScanner, baseDir string) (*openapi.Collector, *chirouter.Wrapper) {
	apiSchema := &openapi.Collector{}
	apiSchema.Reflector().SpecEns().Info.Title = "Projects"
	apiSchema.Reflector().SpecEns().Info.WithDescription("REST API definitions for Projects App")

	serverDesc := "API server"

	apiSchema.Reflector().SpecEns().Info.Version = "0"
	apiSchema.Reflector().Spec.Servers = append(apiSchema.Reflector().Spec.Servers, openapi3.Server{
		URL:         "/api",
		Description: &serverDesc,
	})

	// Setup request decoder and validator.
	validatorFactory := jsonschema.NewFactory(apiSchema, apiSchema)
	decoderFactory := request.NewDecoderFactory()
	decoderFactory.ApplyDefaults = true
	decoderFactory.SetDecoderFunc(rest.ParamInPath, chirouter.PathToURLValues)

	apiRouter := chirouter.NewWrapper(chi.NewRouter())
	apiRouter.Use(
		nethttp.OpenAPIMiddleware(apiSchema),          // Documentation collector.
		request.DecoderMiddleware(decoderFactory),     // Request decoder setup.
		request.ValidatorMiddleware(validatorFactory), // Request validator setup.
		response.EncoderMiddleware,                    // Response encoder setup.
	)

	apiRouter.Route("/v1", func(r chi.Router) {
		openapix.Post(r, "/projects", GetAllProjects(d, baseDir))
	})

	// check array types are marked as non-null; i.e. no items will return "[]" instead of "null"
	openapix.MustCheckNonNullArrays(apiSchema.Reflector().Spec.Components.Schemas.MapOfSchemaOrRefValues)
	openapix.MustNotHaveDuplicateOperationIDOrUnknownSecurity(apiSchema.Reflector().Spec)

	return apiSchema, apiRouter
}

type EmptyStruct struct{}

type ListProjectsResponse struct {
	Projects []domain.Project `json:"projects" nullable:"false"`
}

func GetAllProjects(d *dal.ProjectScanner, baseDir string) *nethttp.Handler {
	return openapix.MustCreateOpenapiEndpoint(
		"Get service info",
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
