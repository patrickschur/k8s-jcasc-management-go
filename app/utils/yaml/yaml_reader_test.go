package yaml

import (
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

var filename = "../../../templates/jcasc_config.yaml"

func TestReadYamlFileAsString(t *testing.T) {
	yamlContent, err := ReadYamlFileAsString(filename)
	if err != nil {
		t.Error("Unable to read YAML from file: ", err)
	} else {
		assert.NotEmpty(t, yamlContent)
	}
}

func TestReadYamlNodeFromFileAsString(t *testing.T) {
	yamlNode, err := ReadYamlNodeFromFileAsString(filename, "$.jenkins")
	if err != nil {
		t.Error("Unable to read/parse file with YAML part to string:", err)
	} else {
		assert.NotEmpty(t, yamlNode)
		assert.True(t, strings.HasPrefix(strings.TrimLeft(yamlNode, " "), "systemMessage"))
	}
}

func TestReadYamlNodeFromFile(t *testing.T) {
	yamlNode, err := ReadYamlNodeFromFile(filename, "$.jenkins")
	if err != nil {
		t.Error("Unable to read/parse file with YAML part to node:", err)
	} else {
		assert.NotNil(t, yamlNode)
		assert.True(t, strings.HasPrefix(strings.TrimLeft(yamlNode.String(), " "), "systemMessage"))
	}
}

func TestReadYamlValueAsStringFromFile(t *testing.T) {
	systemMessage, err := ReadYamlValueAsStringFromFile(filename, "$.jenkins.systemMessage")
	if err != nil {
		t.Error("Unable to marshal:", err)
	} else {
		assert.Equal(t, "##K8S_MGMT_JENKINS_SYSTEM_MESSAGE##", systemMessage)
	}
}

func TestReadYamlValueAsString(t *testing.T) {
	yamlContent, _ := ReadYamlFileAsString(filename)
	systemMessage, err := ReadYamlValueAsString(yamlContent, "$.jenkins.systemMessage")
	if err != nil {
		t.Error("Unable to marshal:", err)
	} else {
		assert.Equal(t, "##K8S_MGMT_JENKINS_SYSTEM_MESSAGE##", systemMessage)
	}
}

func TestReadYamlValueAsStringArrayFromFile(t *testing.T) {
	kubernetesName, err := ReadYamlValueAsStringArrayFromFile(filename, "$.jenkins.clouds[*].kubernetes.name")
	if err != nil {
		t.Error("Unable to marshal or read values from file:", err)
	} else {
		var expectedNameArray = []string{"jenkins-build-slaves"}
		assert.Equal(t, expectedNameArray, kubernetesName)
	}
}

func TestReadYamlValueAsStringArray(t *testing.T) {
	yamlContent, _ := ReadYamlFileAsString(filename)
	permissions, err := ReadYamlValueAsStringArray(yamlContent, "$.jenkins.authorizationStrategy.roleBased.roles.global[0].permissions")
	if err != nil {
		t.Error("Unable to marshal or read values:", err)
	} else {
		var expectedNameArray = []string{"Overall/Administer"}
		assert.Equal(t, expectedNameArray, permissions)
	}
}

func TestReadYamlAsMapFromFile(t *testing.T) {
	jcascConfigMap, err := ReadYamlAsMapFromFile(filename)
	if err != nil {
		t.Error("Unable to marshal:", err)
	} else {
		assert.NotNil(t, jcascConfigMap)
		var jenkinsMap = jcascConfigMap["jenkins"].(map[string]interface{})
		assert.NotNil(t, jenkinsMap)
		assert.Equal(t, "##K8S_MGMT_JENKINS_SYSTEM_MESSAGE##", jenkinsMap["systemMessage"])
	}
}

func TestReadYamlAsMap(t *testing.T) {
	yamlContent, _ := ReadYamlFileAsString(filename)
	jcascConfigMap, err := ReadYamlAsMap(yamlContent)
	if err != nil {
		t.Error("Unable to marshal:", err)
	} else {
		assert.NotNil(t, jcascConfigMap)
		var jenkinsMap = jcascConfigMap["jenkins"].(map[string]interface{})
		assert.NotNil(t, jenkinsMap)
		assert.Equal(t, "##K8S_MGMT_JENKINS_SYSTEM_MESSAGE##", jenkinsMap["systemMessage"])
	}
}
