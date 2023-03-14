package main

import (
	"fmt"
	"time"
)

func task(name string) {
	for i := 0; i < 10; i++ {
		fmt.Println(name, ":", i)
		time.Sleep(time.Second)
	}
}

func thread() {
	go task("Tarefa 1") // com o go inicia uma nova thread T2
	go task("Tarefa 2") //T3
	task("Tarefa 3")    //T1
}

func reader(canal chan int) {
	for x := range canal {
		fmt.Println(x)
	}
}

func publish(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
}

func main_publish_reader() {
	canal := make(chan int) //canal de comunicação entre as threads

	go publish(canal)
	go reader(canal)
	time.Sleep(time.Second * 5)
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d \n", workerId, x)
		time.Sleep(time.Second)
	}
}

func main() {
	ch := make(chan int)
	qtsWorkers := 3

	for i := 0; i < qtsWorkers; i++ {
		go worker(i, ch)
	}

	for i := 0; i < 100000; i++ {
		ch <- i
	}
}
