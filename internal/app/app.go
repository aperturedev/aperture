package app

import (
	"errors"
	"fmt"
	"gopkg.in/yaml.v3"
	"os"
)

const fileName = "aperture.yaml"

func New(id string) *App {
	return &App{ID: id}
}

type App struct {
	ID        string     `yaml:"appId"`
	NetConfig *NetConfig `yaml:"netConfig,omitempty"`
}

type NetConfig struct {
	ProjectName string `yaml:"projectName"`
}

func Load() (*App, error) {
	// TODO get wd

	data, err := os.ReadFile(fileName)
	if err != nil {
		// TODO - Return typed err
		return nil, err
	}

	var app App

	err = yaml.Unmarshal(data, &app)
	if err != nil {
		return nil, err
	}

	return &app, nil
}

func (a *App) Save() error {
	data, err := yaml.Marshal(&a)
	if err != nil {
		return errors.Join(err, fmt.Errorf("could not marshal aperture config"))
	}

	err = os.WriteFile(fileName, data, 0755)
	if err != nil {
		return errors.Join(err, fmt.Errorf("could not aperture config"))
	}

	return nil
}

func (a *App) SetNetConfig(cfg NetConfig) {
	a.NetConfig = &cfg
}
