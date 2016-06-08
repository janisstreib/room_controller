// +build !arm

package motion_sensor

type FakeMotionSensor struct {
}

func (m FakeMotionSensor) MotionActive() bool {
	return false
}

func NewMotionSensor() MotionSensor {
	m := new(FakeMotionSensor)
	return m
}
