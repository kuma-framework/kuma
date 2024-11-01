package shared

import "github.com/kuma-framework/kuma/v2/internal/domain"

var Modules = map[string]domain.Template{
	"kuma-framework/kuma-typescript-rest-services": domain.NewTemplate(
		"TypeScript Rest Services (OpenAPI 2.0)",
		"Create a library TypeScript with services typed for all endpoints described in a file Open API 2.0",
		[]string{"typescript", "openapi", "rest", "library"},
	),
	"kuma-framework/kuma-hello-world": domain.NewTemplate(
		"Hello World",
		"A simple Hello World in Go!",
		[]string{"golang", "example"},
	),
	"kuma-framework/kuma-changelog-generator": domain.NewTemplate(
		"Changelog Generator",
		"Helper to write a good changelog to your project",
		[]string{"changelog", "helper", "markdown"},
	),
	"kuma-framework/kuma-commit-standardizer": domain.NewTemplate(
		"Commit Standardizer",
		"Write conventional commits for your projects",
		[]string{"git", "standardizer"},
	),
}
