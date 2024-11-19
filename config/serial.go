package config

import "time"

var (
	Name        = "COM3"
	Baud        = 9600
	ReadTimeout = time.Second * 5
)
