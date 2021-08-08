package main

import (
	"fmt"
	"sync"
)

type singleton struct {
	count int
	sync.RWMutex
}

var instance singleton

func GetInstance() *singleton {
	return &instance
}

func (s *singleton) AddOne() {
	s.Lock()
	defer s.Unlock()
	s.count++
}

func (s *singleton) GetCount() int {
	s.RLock()
	defer s.RUnlock()
	return s.count
}

func main() {
	obj1 := GetInstance()
	obj1.AddOne()
	fmt.Println(obj1.GetCount())
	obj2 := GetInstance()
	obj2.AddOne()
	fmt.Println(obj2.GetCount())
	obj3 := GetInstance()
	obj3.AddOne()
	fmt.Println(obj3.GetCount())
}
