package main
 
import (
	"fmt"
	"os"
	"os/signal"
	"time"
	"log"
)

type myStruct struct {
	name string
	num int
}

type inputStruct struct {
	input string
	size int
}

func (out *myStruct) method(in *inputStruct) {
	out.num = in.size
}

func StopCircle(ch chan bool) {
	ch <- false
}
 
func main() {
	var a myStruct
	var b inputStruct
	b.input = "4123"
	b.size = 50
	a.method(&b)
	println(a.num)
	c := make(chan os.Signal, 1)
	signal.Notify(c)
 
	ticker := time.NewTicker(time.Second)
	done := make(chan bool)
	go func() {
		for {
			select {
			case <-done:
				goto last
			case <-ticker.C:
				log.Println("Hey!")
			}
		}
	last:
		fmt.Println("after loop")
	}()

	time.Sleep(3 * time.Second)
	ticker.Stop()
	// done <- true
	StopCircle(done)
	log.Println("Done")
	// go func() {
	// 	defer func() { stop <- true }()
	// 	for t := range ticker.C{
	// 		fmt.Printf("tick at: %v\n", t.UTC())

	// 		// select {
	// 		// case <-ticker.C:
	// 		// 	fmt.Println("Тик")
	// 		// case <-stop:
	// 		// 	fmt.Println("Закрытие горутины")
	// 		// 	return
	// 		// }
	// 	}
	// }()
 
	// Блокировка, пока не будет получен сигнал
	ticker.Stop()
	fmt.Println("Приложение остановлено")
}