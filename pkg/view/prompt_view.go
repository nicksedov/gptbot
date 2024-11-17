package view

type PromptsTabView struct {
	PromptViews []PromptView `json:"prompts"`
}

type PromptView struct {
	ID           uint          `json:"id"`
	Title        string        `json:"title"`
	Prompt       string        `json:"prompt"`
	AltText      string        `json:"altText"`
	Hidden       bool          `json:"hidden"`
	PromptParams []PromptParam `json:"prompt_params"`
}

type PromptParam struct {
	Tag   string `json:"tag"`
	Title string `json:"title"`
}

type NewPromptFormView struct {
	Title       string `json:"title"`
	Prompt      string `json:"prompt"`
	AltText     string `json:"altText"`
	ParamTag1   string `json:"p1_name"`
	ParamTitle1 string `json:"p1_title"`
	ParamTag2   string `json:"p2_name"`
	ParamTitle2 string `json:"p2_title"`
	ParamTag3   string `json:"p3_name"`
	ParamTitle3 string `json:"p3_title"`
}
