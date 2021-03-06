package qpay

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
)

type Config struct {
	UseSandbox      bool   `json:"use_sandbox"`
	ReturnRaw       bool   `json:"return_raw"`
	AppId           string `json:"app_id"`
	AppKey          string `json:"app_key"`
	MchId           string `json:"mch_id"`
	Md5Key          string `json:"md5_key"`
	FeeType         string `json:"fee_type"`
	NotifyUrl       string `json:"notify_url"`
	RefundNotifyUrl string `json:"refund_notify_url"`
	RedirectUrl     string `json:"redirect_url"`
}

func (c Config) ParseConfig() interface{} {
	return c
}

func NewConfigWithJson(content []byte) (*Config, error) {
	var config Config
	err := json.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}

func NewConfigWithYaml(content []byte) (*Config, error) {
	var config Config
	err := yaml.Unmarshal(content, &config)
	if err != nil {
		return nil, err
	}
	return &config, nil
}
