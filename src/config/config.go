package config

import (
	"fmt"
	"os"

	"gopkg.in/yaml.v3"
)

type Config struct {
	ASSFile          TASSFile          `yaml:"ASSFile"`
	YouTubeIndexFile TYouTubeIndexFile `type:"YouTubeIndexFile"`
}

type TYouTubeIndexFile struct {
	Template string `yaml:"Template"`
}

type TASSFile struct {
	Style      TStyle `yaml:"Style"`
	Template   string `yaml:"Template"`
	HideBefore int    `yaml:"HideBefore"`
}

type TStyle struct {
	Name            string `yaml:"Name"`
	Fontname        string `yaml:"Fontname"`
	Fontsize        int    `yaml:"Fontsize"`
	PrimaryColour   string `yaml:"PrimaryColour"`
	SecondaryColour string `yaml:"SecondaryColour"`
	OutlineColour   string `yaml:"OutlineColour"`
	BackColour      string `yaml:"BackColour"`
	Bold            int    `yaml:"Bold"`
	Italic          int    `yaml:"Italic"`
	Underline       int    `yaml:"Underline"`
	StrikeOut       int    `yaml:"StrikeOut"`
	ScaleX          int    `yaml:"ScaleX"`
	ScaleY          int    `yaml:"ScaleY"`
	Spacing         int    `yaml:"Spacing"`
	Angle           int    `yaml:"Angle"`
	BorderStyle     int    `yaml:"BorderStyle"`
	Outline         int    `yaml:"Outline"`
	Shadow          int    `yaml:"Shadow"`
	Alignment       int    `yaml:"Alignment"`
	MarginL         int    `yaml:"MarginL"`
	MarginR         int    `yaml:"MarginR"`
	MarginV         int    `yaml:"MarginV"`
	Encoding        int    `yaml:"Encoding"`
}

func NewConfig() *Config {
	t := new(Config)
	return t
}

func (t *Config) Load(filename *string) error {
	data, err := os.ReadFile(*filename)
	if err != nil {
		return fmt.Errorf("Configファイル読み込み失敗 - %w", err)
	}

	if err := yaml.Unmarshal(data, t); err != nil {
		return fmt.Errorf("Configファイル処理失敗 - %w", err)
	}
	return nil
}
