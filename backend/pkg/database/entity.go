package database

type Batteries struct {
	Id    int    `db:"ID"`
	Name  string `db:"Battery"`
	Price int    `db:"Price"`
}

type Tires struct {
	Id    int    `db:"ID"`
	Name  string `db:"Tire"`
	Price int    `db:"Price"`
}

type Wheels struct {
	Id    int    `db:"ID"`
	Name  string `db:"Wheel"`
	Price int    `db:"Price"`
}

type WheelAvailability struct {
	BatteryId string `db:"BatteryID"`
	WheelId   string `db:"WheelID"`
}

type TiresAvailability struct {
	TireId    string `db:"TireID"`
	BatteryId string `db:"BatteryID"`
}

type Orders struct {
	Id          int    `db:"ID"`
	TireName    string `db:"Tire"`
	TireCost    int    `db:"TirePrice"`
	BatteryName string `db:"Battery"`
	BatteryCost int    `db:"BatteryPrice"`
	WheelName   string `db:"Wheel"`
	WheelCost   int    `db:"WheelPrice"`
	Discount    bool   `db:"Discount"`
	NetCost     int    `db:"NetCost"`
	FinalCost   int    `db:"FinalCost"`
}
