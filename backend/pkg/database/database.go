package database

import (
	"fmt"
	"github.com/jmoiron/sqlx"
)

type SqlManager struct {
	db *sqlx.DB
}

func NewSqlManager(connString string) (*SqlManager, error) {
	conn, err := sqlx.Connect("mysql", connString)

	if err != nil {
		return nil, err
	}
	return &SqlManager{
		db: conn,
	}, nil
}

func (c *SqlManager) ListAllBatteries() ([]Batteries, error) {
	var battery []Batteries
	err := c.db.Select(&battery, "SELECT * FROM Batteries")
	return battery, err
}

func (c *SqlManager) ListWheelsByBattery(batteryName string) ([]Wheels, error) {
	var wheel []Wheels
	statement := fmt.Sprintf(
		`SELECT W.Wheel, W.Price FROM Batteries 
				JOIN WheelAvailability WA ON Batteries.Battery = WA.BatteryName 
				JOIN Wheels W ON WA.WheelName = W.Wheel
                WHERE BatteryName = '%s'`, batteryName)

	err := c.db.Select(&wheel, statement)
	return wheel, err
}

func (c *SqlManager) ListTiresByWheel(wheelName string) ([]Tires, error) {
	var tires []Tires
	statement := fmt.Sprintf(
		`SELECT T.Tire, T.Price FROM Wheels
 				JOIN TiresAvailability TA ON Wheels.Wheel = TA.WheelName
 				JOIN Tires T ON TA.TireName = T.Tire
				WHERE WheelName = "%s"`, wheelName)

	err := c.db.Select(&tires, statement)
	return tires, err
}
