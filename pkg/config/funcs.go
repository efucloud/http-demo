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

package config

import (
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"
	"os"
)

func logConfig(conf *LogConfig) {

	writeSyncer := zapcore.AddSync(&lumberjack.Logger{
		Filename:   conf.Filename,
		MaxSize:    conf.MaxSize,
		MaxBackups: conf.MaxBackups,
		MaxAge:     conf.MaxAge,
		Compress:   conf.Compress,
	})
	var encoderConfig zapcore.EncoderConfig
	if conf.Production {
		encoderConfig = zap.NewProductionEncoderConfig()
	} else {
		encoderConfig = zap.NewDevelopmentEncoderConfig()
	}
	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder
	encoderConfig.EncodeLevel = zapcore.CapitalLevelEncoder
	encoder := zapcore.NewConsoleEncoder(encoderConfig)
	core := zapcore.NewCore(encoder, zapcore.NewMultiWriteSyncer(writeSyncer, zapcore.AddSync(os.Stdout)), zapcore.DebugLevel)
	logger := zap.New(core, zap.AddCaller())
	Logger = logger.Sugar()

}

func (c *Config) Init() {
	if c.LogConfig == nil {
		c.LogConfig = new(LogConfig)
		c.LogConfig.Filename = "./log/http-demo.log"
		c.LogConfig.MaxAge = 30
		c.LogConfig.MaxSize = 1
		c.LogConfig.MaxBackups = 10
		c.LogConfig.Compress = false
	}
	logConfig(c.LogConfig)

}
