package v1

import "github.com/manifoldco/promptui"

var (
	PromptTemplate = &promptui.PromptTemplates{
		Prompt:          "{{ . }} ",
		Valid:           "{{ . | green }} ",
		Invalid:         "{{ . | red }} ",
		Success:         "{{ . | bold }} ",
		ValidationError: "{{ . | red }} ",
	}

	// select template for git
	GitSelectTemplate = &promptui.SelectTemplates{
		Label:    "{{ . }}?",
		Active:   "{{ .Name | cyan }}",
		Inactive: "{{ .Name | cyan }}",
		Selected: "{{ .Name | red | cyan }}",
		Details: `\n--------- Git ----------
{{ "Name:" | faint }}	{{ .Name }}`,
	}
)

func NewUserPrompt() promptui.Prompt {
	return promptui.Prompt{}
}
