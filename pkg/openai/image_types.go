package openai

type Model string
type Size string

const (
	DALLE_2 Model = "dall-e-2"
	DALLE_3 Model = "dall-e-3"
	 
	DALLE_2_SMALL Size = "256x256"
	DALLE_2_MID Size = "512x512"
	DALLE_2_LARGE Size = "1024x1024"
	DALLE_3_SQUARE Size = "1024x1024"
	DALLE_3_PORTRAIT Size = "1024x1792"
	DALLE_3_LANDSCAPE Size = "1792x1024"
)