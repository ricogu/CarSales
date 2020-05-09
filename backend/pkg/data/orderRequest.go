package data

type OrderSubmission struct {
	CustomerName string `json:"CustomerName"`
	BatteryId    int    `json:"BatteryId"`
	WheelId      int    `json:"WheelId"`
	TireId       int    `json:"TireId"`
}
