package helpers

import "os"

var (
	Public = false
	Name   = os.Getenv("Name_Bot")
	Owner  = os.Getenv("Owner_Number")
)

func SetName(name string) {
	Name = name
}
