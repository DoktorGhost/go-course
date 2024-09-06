package main

import (
	"fmt"
)

type Order interface {
	AddItem(item string, quantity int) error
	RemoveItem(item string) error
	GetOrderDetails() map[string]int
}

// DineInOrder — заказ на месте.
type DineInOrder struct {
	orderDetails map[string]int
}

func (d *DineInOrder) AddItem(item string, quantity int) error {
	d.orderDetails[item] += quantity
	return nil
}
func (d *DineInOrder) RemoveItem(item string) error {
	if _, ok := d.orderDetails[item]; ok {
		delete(d.orderDetails, item)
		return nil
	}
	return fmt.Errorf("item %s does not exist", item)

}
func (d *DineInOrder) GetOrderDetails() map[string]int {
	fmt.Println("Заказ:")
	for k, v := range d.orderDetails {
		fmt.Printf("%s: %d\n", k, v)
	}
	return d.orderDetails
}

// TakeAwayOrder — заказ на вынос.
type TakeAwayOrder struct {
	orderDetails map[string]int
}

func (t *TakeAwayOrder) AddItem(item string, quantity int) error {
	t.orderDetails[item] += quantity
	return nil
}
func (t *TakeAwayOrder) RemoveItem(item string) error {
	if _, ok := t.orderDetails[item]; ok {
		delete(t.orderDetails, item)
		return nil
	}
	return fmt.Errorf("item %s does not exist", item)
}
func (t *TakeAwayOrder) GetOrderDetails() map[string]int {
	fmt.Println("Заказ:")
	for k, v := range t.orderDetails {
		fmt.Printf("%s: %d\n", k, v)
	}
	return t.orderDetails
}

func ManageOrder(o Order) {
	o.AddItem("Pizza", 2)
	o.AddItem("Burger", 1)
	o.RemoveItem("Pizza")
	fmt.Println(o.GetOrderDetails())
}
func main() {
	dineIn := &DineInOrder{orderDetails: make(map[string]int)}
	takeAway := &TakeAwayOrder{orderDetails: make(map[string]int)}
	ManageOrder(dineIn)
	ManageOrder(takeAway)
}
