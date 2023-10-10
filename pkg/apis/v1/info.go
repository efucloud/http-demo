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

package v1

import (
	"github.com/efucloud/http-demo/pkg/apis/filters"
	"github.com/efucloud/http-demo/pkg/config"
	"github.com/emicklei/go-restful/v3"
	"io"
	"net/http"
)

type InfoResource struct {
}
type RequestData struct {
	RequestURI    string      `json:"requestUri" description:""`
	RemoteAddr    string      `json:"remoteAddr" description:""`
	RequestHeader http.Header `json:"requestHeader" description:""`
	RequestBody   string      `json:"requestBody" description:""`
}
type Body struct {
}

var histories []RequestData

func (r InfoResource) AddWebService(ws *restful.WebService) {

	ws.Route(ws.GET("/health").
		Doc("健康检查").
		Notes("健康检查").
		To(r.health).
		Returns(http.StatusOK, "成功", "ok"))
	ws.Route(ws.GET("/info").
		Doc("查看应用信息").
		Notes("查看应用的编译信息").
		To(r.info).
		Returns(http.StatusOK, "成功", nil).
		Filter(filters.Log))
	ws.Route(ws.GET("/history").
		Doc("查看请求历史").
		Notes("查看请求历史").
		To(r.history).
		Returns(http.StatusOK, "成功", []RequestData{}).
		Filter(filters.Log))
	ws.Route(ws.GET(config.APIPrefix+"/{address:*}").
		Doc("Get").
		Param(ws.PathParameter("address", "address")).
		Notes("接收任意Get请求,返回请求的详细信息").
		To(r.request).
		Returns(http.StatusOK, "成功", RequestData{}).
		Filter(filters.Log))
	ws.Route(ws.POST(config.APIPrefix+"/{address:*}").
		Doc("Get").
		Reads(Body{}).
		Param(ws.PathParameter("address", "address")).
		Notes("接收任意Get请求,返回请求的详细信息").
		To(r.request).
		Returns(http.StatusOK, "成功", RequestData{}).
		Filter(filters.Log))

}
func (r InfoResource) history(req *restful.Request, resp *restful.Response) {
	_ = resp.WriteAsJson(histories)
}
func (r InfoResource) request(req *restful.Request, resp *restful.Response) {
	var reqData RequestData
	reqData.RequestURI = req.Request.RequestURI
	reqData.RemoteAddr = req.Request.RemoteAddr
	reqData.RequestHeader = req.Request.Header
	data, _ := io.ReadAll(req.Request.Body)
	reqData.RequestBody = string(data)
	histories = append(histories, reqData)
	_ = resp.WriteAsJson(reqData)
}
func (r InfoResource) info(req *restful.Request, resp *restful.Response) {
	data := make(map[string]interface{})
	data["application "] = config.ApplicationName
	data["buildDate"] = config.BuildDate
	data["goVersion"] = config.GoVersion
	data["commit"] = config.Commit
	_ = resp.WriteAsJson(data)
}

func (r InfoResource) health(req *restful.Request, resp *restful.Response) {
	_, _ = resp.Write([]byte("ok"))
}
