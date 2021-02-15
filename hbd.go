package main

import (
	"fmt"
	"time")


func birthday(time time.Time) {

	today := time.Now()
	if time == today {
		fmt.Println("Happy Birthday SMANKA 33")
	}

}