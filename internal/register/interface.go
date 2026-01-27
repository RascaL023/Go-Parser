package register
//
// import "encoding/json"
//
// var RegisteredTool = map[string]Tool{}
//
// type Parser interface {
// 	Parse(json.RawMessage) (any, error)
// }
//
// type Resolver interface {
// 	Resolve(input any) (any, error)
// }
//
// type Renderer interface {
// 	Render(templatePath, outputPath string, data any) error
// }
//
// type Tool interface {
//     Name() string
// 		Parser
// 		Resolver
// 		Renderer
// }
//
// func RegisterTool(t Tool) {
//     RegisteredTool[t.Name()] = t;
// }
//
// func GetTool(name string) (Tool, bool) {
//     p, ok := Registered[name];
//     return p, ok;
// }
//
