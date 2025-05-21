package soa

import (
	"math/rand"
)

const Size = 100000

type Struct struct {
	x int
	y int
	z int
}

type StructOfArrays struct {
	xs [Size]int
	ys [Size]int
	zs [Size]int
}

func NewStructOfArrays() StructOfArrays {
	xs := [Size]int{}
	ys := [Size]int{}
	zs := [Size]int{}

	for i := 0; i < Size; i++ {
		xs[i] = rand.Int()
		ys[i] = rand.Int()
		zs[i] = rand.Int()
	}

	return StructOfArrays{
		xs: xs,
		ys: ys,
		zs: zs,
	}
}

func NewArrayOfStructs() [Size]Struct {
	aos := [Size]Struct{}
	for i := 0; i < Size; i++ {
		aos[i] = Struct{x: rand.Int(), y: rand.Int(), z: rand.Int()}
	}

	return aos
}

func LoopArrayLinear(aos []Struct) {
	result := 0
	for i := 0; i < len(aos); i++ {
		result += aos[i].x + aos[i].y + aos[i].z
	}
}

func LoopArrayRandom(aos [Size]Struct) {
	result := 0
	for i := 0; i < len(aos); i++ {
		index := rand.Intn(len(aos))
		result += aos[index].x + aos[index].y + aos[index].z
	}
}

func LoopStructLinear(soa StructOfArrays) {
	result := 0
	for i := 0; i < len(soa.xs); i++ {
		result += soa.xs[i] + soa.ys[i] + soa.ys[i]
	}
}

func LoopStructRandom(soa StructOfArrays) {
	result := 0
	for i := 0; i < len(soa.xs); i++ {
		index := rand.Intn(len(soa.xs))
		result += soa.xs[index] + soa.ys[index] + soa.ys[index]
	}
}
