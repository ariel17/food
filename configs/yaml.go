package configs

import "os"

const defaultYamlPath = "recipes.yaml"

var yamlPath string

// GetYamlPath returns the file path to YAML recipes.
func GetYamlPath() string {
	return yamlPath
}

func init() {
	yamlPath = os.Getenv("YAML_PATH")
	if yamlPath == "" {
		yamlPath = defaultYamlPath
	}
}