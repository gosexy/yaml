package main

import (
	"github.com/gosexy/yaml"
)

func main() {
	settings := yaml.New()
	settings.Set("success", true)
	settings.Set("nested", "tree", 1)
	settings.Set("another", "nested", "tree", []int{1, 2, 3})
	settings.Write("test.yaml")
}
