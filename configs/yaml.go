package configs

import "os"

const defaultYAMLPath = "plates.yaml"

var yamlPath string

// GetYAMLPath returns the file path to YAML recipes.
func GetYAMLPath() string {
	return yamlPath
}

func init() {
	yamlPath = os.Getenv("YAML_PATH")
	if yamlPath == "" {
		yamlPath = defaultYAMLPath
	}
}