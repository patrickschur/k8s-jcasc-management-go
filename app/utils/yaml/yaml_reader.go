package yaml

import (
	"fmt"
	"github.com/goccy/go-yaml"
	"github.com/goccy/go-yaml/ast"
	"io/ioutil"
	"k8s-management-go/app/utils/loggingstate"
	"strings"
)

func ReadYamlFileAsString(file string) (yamlFile string, err error) {
	// read file
	yamlByte, err := ioutil.ReadFile(file)
	// check for errors
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(
			"  -> Unable get read YAML file",
			fmt.Sprintf("File: [%s]\n\nError:\n%v", file, err.Error()))
		return yamlFile, err
	}
	return string(yamlByte), err
}

// ReadYamlNodeFromFileAsString reads a node of a YAML file and returns it as String
func ReadYamlNodeFromFileAsString(file string, yamlPath string) (yamlNodeAsString string, err error) {
	// read the file
	yamlNode, err := ReadYamlNodeFromFile(file, yamlPath)
	if err != nil {
		return yamlNodeAsString, err
	}

	return yamlNode.String(), err
}

// ReadYamlNodeFromFile reads a node of a YAML file and returns it as Node
func ReadYamlNodeFromFile(file string, yamlPath string) (yamlNode ast.Node, err error) {
	// read the file
	yamlString, err := ReadYamlFileAsString(file)
	if err != nil {
		return yamlNode, err
	}

	// set the path
	path, err := yaml.PathString(yamlPath)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(
			"  -> Unable get data from YAML",
			fmt.Sprintf("File: [%s]\nPath: [%s]\nError:\n%v", file, yamlPath, err.Error()))
		return yamlNode, err
	}

	// read content from path
	yamlNode, err = path.ReadNode(strings.NewReader(yamlString))
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(
			"  -> Unable to read node from YAML",
			fmt.Sprintf("File: [%s]\nPath: [%s]\nError:\n%v", file, yamlPath, err.Error()))
	}
	return yamlNode, err
}

// ReadYamlValueAsStringFromFile reads a value with YAML path as string directly from a file
func ReadYamlValueAsStringFromFile(file string, yamlPath string) (value string, err error) {
	// read file
	yamlFile, err := ReadYamlFileAsString(file)
	if err != nil {
		return value, err
	}
	// read values
	return ReadYamlValueAsString(yamlFile, yamlPath)
}

// ReadYamlValueAsString reads a value with YAML path from a given YAML string
func ReadYamlValueAsString(yamlAsString string, yamlPath string) (value string, err error) {
	// set the path
	path, err := yaml.PathString(yamlPath)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(
			"  -> Unable get read yaml path from YAML file",
			fmt.Sprintf("YAML: [%s]\nPath: [%s]\nError:\n%v", yamlAsString, yamlPath, err.Error()))
		return value, err
	}

	// read from the path
	err = path.Read(strings.NewReader(yamlAsString), &value)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(
			"  -> Unable get read string from YAML file",
			fmt.Sprintf("YAML: [%s]\nPath: [%s]\nError:\n%v", yamlAsString, yamlPath, err.Error()))
		return value, err
	}

	return value, err
}

// ReadYamlValueAsStringArrayFromFile reads a value with a YAML path directly from a file
func ReadYamlValueAsStringArrayFromFile(file string, yamlPath string) (values []string, err error) {
	// read file
	yamlFile, err := ReadYamlFileAsString(file)
	if err != nil {
		return values, err
	}
	// read values
	return ReadYamlValueAsStringArray(yamlFile, yamlPath)
}

// ReadYamlValueAsStringArray reads a value with a YAML path from a given YAML string
func ReadYamlValueAsStringArray(yamlAsString string, yamlPath string) (values []string, err error) {
	// set the path
	path, err := yaml.PathString(yamlPath)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(
			"  -> Unable get read yaml path from YAML file",
			fmt.Sprintf("YAML: [%s]\nPath: [%s]\nError:\n%v", yamlAsString, yamlPath, err.Error()))
		return values, err
	}

	// read from the path
	err = path.Read(strings.NewReader(yamlAsString), &values)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails("  -> Unable get read string array from YAML file",
			fmt.Sprintf("YAML: [%s]\nPath: [%s]\nError:\n%v", yamlAsString, yamlPath, err.Error()))
		return values, err
	}

	return values, err
}

// ReadYamlAsMapFromFile reads a YAML file and converts the content to a map
func ReadYamlAsMapFromFile(filename string) (yamlMap map[string]interface{}, err error) {
	yamlAsString, err := ReadYamlFileAsString(filename)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(
			"  -> Unable get read YAML file",
			fmt.Sprintf("YAML file: [%s]\nError:\n%v", filename, err.Error()))
	} else {
		yamlMap, err = ReadYamlAsMap(yamlAsString)
	}
	return yamlMap, err
}

// ReadYamlAsMap reads a YAML string and converts it to a map
func ReadYamlAsMap(yamlAsString string) (yamlMap map[string]interface{}, err error) {
	err = yaml.Unmarshal([]byte(yamlAsString), &yamlMap)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(
			"  -> Unable get read YAML file",
			fmt.Sprintf("Unable to parse YAML to map:\n%v", err.Error()))
	}
	return yamlMap, err
}
