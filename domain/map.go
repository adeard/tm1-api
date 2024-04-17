package domain

type MapRequestData struct {
	Title string `json:"title" form:"title"`
	Lat   string `json:"lat" form:"lat"`
	Lng   string `json:"lng" form:"lng"`
}
