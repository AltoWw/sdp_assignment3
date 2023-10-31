package main

import "fmt"

type command interface {
	execute()
}

type device interface { 
	on()
	off()
}

type onCommand struct { 
	device device
}

func (r *onCommand) setDevice(device device) {
	r.device = device
}

func (loc onCommand) execute() {
	loc.device.on()
}

type offCommand struct { 
	device device
}

func (r *offCommand) setDevice(device device) {
	r.device = device
}

func (loc *offCommand) execute() {
	loc.device.off()
}

type lightBulb struct {
}

func (lb lightBulb) on() {
	fmt.Println("Light bulb ON!")
}

func (lb lightBulb) off() {
	fmt.Println("Light bulb OFF!")
}

type remote struct {
	cmd command
}

func (r *remote) setCommand(cmd command) {
	r.cmd = cmd
}

func (r *remote) pressButton() {
	r.cmd.execute()
}

func main() {
	lightBulb := &lightBulb{}

	onCommand := &onCommand{}
	onCommand.setDevice(lightBulb)

	offCommand := &offCommand{}
	offCommand.setDevice(lightBulb)

	oncmd := &remote{}
	oncmd.setCommand(onCommand)
	oncmd.pressButton()

	offcmd := &remote{}
	offcmd.setCommand(offCommand)
	offcmd.pressButton()

}