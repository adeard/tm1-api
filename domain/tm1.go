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
