package dto

import "fmt"

type OrderDirection string

const (
	OrderByASC      OrderDirection = "ASC"
	OrderByDESC     OrderDirection = "DESC"
	OrderDefaultKey                = "record_id"
)

type OrderParam struct {
	Key       string         `json:"key" form:"key"`
	Direction OrderDirection `json:"direction" form:"direction"`
}

func (a OrderParam) String() string {
	if a.Key == "" {
		a.Key = OrderDefaultKey
	}

	key := a.Key
	direction := "DESC"
	if a.Direction == OrderByASC {
		direction = "ASC"
	}

	return fmt.Sprintf("%s %s", key, direction)
}
