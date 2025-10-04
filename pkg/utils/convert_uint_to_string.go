package utils

import (
	"strconv"
)

func ConvertUintToId(id uint) string {
	Id := strconv.Itoa(int(id))
	return Id
}
