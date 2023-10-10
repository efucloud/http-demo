/*
Copyright 2022 The efucloud.com Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package apis

import (
	"github.com/efucloud/http-demo/pkg/apis/v1"
	"github.com/efucloud/http-demo/pkg/embeds"
	restfulspec "github.com/emicklei/go-restful-openapi/v2"
	"github.com/emicklei/go-restful/v3"
	"github.com/go-openapi/spec"
	"net/http"
)

func AddResources() {
	restful.DefaultRequestContentType(restful.MIME_JSON)
	restful.DefaultResponseContentType(restful.MIME_JSON)
	container := restful.DefaultContainer
	container.Router(restful.CurlyRouter{})
	container.Filter(container.OPTIONSFilter)
	cors := restful.CrossOriginResourceSharing{
		AllowedHeaders: []string{"Content-Type", "Accept", "*"},
		AllowedMethods: []string{"GET", "POST", "DELETE", "PUT", "PATCH", "*"},
		CookiesAllowed: true,
		Container:      container,
	}
	container.Filter(cors.Filter)
	ws := new(restful.WebService)

	ws.Consumes(restful.MIME_JSON).Produces(restful.MIME_JSON)
	v1.InfoResource{}.AddWebService(ws)

	container.Add(ws)
	//for _, ws := range container.RegisteredWebServices() {
	//	for _, route := range ws.Routes() {
	//		config.Logger.Debugf("router name: %s, method: %s, path: %s", route.Doc, route.Method, route.Path)
	//	}
	//}

}

func AddSwagger() {
	c := restfulspec.Config{
		WebServices:                   restful.RegisteredWebServices(), // you control what services are visible
		APIPath:                       "/apidocs.json",
		PostBuildSwaggerObjectHandler: enrichSwaggerObject}
	restful.DefaultContainer.Add(restfulspec.NewOpenAPIService(c))
	http.Handle("/swagger/", http.StripPrefix("/swagger/", http.FileServer(embeds.SwaggerFileSystem())))
}
func enrichSwaggerObject(swo *spec.Swagger) {
	swo.Info = &spec.Info{
		InfoProps: spec.InfoProps{
			Title:       "http demo",
			Description: "golang http demo",
			Contact: &spec.ContactInfo{
				ContactInfoProps: spec.ContactInfoProps{
					Name:  "please change",
					Email: "changeme@aliyun.com",
					URL:   ""},
			},
			License: &spec.License{
				LicenseProps: spec.LicenseProps{
					Name: "MIT",
					URL:  "http://mit.org"},
			},
			Version: "v1.0.0",
		},
	}

}
