ackage main

import (
	"fmt"
	"reflect"
	"strings"
)

var (
	ON  = "ON"
	OFF = "OFF"
)

type Device interface {
	TurnOn() error
	TurnOff() error
	GetStatus() string
}

type AutomationRule interface {
	Execute(devices []Device) error
}

// ------------------------------------------------------
type Light struct {
	Status     string
	Brightness int
}

func NewLight() *Light {
	return &Light{
		Status: OFF,
	}
}

func (l *Light) TurnOn() error {
	if l.Status == OFF {
		l.Status = ON
	}
	return nil
}

func (l *Light) TurnOff() error {
	if l.Status == ON {
		l.Status = OFF
	}
	return nil
}

func (l *Light) GetStatus() string {
	return l.Status
}

func (l *Light) ChangeBrightness(val int) error {
	if val < 0 || val > 100 {
		return fmt.Errorf("Cant set brightness %d", val)
	}
	l.Brightness = val
	return nil
}

// ------------------------------------------------------

type Thermostat struct {
	Status      string
	Temperature int
}

func NewThermostat() *Thermostat {
	return &Thermostat{Status: OFF}
}

func (t *Thermostat) TurnOn() error {
	if t.Status == OFF {
		t.Status = ON
	}
	return nil
}

func (t *Thermostat) TurnOff() error {
	if t.Status == ON {
		t.Status = OFF
	}
	return nil
}

func (t *Thermostat) GetStatus() string {
	return t.Status
}

func (t *Thermostat) ChangeTemperature(val int) error {
	if val < 18 || val > 35 {
		return fmt.Errorf("Cant set temperature %d", val)
	}
	t.Temperature = val
	return nil
}

// ------------------------------------------------------

type MorningRoutine struct {
}

func (mr *MorningRoutine) Execute(devices []Device) error {
	for _, device := range devices {
		if light, ok := device.(*Light); ok {
			light.TurnOn()
			if err := light.ChangeBrightness(100); err != nil {
				return err
			}
		}

		if thermostat, ok := device.(*Thermostat); ok {
			thermostat.TurnOn()
			if err := thermostat.ChangeTemperature(22); err != nil {
				return err
			}
		}
	}
	return nil
}

// ------------------------------------------------------

type SmartHomeController struct {
	Devices []Device
}

func (shc *SmartHomeController) AddDevice(device Device) {
	shc.Devices = append(shc.Devices, device)
}

func (shc *SmartHomeController) RemoveDevice(device Device) {
	for i, d := range shc.Devices {
		if d == device {
			shc.Devices = append(shc.Devices[:i], shc.Devices[i+1:]...)
		}
	}
}

func (shc *SmartHomeController) RunAutomation(rule AutomationRule) error {
	return rule.Execute(shc.Devices)
}

func (shc *SmartHomeController) PrintDeviceStatuses() {
	for _, device := range shc.Devices {
		fmt.Print(strings.Split(reflect.TypeOf(device).String(), ".")[1], " - ", device.GetStatus(), "\n")
	}
}

// ------------------------------------------------------

func main() {
	shc := SmartHomeController{}

	light := NewLight()
	thermostat := NewThermostat()

	shc.AddDevice(light)
	shc.AddDevice(thermostat)

	mr := MorningRoutine{}

	shc.RunAutomation(&mr)

	shc.PrintDeviceStatuses()
}
