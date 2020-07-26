package jcascactions

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var filename = "../../../projects/example-project/jcasc_config.yaml"

func TestReadYamlValueAsStringFromFile(t *testing.T) {
	systemMessage, err := ReadYamlValueAsStringFromFile(filename, "$.jenkins.systemMessage")
	if err != nil {
		t.Error("Unable to marshal:", err)
	} else {
		assert.Equal(t, "Jenkins instance for namespace [example-project]", systemMessage)
	}
}

func TestReadYamlValueAsStringArrayFromFile(t *testing.T) {
	kubernetesName, err := ReadYamlValueAsStringArrayFromFile(filename, "$.jenkins.clouds[*].kubernetes.name")
	if err != nil {
		t.Error("Unable to marshal:", err)
	} else {
		var expectedNameArray = []string{"jenkins-build-slaves"}
		assert.Equal(t, expectedNameArray, kubernetesName)
	}
}

func TestReadYamlValueAsStringArrayFromFileWithClouds(t *testing.T) {
	cloudNodes, err := ReadYamlNodeFromFile(filename, "$.jenkins.clouds[*]")
	if err != nil {
		t.Error("Unable to marshal:", err)
	} else {
		assert.NotEqual(t, 0, len(cloudNodes))
	}
}

func TestCode(t *testing.T) {
	// read clouds
	cloudNodes, err := ReadYamlNodeFromFile(filename, "$.jenkins.clouds[*].kubernetes.templates")
	if err != nil {
		t.Error("Unable to marshal:", err)
	} else {
		assert.NotEqual(t, 0, len(cloudNodes))
	}

	//nodeYamlByte, _ := ioutil.ReadFile("../../../templates/cloud-templates/node.yaml")
}
