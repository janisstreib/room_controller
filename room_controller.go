package main

import "fmt"
import "time"
import "github.com/janisstreib/room_controller/motion_sensor"

func addInts(x int, y int) int {
	return x + y
}

func main() {
	fmt.Printf("Hello World\n")
	m := motion_sensor.NewMotionSensor()
	for {
		fmt.Printf("%d\n", m.MotionActive())
		time.Sleep(100 * time.Millisecond)
	}
}
