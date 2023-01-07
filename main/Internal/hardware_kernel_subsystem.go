package hardware

import (
	"errors"
	"fmt"
	"time"

	"gobot.io/x/gobot"
	"gobot.io/x/gobot/drivers/gpio"
	"gobot.io/x/gobot/platforms/firmata"
)

// HardwareKernelSubsystem represents a hardware kernel subsystem
type HardwareKernelSubsystem struct {
	board    *firmata.Adaptor
	moisture *gpio.AnalogSensorDriver
	pump     *gpio.RelayDriver
	light    *gpio.LedDriver
}

// NewHardwareKernelSubsystem returns a new HardwareKernelSubsystem instance
func NewHardwareKernelSubsystem(port string) (*HardwareKernelSubsystem, error) {
	board, err := firmata.NewAdaptor(port)
	if err != nil {
		return nil, err
	}
	moisture := gpio.NewAnalogSensorDriver(board, "0")
	pump := gpio.NewRelayDriver(board, "3")
	light := gpio.NewLedDriver(board, "13")
	return &HardwareKernelSubsystem{board, moisture, pump, light}, nil
}
