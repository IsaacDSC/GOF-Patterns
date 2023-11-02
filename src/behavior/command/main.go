package main

import "fmt"

// [ * ] Contract command existents equal all device
type Command interface {
	execute()
}

// [ * ]  Implement interface button tiger button on and button off
type Button struct {
	command Command
}

func (b *Button) press() {
	b.command.execute()
}

// [ * ]  Contract my device commands registered
type Device interface {
	on()
	off()
}

// [ * ]  On Command Driver initialize

type OnCommand struct {
	device Device
}

func (c *OnCommand) execute() {
	c.device.on()
}

// [ * ]  Off Command end Driver

type OffCommand struct {
	device Device
}

func (c *OffCommand) execute() {
	c.device.off()
}

// [ * ] concrete implementation
type Pendrive struct {
	isRunning bool
}

func (t *Pendrive) on() {
	t.isRunning = true
	fmt.Println("Turning Pendrive on")
}

func (t *Pendrive) off() {
	t.isRunning = false
	fmt.Println("Turning Pendrive off")
}

func main() {
	// [ * ] create instance driver
	pendrive := new(Pendrive)
	// [ * ] register commands in pendrive
	onCommand := &OnCommand{device: pendrive}
	offCommand := &OnCommand{device: pendrive}
	// [ * ] register buttons in commands
	btnOn := &Button{command: onCommand}
	btnOff := &Button{command: offCommand}

	// [ * ] implementation logic initialize and shutdown pendrive
	btnOn.press()
	btnOff.press()
}
