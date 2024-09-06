package main

import "fmt"

type Mover interface {
	Move() string
	Speed() int
	MaxSpeed() int
	MinSpeed() int
}

type BaseMover struct {
	speed int
}

func (bm BaseMover) Speed() int {
	return bm.speed
}

func (bm BaseMover) MaxSpeed() int {
	return 120
}

func (bm BaseMover) MinSpeed() int {
	return 10
}

type FastMover struct {
	BaseMover
}

func (fm FastMover) Move() string {
	return fmt.Sprintf("Fast mover! Moving at speed: %v", fm.Speed())
}

type SlowMover struct {
	BaseMover
}

func (sm SlowMover) Move() string {
	return fmt.Sprintf("Slow mover... Moving at speed: %v", sm.Speed())
}

func main() {
	var movers []Mover
	fm := FastMover{BaseMover{100}}
	sm := SlowMover{BaseMover{10}}
	movers = append(movers, fm, sm)
	for _, mover := range movers {
		fmt.Println(mover.Move())
		fmt.Println("Maximum speed:", mover.MaxSpeed())
		fmt.Println("Minimum speed:", mover.MinSpeed())
	}
}
