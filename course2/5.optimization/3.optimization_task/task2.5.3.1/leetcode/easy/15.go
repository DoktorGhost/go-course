/*
15. Дизайн парковочной системы
Спроектировать парковочную систему для парковки. На парковке есть три вида парковочных мест: большие, средние и маленькие, с фиксированным количеством мест для каждого размера.

Реализуйте класс ParkingSystem:

ParkingSystem(int big, int medium, int small) Инициализирует объект класса ParkingSystem. Количество мест для каждого парковочного места указано в конструкторе.

bool addCar(int carType) Проверяет, есть ли парковочное место типа carType для автомобиля, который хочет попасть на стоянку. carType может быть трех видов: большой, средний или маленький, которые представлены цифрами 1, 2 и 3 соответственно. Автомобиль может парковаться только на парковочном месте своего типа автомобиля carType. Если места нет, верните false, в противном случае припаркуйте машину на месте такого размера и верните true.

Пример 1:

# Input

["ParkingSystem", "addCar", "addCar", "addCar", "addCar"]

[[1, 1, 0], [1], [2], [3], [1]]

# Output

[null, true, true, false, false]

# Объяснение

ParkingSystemparkingSystem = новая ParkingSystem(1, 1, 0);

parkingSystem.addCar(1); // возвращаем true, потому что есть 1 свободный слот для большой машины

parkingSystem.addCar(2); // вернем true, потому что есть 1 свободный слот для средней машины

parkingSystem.addCar(3); // вернем false, потому что нет свободного слота для маленькой машины

parkingSystem.addCar(1); // вернем false, потому что нет свободного места для большой машины. Оно уже занято.

Ограничения:

0 <= большой, средний, маленький <= 1000

carType — 1, 2 или 3.

В addCar будет сделано не более 1000 запросов.

type ParkingSystem struct {

}

func Constructor(big int, medium int, small int) ParkingSystem {

}

func (this \*ParkingSystem) AddCar(carType int) bool {

}
https://leetcode.com/problems/design-parking-system/
*/
package main

type ParkingSystem struct {
	big    int
	medium int
	small  int
}

func Constructor(big int, medium int, small int) ParkingSystem {
	res := ParkingSystem{big, medium, small}
	return res
}

func (this *ParkingSystem) AddCar(carType int) bool {
	switch carType {
	case 1:
		if this.big >= 1 {
			this.big--
			return true
		}
	case 2:
		if this.medium >= 1 {
			this.medium--
			return true
		}
	case 3:
		if this.small >= 1 {
			this.small--
			return true
		}
	}

	return false
}
