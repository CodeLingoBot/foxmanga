package bootstrap
//
//import (
//	"github.com/facebookgo/inject"
//)
//
//var (
//	g inject.Graph
//)
//
//func init() {
//	//_m := migration.GetMigration(database)
//	//err := g.Provide(
//	//	&inject.Object{Value: database, Name: "database"},
//	//	//&inject.Object{Value: category.GetCategoryService(), Name: "caterepo"},
//	//	&inject.Object{Value: category.CreateRepository(), Name: "caterepo"},
//	//)
//	//if err != nil {
//	//	fmt.Fprintln(os.Stderr, err)
//	//	os.Exit(1)
//	//}
//	//
//	//// Here the Populate call is creating instances of NameAPI &
//	//// PlanetAPI, and setting the HTTPTransport on both to the
//	//// http.DefaultTransport provided above:
//	//if err := g.Populate(); err != nil {
//	//	fmt.Fprintln(os.Stderr, err)
//	//	os.Exit(1)
//	//}
//	//erra := _m.Migrate()
//	//if erra != nil {
//	//	panic(erra)
//	//}
//}
