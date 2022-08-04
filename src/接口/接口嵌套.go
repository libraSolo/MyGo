package main

type Flyer interface {
	fly()
}

type Swimmer interface {
	swim()
}

// 接口组合
type FlyFish interface {
	Flyer
	Swimmer
}

type Fish struct{}

func (fish Fish) fly() {
	println("FlyFish, fly")
}
func (fish Fish) swim() {
	println("FlyFish, swim")
}
func main() {
	var fish FlyFish
	fish = Fish{}
	fish.fly()
	fish.swim()
}
