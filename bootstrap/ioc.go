package bootstrap

import (
	"Project/foxmanga/pkg/category"
	"fmt"
	"github.com/facebookgo/inject"
	"os"
)

var (
	g inject.Graph
)

func init() {
	err := g.Provide(
		&inject.Object{Value: &database, Name: "database"},
		//&inject.Object{Value: category.GetCategoryService(), Name: "caterepo"},
		&inject.Object{Value: category.CreateRepository(), Name: "caterepo"},
	)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	// Here the Populate call is creating instances of NameAPI &
	// PlanetAPI, and setting the HTTPTransport on both to the
	// http.DefaultTransport provided above:
	if err := g.Populate(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
