package main

import "github.com/gosexy/yaml"

func main() {
	settings := yaml.New()
	settings.Set("success", true)
	settings.Write("test.yaml")
}
