package main

import (
	"math/rand"
	"sync"
	"time"
)

type Order struct {
	ID       int
	Complete bool
}

var orders []*Order
var completeOrders map[int]bool
var wg sync.WaitGroup
var processTimes chan time.Duration
var sinceProgramStarted time.Duration
var count int
var limitCount int

func main() {
	count = 30
	limitCount = 5
	processTimes = make(chan time.Duration, count)
	orders = GenerateOrders(count)                 //сгенерировали срез на 30 заказов
	completeOrders = GenerateCompleteOrders(count) //мапа с выполненными заказами
	programStart := time.Now()                     //время старта программы

	LimitSpawnOrderProcessing(limitCount)

	wg.Wait()
	sinceProgramStarted = time.Since(programStart)
	go func() {
		time.Sleep(1 * time.Second)
		close(processTimes)
	}()
	checkTimeDifference(limitCount)
}

func checkTimeDifference(limitCount int) {
	//do not edit
	var averageTime time.Duration
	var orderProcessTotalTime time.Duration
	var orderProcessedCount int
	for v := range processTimes {
		orderProcessedCount++
		orderProcessTotalTime += v
	}
	if orderProcessedCount != count {
		panic("orderProcessedCount != count")
	}
	averageTime = orderProcessTotalTime / time.Duration(orderProcessedCount)

	println("orderProcessTotalTime", orderProcessTotalTime/time.Second)
	println("averageTime", averageTime/time.Second)
	println("sinceProgramStarted", sinceProgramStarted/time.Second)
	println("sinceProgramStarted average", sinceProgramStarted/(time.Duration(orderProcessedCount)*time.Second))
	println("orderProcessTotalTime - sinceProgramStarted", (orderProcessTotalTime-sinceProgramStarted)/time.Second)

	if (orderProcessTotalTime/time.Duration(limitCount)-sinceProgramStarted)/time.Second > 0 {
		panic("(orderProcessTotalTime-sinceProgramStarted)/time.Second > 0")
	}

}

func LimitSpawnOrderProcessing(limitCount int) {
	limit := make(chan struct{}, limitCount) //канал для ограничения горутин
	var t time.Time
	// limit spawn OrderProcesing worker by variable limit
	for _, order := range orders {
		wg.Add(1)
		limit <- struct{}{}
		go func(order *Order) {
			t = time.Now()
			OrderProcessing(order, limit, t)
		}(order)
	}

}

func OrderProcessing(order *Order, limit chan struct{}, t time.Time) {
	//complete orders if they complete
	if completeOrders[order.ID] {
		order.Complete = true
		time.Sleep(1 * time.Second)
	}

	processTimes <- time.Since(t)
	<-limit
	wg.Done()
}

// генерирует заказы
func GenerateOrders(count int) []*Order {
	orders := make([]*Order, count)
	for i := 0; i < count; i++ {
		orders[i] = &Order{ID: i + 1, Complete: false}
	}
	return orders
}

// 50% вероятность завершения заказа
func GenerateCompleteOrders(maxOrderID int) map[int]bool {
	completeOrders := make(map[int]bool)
	for i := 0; i < maxOrderID; i++ {
		if rand.Intn(2) == 0 { // Шанс 50%
			completeOrders[i+1] = true
		}
	}
	return completeOrders
}
