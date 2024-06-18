// Package cmd
// @Author: yinwei
// @File: model
// @Version: 1.0.0
// @Date: 2024/6/17 09:42

package model

type GenCfg struct {
	Host     string   `json:"host" yaml:"host"`
	Port     string   `json:"port" yaml:"port"`
	Database string   `json:"database" yaml:"database"`
	Auth     string   `json:"auth" yaml:"auth"`
	Tables   []string `json:"tables" yaml:"tables"`
	Outpath  string   `json:"outpath" yaml:"outpath"`
}
