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

package server

import (
	"fmt"
	"github.com/efucloud/http-demo/cmd/server/options"
	"github.com/efucloud/http-demo/pkg/apis"
	"github.com/efucloud/http-demo/pkg/config"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/cobra"
	"net/http"
)

func NewRunnerServerCommand() *cobra.Command {
	s := options.NewServerRunOptions()
	cmd := &cobra.Command{
		Use:          "server",
		Long:         `http-demo server`,
		Short:        "http-demo server",
		Example:      `http-demo server`,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			cmd.SilenceUsage = true
			cmd.SilenceErrors = true
			return run(s)
		},
	}
	return cmd
}
func run(o *options.ServerRunOptions) (err error) {
	config.ApplicationConfig.Init()
	apis.AddResources()
	apis.AddSwagger()
	pro := prometheus.NewRegistry()
	http.Handle("/metrics", promhttp.HandlerFor(pro, promhttp.HandlerOpts{}))
	config.Logger.Infof("ready to start server with http and port is: %d", config.ServerPort)
	err = http.ListenAndServe(fmt.Sprintf(":%d", config.ServerPort), nil)
	if err != nil {
		config.Logger.Fatal("ready to start server failed, err: " + err.Error())
	}
	return err
}
