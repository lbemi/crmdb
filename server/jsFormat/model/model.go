package model

type OldJs struct {
	Title  string `json:"title" yaml:"title"`
	Field  string `json:"field" yaml:"field"`
	Crspan string `json:"crspan" yaml:"crspan"`
	Halign string `json:"halign" yaml:"halign"`
	Align  string `json:"align" yaml:"align"`
	//Visible   bool   `yaml:"visible" json:"visible"`
	Formatter string `json:"formatter" yaml:"formatter"`
}

type NewJS struct {
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Type   string   `json:"type"`
	Custom []string `json:"custom"`
	Width  string   `json:"width"`
	Align  string   `json:"align"`
	Order  string   `json:"order"`
	Show   bool     `json:"show"`
}
