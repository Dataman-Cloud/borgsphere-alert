package model

import (
	"time"
)

type Strategy struct {
	ID       uint64    `json:"Id"`
	Name     string    `json:"Name"`
	Type     uint8     `json:"Type"`
	CreateAt time.Time `json:"CreateAt"`
}
