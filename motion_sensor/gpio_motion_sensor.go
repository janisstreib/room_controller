package motion_sensor

// +build !arm


type FakeMotionSensor struct {
}

func (m FakeMotionSensor) MotionActive() bool {
	return false
}

func NewMotionSensor() MotionSensor {
	m := new(FakeMotionSensor)
	return m
}