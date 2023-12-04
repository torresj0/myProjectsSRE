package main


import (	
	"fmt"
	"sort"


)

type dados struct{
	Name string
	Age int
}

type forName []dados

func (ps forName) len()int{
	return len(ps)

}

func (ps forName) less(i, j int) bool { // if item in index I is smaller than the item in position J
	return ps(i).Name < ps(j).Name
	
}

func (ps forName) Swap(i, j int){
	ps[i], ps[j] = ps[j], ps[i]
}


func main (){
	children = []dados{
		{"Joao", 19},
		{"Maria", 13},
	}
	sort.Sort(forName(children))
	fmt.Println(children)
}
