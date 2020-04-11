package config

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"path"
	"runtime"
)

type ApprovalQueueConfig struct {
	MysqlPass       string `yaml:"mysql_pass"`
	MysqlUser       string `yaml:"mysql_user"`
	MysqlHost       string `yaml:"mysql_host"`
	KafkaBrokerHost string `yaml:"kafka_broker_host"`
	KafkaTopic      string `yaml:"kafka_topic"`
}

func Get(appEnv string) (*ApprovalQueueConfig, error) {
	if appEnv == "" {
		return nil, errors.New("environment not provided")
	}

	_, currentFilePath, _, _ := runtime.Caller(0)
	yamlData, err := ioutil.ReadFile(path.Join(path.Dir(currentFilePath), appEnv+"/config.yaml"))
	if err != nil {
		return nil, err
	}

	conf := &ApprovalQueueConfig{}
	err = yaml.Unmarshal(yamlData, conf)
	if err != nil {
		log.Fatal(fmt.Sprintf("config error: error = %s", err.Error()))
		return nil, err
	}

	return conf, nil
}
