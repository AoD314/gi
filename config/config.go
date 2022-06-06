package config

import (
	"io/ioutil"
	"log"
	"main/data"

	"gopkg.in/yaml.v3"
)

// type GlobalConfig struct {
// 	InstallCfg InstallConfig
// 	CpuCount   int
// }

func InstallConfigLoad(path string) (*data.InstallConfig, error) {
	yamlFile, err := ioutil.ReadFile(path)
	if err != nil {
		log.Printf("Can not read file %s.\nGet errors: #%v ", path, err)
		return nil, err
	}

	icfg := &data.InstallConfig{}

	err = yaml.Unmarshal([]byte(yamlFile), icfg)
	if err != nil {
		log.Fatalf("error: %v", err)
		return nil, err
	}

	return icfg, nil
}

// func InitGlobalConfig() GlobalConfig {
// 	cfg := GlobalConfig{}
// 	cfg.CpuCount = runtime.NumCPU()

// 	return cfg
// }

// func ConfigFile(name string, gcfg GlobalConfig) {

// 	tmpl, _ := template.ParseFiles("templates/" + name)
// 	tmpl.Execute(os.Stdout, gcfg)
// }
