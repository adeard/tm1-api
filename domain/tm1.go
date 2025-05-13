package domain

type Tm1RequestData struct {
	Cells any
	Value any
}

type Tm1RequestDynamicData struct {
	Updates any
}

type Tm1DynamicRequestData struct {
	Tm1RequestDynamicData
	Cubes string `json:"Cubes"`
	Port  string `json:"Port"`
	Url   string `json:"Url"`
}

type Tm1DynamicInputData struct {
	Cells any `json:"Cells"`
	Value any `json:"Value"`
}

type Tm1GetElementData struct {
	Dimensions  string `json:"Dimensions"`
	Hierarchies string `json:"Hierarchies"`
	Elements    string `json:"Elements"`
}

type Tm1GetElementRequestData struct {
	Tm1GetElementData
	Port string `json:"Port"`
	Url  string `json:"Url"`
}

type Tm1AddElementRequestData struct {
	Port        string `json:"Port"`
	Url         string `json:"Url"`
	Dimensions  string `json:"Dimensions"`
	Hierarchies string `json:"Hierarchies"`
	Parents     string `json:"Parents"`
	Elements    string `json:"Elements"`
}
