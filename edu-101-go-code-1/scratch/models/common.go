package models

type TranslateRequest struct {
	Name     string
	Language string
}

type TranslateOutput struct {
	Greeting string
}
