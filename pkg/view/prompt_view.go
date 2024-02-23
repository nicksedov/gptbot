package view

type PromptsTabView struct {
	PromptViews []PromptFormView `json:"prompts"`
}

type PromptFormView struct {
	Title        string `json:"title"`
	Prompt       string `json:"prompt"`
	AltText      string `json:"altText"`
	PromptParams []PromptParam  `json:"prompt_params"`
}

type PromptParam struct {
	Tag   string `json:"tag"`
	Title string `json:"title"`
}
