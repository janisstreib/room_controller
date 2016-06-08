package main

import "fmt"
import "github.com/janisstreib/room_controller/motion_sensor"


func addInts(x int, y int) int {
    return x + y
}

func main() {
    fmt.Printf("Hello World\n")
    m := motion_sensor.NewGPIOMotionSensor(7)
    for {
	fmt.Printf("%d\n", m.MotionActive())
    }
}
