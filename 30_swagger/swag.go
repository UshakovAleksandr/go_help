package main

import (
	"embed"
	"html/template"
	"io/fs"
	"net/http"
	"runtime"

	"github.com/julienschmidt/httprouter"
)

var (
	// AppName - application name.
	AppName string
	// AppVersion - version application.
	AppVersion string
	// SourceCode - source code (git link).
	SourceCode string
	// Branch - branch source code.
	Branch string
	// Commit - commit in source code.
	Commit string
	// BuildTime - time build application.
	BuildTime string
	// LambdaVersion - version current module.
	LambdaVersion = "0.1.0"
	// Language - golang version.
	Language = runtime.Version()
)

// BuildInfo holds build information about the app.
var BuildInfo = buildInfo{
	AppName,
	AppVersion,
	SourceCode,
	Branch,
	Commit,
	BuildTime,
	LambdaVersion,
	Language,
}

type buildInfo struct {
	AppName       string `json:"app_name,omitempty"`
	AppVersion    string `json:"app_version,omitempty"`
	SourceCode    string `json:"source_code,omitempty"`
	Branch        string `json:"branch,omitempty"`
	Commit        string `json:"commit,omitempty"`
	BuildTime     string `json:"build_time,omitempty"`
	LambdaVersion string `json:"lambda_version,omitempty"`
	Language      string `json:"language,omitempty"`
}

//go:embed swagger-ui
var swaggerFS embed.FS

//go:embed openapi
var apidocs embed.FS

// openapiHandler get dynamic spec for http-server.
func openapiHandler(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	// Copy source template
	tmpl, _ := template.ParseFS(apidocs, "openapi/openapi.yaml")
	tmpl.Execute(w, BuildInfo) //nolint:errcheck
}

// NewSwaggerHandler returns Handler for endpoint `/swagger/*`.
func NewSwaggerHandler() http.FileSystem {
	fswagger, _ := fs.Sub(swaggerFS, "swagger-ui")
	return http.FS(fswagger)
}




























//type Request struct {
//	Value int    `json:"value"`
//	Price int    `json:"price"`
//	Desc  string `json:"desc"`
//	Foo
//}
//
//type Foo struct {
//	foo string
//	Bar
//	Buz
//}
//
//type Bar struct {
//	bar int
//}
//
//type Buz struct {
//	buz bool
//}
//
//
//
//func structDefine(elem reflect.Value) {
//
//	if elem.Kind() == reflect.Struct {
//		d := make(map[string]map[string]interface{})
//		d[elem.Type().Name()] = map[string]interface{}{}
//
//		for i := 0; i < elem.NumField(); i++ {
//			//fmt.Println(elem.Kind()) = struct
//			// тип поля структуры
//			if elem.Field(i).Kind() != reflect.Struct {
//				fieldType := elem.Field(i).Type().Name()
//				// навание поля структуры
//				fieldName := elem.Type().Field(i).Name
//				d[elem.Type().Name()][fieldName] = fieldType
//				fmt.Println(d)
//			} else {
//				//fmt.Println(elem.Field(i).Type().Name())
//				d[elem.Type().Name()][elem.Type().Field(i).Name] = elem.Field(i).Type().Name()
//				structDefine(elem.Field(i))
//			}
//		}
//		for k, v := range d {
//			global[k] = v
//		}
//	}
//}




















