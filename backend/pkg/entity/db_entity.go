package entity

type batteries struct {
	name  string `db:"Battery"`
	price int    `db:"Price"`
}

type tires struct {
	name  string `db:"Tire"`
	price int    `db:"Price"`
}

type wheels struct {
	name  string `db:"Wheel"`
	price int    `db:"Price"`
}

type WheelAvailability struct {
	battery string `db:"BatteryName"`
	wheel   string `db:"WheelName"`
}

type TiresAvailability struct {
	tire  string `db:"TireName"`
	wheel string `db:"WheelName"`
}
