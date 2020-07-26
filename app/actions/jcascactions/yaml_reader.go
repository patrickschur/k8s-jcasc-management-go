package jcascactions

import (
	"fmt"
	"github.com/goccy/go-yaml"
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

// ReadYamlNode reads a node of a YAML file and returns it as String
func ReadYamlNodeFromFile(file string, yamlPath string) (yamlNodeAsString string, err error) {
	// read the file
	yamlString, err := ReadYamlFileAsString(file)
	if err != nil {
		return yamlNodeAsString, err
	}

	// set the path
	path, err := yaml.PathString(yamlPath)
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(
			"  -> Unable get data from YAML",
			fmt.Sprintf("File: [%s]\nPath: [%s]\nError:\n%v", file, yamlPath, err.Error()))
		return yamlNodeAsString, err
	}

	// read content from path
	yamlNode, err := path.ReadNode(strings.NewReader(yamlString))
	if err != nil {
		loggingstate.AddErrorEntryAndDetails(
			"  -> Unable to read node from YAML",
			fmt.Sprintf("File: [%s]\nPath: [%s]\nError:\n%v", file, yamlPath, err.Error()))
		return yamlNodeAsString, err
	}

	return yamlNode.String(), err
}

func ReadYamlValueAsStringFromFile(file string, yamlPath string) (value string, err error) {
	// read file
	yamlFile, err := ReadYamlFileAsString(file)
	if err != nil {
		return value, err
	}
	// read values
	return ReadYamlValueAsString(yamlFile, yamlPath)
}

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

func ReadYamlValueAsStringArrayFromFile(file string, yamlPath string) (values []string, err error) {
	// read file
	yamlFile, err := ReadYamlFileAsString(file)
	if err != nil {
		return values, err
	}
	// read values
	return ReadYamlValueAsStringArray(yamlFile, yamlPath)
}

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
