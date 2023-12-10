package service

type Forecast struct {
	Status    Status    `json:"status"`
	Latitude  string    `json:"latitude"`
	Longitude string    `json:"longitude"`
	Forecast  []Weather `json:"forecast"`
}
type Coordinate struct {
	Id        uint    `json:"id"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
}
type Status struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
	Error   bool   `json:"error"`
}
type Weather struct {
	Time    string `json:"time"`
	Weather string `json:"weather"`
}
