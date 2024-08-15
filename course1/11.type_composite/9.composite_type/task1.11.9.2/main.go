package main

import "fmt"

type TVer interface {
	switchOFF() bool
	switchOn() bool
	GetStatus() bool
	GetModel() string
}

type LGer struct {
	status bool
	model  string
}

func (l *LGer) switchOFF() bool {
	l.status = false
	return true
}
func (l *LGer) switchOn() bool {
	l.status = true
	return true
}
func (l *LGer) GetStatus() bool {
	return l.status
}
func (l *LGer) GetModel() string {
	return l.model
}
func (l *LGer) LGHub() string {
	return "LG Hub"
}

type Samsunger struct {
	status bool
	model  string
}

func (s *Samsunger) switchOFF() bool {
	s.status = false
	return true
}
func (s *Samsunger) switchOn() bool {
	s.status = true
	return true
}
func (s *Samsunger) GetStatus() bool {
	return s.status
}
func (s *Samsunger) GetModel() string {
	return s.model
}
func (s *Samsunger) SamsungHub() string {
	return "Samsung Hub"
}

func operateTV(tv TVer) {
	fmt.Println("Model:", tv.GetModel())
	fmt.Println("Is TV on?", tv.GetStatus())
	tv.switchOFF()
	fmt.Println("Turned off. Is TV on?", tv.GetStatus())
	tv.switchOn()
	fmt.Println("Turned on. Is TV on?", tv.GetStatus())
}

func main() {
	samsungTV := &Samsunger{
		status: true,
		model:  "Samsung XL-100500",
	}

	lgTV := &LGer{
		status: true,
		model:  "LG SmartTV-3000",
	}

	operateTV(samsungTV)
	fmt.Println()
	operateTV(lgTV)
}
