package main

import "fmt"

type Human struct {
	Name string
	Age  int
}

func (h *Human) SayName() {
	fmt.Println(h.Name)
}

func (h *Human) SayAge() {
	fmt.Println("Этот метод 'SayAge()' не переопределен:", h.Age)
}


type Action struct {
	Human // встаивание
}

func (a *Action) SayName() {
	fmt.Println("Этот метод 'SayName()' переопределен встаиванием в структуру Action :", a.Human.Name)
}


func main() {
	Bob := Human{Name: "Bob", Age: 20}
	Act := Action{Human: Bob}
	Act.SayName() // переопределен
	Act.SayAge() // не переопределен но доступен без указания структуры (напрямую)

}
