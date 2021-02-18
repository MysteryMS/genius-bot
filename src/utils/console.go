package utils

import (
	"encoding/json"
	"github.com/fatih/color"
	"io/ioutil"
	"mystery.tech/m/v2/src/structs"
)

func Info(msg string) {
	color.Cyan("[INFO] %s", msg)
}

func Warn(msg string) {
	w := color.New(color.FgYellow).Add(color.Bold)
	_, _ = w.Printf("[WARN] %s", msg)
}

func Debug(msg string) {
	byteValue, _ := ioutil.ReadFile("config.json")
	var config structs.Config
	_ = json.Unmarshal(byteValue, &config)

	if config.Debug {
		color.White("[DEBUG] \n%s", msg)
	}
}

func Fatal(msg string) {
	f := color.New(color.FgRed).Add(color.Bold)
	_, _ = f.Printf("[FATAL ERR] %s", msg)
}
