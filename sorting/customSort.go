package sorting

import (
	"fmt"
	"sort"
)

type Car struct {
	Name       string
	Horsepower int
	Topspeed   int
}

type ByHorsepower []Car
type ByTopspeed []Car

func (hp ByHorsepower) Len() int           { return len(hp) }
func (hp ByHorsepower) Less(i, j int) bool { return hp[i].Horsepower < hp[j].Horsepower }
func (hp ByHorsepower) Swap(i, j int)      { hp[i], hp[j] = hp[j], hp[i] }

func (ts ByTopspeed) Len() int           { return len(ts) }
func (ts ByTopspeed) Less(i, j int) bool { return ts[i].Topspeed > ts[j].Topspeed }
func (ts ByTopspeed) Swap(i, j int)      { ts[i], ts[j] = ts[j], ts[i] }

func CustomSort() {
	cars := []Car{
		{"nascar", 670, 321},
		{"fusca", 36 , 200},
		{"f1", 1000, 378},
		{"uno", 58 , 450},
	}
	fmt.Println(cars)
	sort.Sort(ByHorsepower(cars))
	
	fmt.Println(cars)
	sort.Sort(ByTopspeed(cars))
	
	fmt.Println(cars)
}