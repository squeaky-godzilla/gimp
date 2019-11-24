package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"

	"gopkg.in/yaml.v2"
)

// DoomConfig definition
type DoomConfig struct {
	DoomRun     `yaml:"DoomRun"`
	DoomDefault `yaml:"DoomDefault"`
}

// DoomRun defines the configuration of iwad and pwads
type DoomRun struct {
	Iwad         string   `yaml:"iwad"`
	Pwads        []string `yaml:"pwads"`
	DefaultsPath string   `yaml:"defaultsPath"`
}

// DoomDefault defines the folders and binary
type DoomDefault struct {
	Gzdoom     string `yaml:"gzdoom"`
	IwadFolder string `yaml:"iwadFolder"`
	PwadFolder string `yaml:"pwadFolder"`
}

func loadYAML(configPath string, config DoomConfig) DoomConfig {

	configPath, _ = filepath.Abs(configPath)

	configFile, err := ioutil.ReadFile(configPath)
	if err != nil {
		fmt.Printf("Error reading YAML file: %s\n", err)
	}

	err = yaml.Unmarshal(configFile, &config)
	if err != nil {
		fmt.Printf("Error parsing YAML file: %s\n", err)
	}

	return config

}

func assembleCommand(config DoomConfig) []string {
	var commandList []string

	// add binary

	commandList = append(commandList, config.DoomDefault.Gzdoom)

	// add IWAD

	commandList = append(commandList, "-iwad", path.Join(config.DoomDefault.IwadFolder, config.DoomRun.Iwad))

	// add PWADs

	commandList = append(commandList, "-file")

	for _, pwad := range config.DoomRun.Pwads {
		pwad = path.Join(config.DoomDefault.PwadFolder, pwad)
		commandList = append(commandList, pwad)
	}

	return commandList

}

func main() {

	var config DoomConfig

	fmt.Println("Parsing doom YAML run file")

	var runFileName string
	var debugMode bool
	flag.StringVar(&runFileName, "r", "", "doom YAML run config.")
	flag.BoolVar(&debugMode, "d", false, "enable the debug mode")
	flag.Parse()

	if runFileName == "" {
		fmt.Println("Please provide doom YAML run config by using -r option")
		return
	}

	config = loadYAML(runFileName, config)
	config = loadYAML(config.DoomRun.DefaultsPath, config)

	// fmt.Println(assembleCommand(config))

	DoomCommand := assembleCommand(config)

	// fmt.Println(DoomCommand[1:])

	if debugMode == true {
		fmt.Printf("%+v\n", config)
		fmt.Println(DoomCommand)
		os.Exit(1)
	}

	cmd := exec.Command(DoomCommand[0], DoomCommand[1:]...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	err := cmd.Run()
	if err != nil {
		log.Fatalf("Running zdoom failed with: %s\n", err)
	}

}
