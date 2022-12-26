package models

type SearchStruct struct {
	Regex bool   `json:"regex"`
	Value string `json:"value"`
}
type Column struct {
	Data       int          `json:"data"`
	Name       string       `json:"name"`
	Orderable  bool         `json:"orderable"`
	Search     SearchStruct `json:"search"`
	Searchable bool         `json:"searchable"`
}
type Order struct {
	Column int    `json:"column"`
	Dir    string `json:"dir"`
}
type DTJsonPengguna struct {
	Columms []Column     `json:"columns"`
	Draw    int          `json:"draw"`
	Length  int          `json:"length"`
	Orders  []Order      `json:"order"`
	Search  SearchStruct `json:"search"`
	Start   int          `json:"start"`
	Role    string       `json:"role"`
}
