package main

import (
	"encoding/json"
	"github.com/bytedance/sonic"
	"testing"
)

type Student struct {
	Id  string
	age int
	sex bool
}

type Class struct {
	Id       string
	Students []Student
}

var (
	s = Student{
		Id:  "21",
		age: 18,
		sex: true,
	}
	c = Class{
		Id:       "5",
		Students: []Student{s, s, s},
	}
)

func BenchmarkJson(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, _ := json.Marshal(c)
		var c2 Class
		_ = json.Unmarshal(bytes, &c2)
	}
}

func BenchmarkSonic(b *testing.B) {
	for i := 0; i < b.N; i++ {
		bytes, _ := sonic.Marshal(c)
		var c2 Class
		_ = sonic.Unmarshal(bytes, &c2)
	}
}
