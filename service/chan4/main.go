package chan4

import (
	"fmt"
	"time"
)

// --------------------------- dataSyncExample ------------------
func dataSyncExample() {
	fmt.Println("非缓冲通道实现数据同步")
	c := make(chan int)
	done := make(chan string)
	go func(ch chan<- int, x int) {
		time.Sleep(time.Second * 5)
		ch <- x * x * x
	}(c, 10)
	go func(ch <-chan int) {
		data := <-ch
		fmt.Println(data)
		s := "dataSyncExample 程序结束"
		done <- s
	}(c)
	fmt.Println("dataSyncExample 程序开始")
	fmt.Println(<-done)
}

// --------------------------- timeOutExample ------------------
func timeOutExampleWork() <-chan int {
	ch := make(chan int)
	go func() {
		fmt.Println("timeOutExample 开始工作")
		time.Sleep(time.Second * 3)
		ch <- 0
	}()
	return ch
}

func timeOutExample() {
	fmt.Println("非缓冲通道实现超时控制")
	select {
	case <-timeOutExampleWork():
		fmt.Println("timeOutExample 任务在规定时间内完成！")
	case <-time.After(time.Second * 3):
		fmt.Println("timeOutExample 任务超时了！！！")
	}
}

// --------------------------- consumerExample ------------------

func producer(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func consumer(ch chan int, done chan string) {
	for i := 0; i < 5; i++ {
		go func(id int) {
			for {
				data, ok := <-ch
				if !ok {
					done <- "done"
					fmt.Printf("done: %v --- %v\n", id, data)
					return
				} else {
					fmt.Printf("消费者 %v，完成任务 %v\n", id, data)
					time.Sleep(time.Second * 2)
				}
			}
		}(i)

	}
}

func consumerExample() {
	fmt.Println("带缓冲通道实现生产者 - 消费者模型")
	done := make(chan string)
	ch := make(chan int, 10)
	go producer(ch)
	go consumer(ch, done)
	<-done
	fmt.Println("任务完成")
}

// --------------------------- concurrencyExample ------------------

func handleEvent(done chan string, task chan bool) {
	for i := 0; i < 10; i++ {
		task <- true
		go func(id int) {
			fmt.Printf("处理事件 %v\n", id)
			time.Sleep(time.Second * 1)
			<-task
			if id == 9 {
				done <- "done"
			}
		}(i)
	}

}

func concurrencyExample() {
	fmt.Println("带缓冲通道实现并发数量控制的示例")
	done := make(chan string)
	task := make(chan bool, 2) //并发数控制为2
	go handleEvent(done, task)
	<-done
	fmt.Println("任务完成")
}

func Main() {
	fmt.Println("chan4")
	fmt.Println("")
	dataSyncExample()
	fmt.Println("")
	timeOutExample()
	fmt.Println("")
	consumerExample()
	fmt.Println("")
	concurrencyExample()

	fmt.Println("")
	fmt.Println("channel 使用注意事项 ")
	fmt.Println("关闭一个 nil 通道或者一个已经关闭的通道将产生一个 panic。 ")
	fmt.Println("向一个已关闭的通道发送数据也将导致 panic。 ")
	fmt.Println("向一个 nil 通道发送数据或者从一个 nil 通道接收数据将使当前协程永久阻塞。 ")
}
