package main

type ParkingSystem struct {
	slots [4]int
}


func Constructor(big int, medium int, small int) ParkingSystem {
	ps := ParkingSystem{}
	ps.slots[1], ps.slots[2], ps.slots[3] = big, medium, small
	return ps
}


func (this *ParkingSystem) AddCar(carType int) bool {
	this.slots[carType]--
	return this.slots[carType] >= 0
}


/**
 * Your ParkingSystem object will be instantiated and called as such:
 * obj := Constructor(big, medium, small);
 * param_1 := obj.AddCar(carType);
 */
