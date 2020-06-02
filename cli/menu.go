package cli

import (
	"fmt"
	"github.com/manifoldco/promptui"
	"k8s-management-go/constants"
	"strings"
)

// Menu struct
type menu struct {
	Name        string
	Description string
	Spacer      string
}

func Menu(info string, err error) string {
	// clear screen
	ClearScreen()

	// If errors not emtpy, show them
	if err != nil {
		fmt.Printf("%v[ERROR]: %v%v", constants.ColorError, err, constants.ColorNormal)
		fmt.Println()
		fmt.Println()
	}
	// If infos are available, show them
	if info != "" {
		fmt.Printf("%v[INFO]: %v%v", constants.ColorInfo, info, constants.ColorNormal)
		fmt.Println()
		fmt.Println()
	}

	// Menu structure
	menuStructure := []menu{
		{Name: constants.CommandInstall, Spacer: "                      .-|-:. ", Description: "Install Jenkins of a project"},
		{Name: constants.CommandUninstall, Spacer: "                    .-|-:. ", Description: "Uninstall Jenkins of a project"},
		{Name: constants.CommandUpgrade, Spacer: "                      .-|-:. ", Description: "Upgrade Jenkins in a project"},
		{Name: constants.CommandEncryptSecrets, Spacer: "               .-|-:. ", Description: "Encrypt the secrets file"},
		{Name: constants.CommandDecryptSecrets, Spacer: "               .-|-:. ", Description: "Decrypt the secrets file"},
		{Name: constants.CommandApplySecrets, Spacer: "                 .-|-:. ", Description: "Apply secrets of a project to Kubernetes"},
		{Name: constants.CommandApplySecretsToAll, Spacer: "            .-|-:. ", Description: "Apply secrets to all projects in Kubernetes"},
		{Name: constants.CommandCreateProject, Spacer: "                .-|-:. ", Description: "Create a new project"},
		{Name: constants.CommandCreateDeploymentOnlyProject, Spacer: "  .-|-:. ", Description: "Create a new deployment only project"},
		{Name: constants.CommandCreateJenkinsUserPassword, Spacer: "    .-|-:. ", Description: "Create a password for Jenkins user"},
		{Name: constants.CommandQuit, Spacer: "                         .-|-:. ", Description: "Quit"},
	}

	// Template for displaying menu
	templates := &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "\U000027A4 {{ .Name | green }}{{ .Spacer }}{{ .Description | white }}",
		Inactive: "  {{ .Name | cyan }}{{ .Spacer }}{{ .Description | red }}",
		Selected: "\U000027A4 {{ .Name | red | cyan }}",
		Details: `
--------- Menu selection ----------
{{ "Command    :" | faint }}	{{ .Name }}
{{ "Description:" | faint }}	{{ .Description }}`,
	}

	// searcher (with "/")
	searcher := func(input string, index int) bool {
		menuItem := menuStructure[index]
		name := strings.Replace(strings.ToLower(menuItem.Name), " ", "", -1)
		input = strings.Replace(strings.ToLower(input), " ", "", -1)

		return strings.Contains(name, input)
	}

	prompt := promptui.Select{
		Label:     "Please select the command you want to execute",
		Items:     menuStructure,
		Templates: templates,
		Size:      12,
		Searcher:  searcher,
	}

	i, _, err := prompt.Run()

	if err != nil {
		fmt.Printf("Prompt failed %v\n", err)
		return constants.ErrorPromptFailed
	}

	return menuStructure[i].Name
}