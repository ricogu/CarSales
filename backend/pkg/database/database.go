package database

import (
	"errors"
	"fmt"
	"github.com/jmoiron/sqlx"
	"time"
)

const (
	BASEPRICE = 12000
	DISCOUNT  = 2000
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

func (c *SqlManager) ListAllOrders() ([]Orders, error) {
	var order []Orders
	getStatement := fmt.Sprintf(
		`select Orders.CustomerName,Orders.BasePrice, Orders.ID,B.Battery as Battery, B.Price as BatteryPrice,
                    T.Tire as Tire , T.Price as TirePrice, 
                    W.Wheel as Wheel, W.Price as WheelPrice, 
                    Discount, NetCost, FinalCost from Orders
					join Batteries B on Orders.BatteryID = B.ID
					join Tires T on Orders.TireID = T.ID
					join Wheels W on Orders.WheelID = W.ID`)
	err := c.db.Select(&order, getStatement)

	//avoid return nil when there is no order
	if order == nil {
		return []Orders{}, err
	}

	return order, err
}

func (c *SqlManager) GetABattery(batteryId int) (Batteries, error) {
	var battery Batteries
	statement := fmt.Sprintf(`SELECT * FROM Batteries WHERE ID = '%d'`, batteryId)
	err := c.db.Get(&battery, statement)
	return battery, err
}

func (c *SqlManager) GetAWheel(wheelId int) (Wheels, error) {
	var wheel Wheels
	statement := fmt.Sprintf(`SELECT * FROM Wheels WHERE ID = '%d'`, wheelId)
	err := c.db.Get(&wheel, statement)
	return wheel, err
}

func (c *SqlManager) GetATire(tireId int) (Tires, error) {
	var tire Tires
	statement := fmt.Sprintf(`SELECT * FROM Tires WHERE ID = '%d'`, tireId)
	err := c.db.Get(&tire, statement)
	return tire, err
}

func (c *SqlManager) SubmitOrder(customerName string, batteryId int, tireId int, wheelId int) (Orders, error) {
	var order Orders

	battery, err := c.GetABattery(batteryId)
	if err != nil {
		return order, err
	}

	wheel, err := c.GetAWheel(wheelId)
	if err != nil {
		return order, err
	}

	tire, err := c.GetATire(tireId)
	if err != nil {
		return order, err
	}

	//check if the wheel selection respect the relation defined
	if !validateWheelAvailability(c.db, batteryId, wheelId) {
		return order, errors.New("wheel cannot be selected based on chosen battery")
	}

	//check if the tire selection respect the relation defined
	if !validateTiresAvailability(c.db, tireId, wheelId) {
		return order, errors.New("tire cannot be selected based on chosen wheel")
	}

	lastFriday := isLastFriday(time.Now())
	netPrice := battery.Price + wheel.Price + tire.Price + BASEPRICE
	var finalPrice int
	if lastFriday {
		finalPrice = netPrice - DISCOUNT
	} else {
		finalPrice = netPrice
	}

	//insert order records into DB
	insertStatement := fmt.Sprintf(`INSERT INTO Orders (CustomerName, BasePrice, TireID, WheelID, BatteryID, Discount, NetCost, FinalCost )
				values ('%s',%d,%d,%d,%d,%t,%d,%d)`, customerName, BASEPRICE,
		tireId,
		wheelId,
		batteryId,
		lastFriday,
		netPrice,
		finalPrice,
	)

	lastInsertID, err := c.db.MustExec(insertStatement).LastInsertId()
	if err != nil {
		return order, err
	}

	getStatement := fmt.Sprintf(
		`select Orders.CustomerName,Orders.BasePrice, Orders.ID,B.Battery as Battery, B.Price as BatteryPrice,
                    T.Tire as Tire , T.Price as TirePrice, 
                    W.Wheel as Wheel, W.Price as WheelPrice, 
                    Discount, NetCost, FinalCost from Orders
					join Batteries B on Orders.BatteryID = B.ID
					join Tires T on Orders.TireID = T.ID
					join Wheels W on Orders.WheelID = W.ID
					where Orders.ID = %d
					`, lastInsertID)

	err = c.db.Get(&order, getStatement)
	if err != nil {
		return order, err
	}

	return order, nil

}

func validateWheelAvailability(db *sqlx.DB, batteryId int, wheelId int) bool {
	var num []int
	statement := fmt.Sprintf(`SELECT COUNT(*) FROM WheelAvailability WHERE BatteryID = '%d' AND WheelID = '%d'`, batteryId, wheelId)
	err := db.Select(&num, statement)
	if err != nil {
		return false
	}
	return num[0] > 0
}

func validateTiresAvailability(db *sqlx.DB, tireID int, wheelId int) bool {
	var num []int
	statement := fmt.Sprintf(`SELECT COUNT(*) FROM TiresAvailability WHERE TireID = '%d' AND WheelID = '%d'`, tireID, wheelId)
	err := db.Select(&num, statement)
	if err != nil {
		return false
	}
	return num[0] > 0
}

func isLastFriday(t time.Time) bool {
	y := t.Year()
	m := t.Month()
	firstDayOfNextMonth := time.Date(y, m+1, 1, 0, 0, 0, 0, time.UTC).Add(-24 * time.Hour)
	lastFridayOfThisMonth := firstDayOfNextMonth.Add(-time.Duration((firstDayOfNextMonth.Weekday()+7-time.Friday)%7) * 24 * time.Hour)
	return t.YearDay() == lastFridayOfThisMonth.YearDay()
}

func (c *SqlManager) ListWheelsByBattery(batteryId int) ([]Wheels, error) {
	var wheel []Wheels
	statement := fmt.Sprintf(
		`SELECT W.ID, W.Wheel, W.Price FROM Batteries 
				JOIN WheelAvailability WA on Batteries.ID = WA.BatteryID
				JOIN Wheels W on WA.WheelID = W.ID
                WHERE Batteries.ID = '%d'`, batteryId)

	err := c.db.Select(&wheel, statement)
	return wheel, err
}

func (c *SqlManager) ListTiresByWheel(wheelId int) ([]Tires, error) {
	var tires []Tires
	statement := fmt.Sprintf(
		`SELECT T.ID, T.Tire, T.Price FROM Wheels
 				JOIN TiresAvailability TA on Wheels.ID = TA.WheelID
 				JOIN Tires T on TA.TireID = T.ID
				WHERE Wheels.ID = '%d'`, wheelId)

	err := c.db.Select(&tires, statement)
	return tires, err
}
