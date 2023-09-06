package main

import "fmt"

type CentralHeating struct {
}

func (ch CentralHeating) Heat(temperature float64) {
	fmt.Printf("Heating with temperature %f\n", temperature)
}

type MusicPlayer struct {
}

func (mp MusicPlayer) Play(song string) {
	fmt.Printf("Playing music : %s", song)
}

type Bulbs struct {
}

func (b *Bulbs) Illuminate(room string) {
	fmt.Println(room, "<< Shining >>")
}
func (b *Bulbs) UnIlluminate(room string) {
	fmt.Println(room, ".. dark ..")
}

type Command interface {
	execute()
}

type PlayMusicCmd struct {
	mp   MusicPlayer
	song string
}

func (cmd PlayMusicCmd) execute() {
	cmd.mp.Play(cmd.song)
}

type HeatCmd struct {
	centralHeating CentralHeating
}

func (cmd HeatCmd) execute() {
	cmd.centralHeating.Heat(27)
}

type CoolCmd struct {
	centralHeating CentralHeating
}

func (cmd CoolCmd) execute() {
	cmd.centralHeating.Heat(18)
}

type ShyHallCmd struct {
	bulbs Bulbs
}

func (cmd ShyHallCmd) execute() {
	cmd.bulbs.Illuminate("Hall")
}

type ShyKitchenCmd struct {
	bulbs Bulbs
}

func (cmd ShyKitchenCmd) execute() {
	cmd.bulbs.Illuminate("Kitchen")
}

type TaskPanel struct {
}

func (tp TaskPanel) submit(cmd Command) {
	cmd.execute()
}

func main() {
	taskPanel := TaskPanel{}

	shyHall := ShyHallCmd{}

	heating := HeatCmd{}

	taskPanel.submit(shyHall)
	taskPanel.submit(heating)
}
