package motion_sensor
// #cgo LDFLAGS: -lwiringPi
// #include <wiringPi.h>
import "C"

type GpioMotionSensor struct {
	pin int
}

func (g GpioMotionSensor) MotionActive() bool {
	if C.digitalRead(C.int(g.pin)) == 1 {
		return true
	}
	return false
}

func NewGPIOMotionSensor(pin int) *GpioMotionSensor {
	g := new(GpioMotionSensor)
	g.pin = pin
	C.wiringPiSetupPhys()
	C.pinMode(C.int(pin), C.INPUT)
	return g
}
