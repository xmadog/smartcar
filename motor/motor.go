package motor

import (
	"github.com/stianeikeland/go-rpio"
)

type MotorPins struct {
	IN1 *uint8
	IN2 *uint8
	PWM *uint8
}

type motorInfo struct {
	in1 *rpio.Pin
	in2 *rpio.Pin
	pwm *rpio.Pin
}

type Motor struct {
	motor1 *motorInfo
	motor2 *motorInfo
	motor3 *motorInfo
	motor4 *motorInfo
}

func NewMotorPins(in1, in2, pwm *uint8) *MotorPins {
	return &MotorPins{in1, in2, pwm}
}

func NewMotor(motor1, motor2, motor3, motor4 *MotorPins) *Motor {
	motor := &Motor{
		motor1: newMotor(motor1),
		motor2: newMotor(motor2),
		motor3: newMotor(motor3),
		motor4: newMotor(motor4),
	}
	return motor
}

func newMotor(pins *MotorPins) *motorInfo {
	in1, in2 := rpio.Pin(*pins.IN1), rpio.Pin(*pins.IN2)
	var pwm *rpio.Pin = nil
	if pins.PWM != nil {
		pwmPin := rpio.Pin(*pins.PWM)
		pwm = &pwmPin
	}

	m := &motorInfo{
		in1: &in1,
		in2: &in2,
		pwm: pwm,
	}

	m.in1.Output()
	m.in2.Output()

	if m.pwm != nil {
		m.pwm.Output()

		m.pwm.Mode(rpio.Pwm)
		m.pwm.Freq(100)
		m.pwm.DutyCycle(0, 10)
	}

	return m
}

func (controller *Motor) Forward() {
	controller.motor1.in1.High()
	controller.motor1.in2.Low()
	if controller.motor1.pwm != nil {
		controller.motor1.pwm.DutyCycle(10, 10)
	}
	//motor.motor1.pwm.Low()

	controller.motor2.in1.High()
	controller.motor2.in2.Low()
	if controller.motor2.pwm != nil {
		controller.motor2.pwm.DutyCycle(10, 10)
	}
	//motor.motor2.pwm.Low()

	controller.motor3.in1.High()
	controller.motor3.in2.Low()
	if controller.motor3.pwm != nil {
		controller.motor3.pwm.DutyCycle(10, 10)
	}
	//motor.motor3.pwm.Low()

	controller.motor4.in1.High()
	controller.motor4.in2.Low()
	if controller.motor4.pwm != nil {
		controller.motor4.pwm.DutyCycle(10, 10)
	}
	//motor.motor4.pwm.Low()
}

func (controller *Motor) Backward() {
	controller.motor1.in1.Low()
	controller.motor1.in2.High()
	if controller.motor1.pwm != nil {
		controller.motor1.pwm.DutyCycle(10, 10)
	}
	//motor.motor1.pwm.Low()

	controller.motor2.in1.Low()
	controller.motor2.in2.High()
	if controller.motor2.pwm != nil {
		controller.motor2.pwm.DutyCycle(10, 10)
	}
	//motor.motor2.pwm.Low()

	controller.motor3.in1.Low()
	controller.motor3.in2.High()
	if controller.motor3.pwm != nil {
		controller.motor3.pwm.DutyCycle(10, 10)
	}
	//motor.motor3.pwm.Low()

	controller.motor4.in1.Low()
	controller.motor4.in2.High()
	if controller.motor4.pwm != nil {
		controller.motor4.pwm.DutyCycle(10, 10)
	}
	//motor.motor4.pwm.Low()
}

func (controller *Motor) TurnLeft() {
	controller.motor1.in1.High()
	controller.motor1.in2.Low()
	if controller.motor1.pwm != nil {
		controller.motor1.pwm.DutyCycle(10, 10)
	}
	//motor.motor1.pwm.Low()

	controller.motor2.in1.Low()
	controller.motor2.in2.High()
	if controller.motor2.pwm != nil {
		controller.motor2.pwm.DutyCycle(10, 10)
	}
	//motor.motor2.pwm.Low()

	controller.motor3.in1.High()
	controller.motor3.in2.Low()
	if controller.motor3.pwm != nil {
		controller.motor3.pwm.DutyCycle(10, 10)
	}
	//motor.motor3.pwm.Low()

	controller.motor4.in1.Low()
	controller.motor4.in2.High()
	if controller.motor4.pwm != nil {
		controller.motor4.pwm.DutyCycle(10, 10)
	}
	//motor.motor4.pwm.Low()
}

func (controller *Motor) TurnRight() {
	controller.motor1.in1.Low()
	controller.motor1.in2.High()
	if controller.motor1.pwm != nil {
		controller.motor1.pwm.DutyCycle(10, 10)
	}
	//motor.motor1.pwm.Low()

	controller.motor2.in1.High()
	controller.motor2.in2.Low()
	if controller.motor2.pwm != nil {
		controller.motor2.pwm.DutyCycle(10, 10)
	}
	//motor.motor2.pwm.Low()

	controller.motor3.in1.Low()
	controller.motor3.in2.High()
	if controller.motor3.pwm != nil {
		controller.motor3.pwm.DutyCycle(10, 10)
	}
	//motor.motor3.pwm.Low()

	controller.motor4.in1.High()
	controller.motor4.in2.Low()
	if controller.motor4.pwm != nil {
		controller.motor4.pwm.DutyCycle(10, 10)
	}
	//motor.motor4.pwm.Low()
}

func (controller *Motor) Close() {
}
