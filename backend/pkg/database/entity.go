package database

type Batteries struct {
	Name  string `db:"Battery"`
	Price int    `db:"Price"`
}

type Tires struct {
	Name  string `db:"Tire"`
	Price int    `db:"Price"`
}

type Wheels struct {
	Name  string `db:"Wheel"`
	Price int    `db:"Price"`
}

type WheelAvailability struct {
	Battery string `db:"BatteryName"`
	Wheel   string `db:"WheelName"`
}

type TiresAvailability struct {
	Tire  string `db:"TireName"`
	Wheel string `db:"WheelName"`
}
