package main
import "fmt"

func main() {
	// for 循环
	for i := 0 ; i < 10 ; i++ {
		fmt.Println(i);
	}

    // while 循环
	var pin int = 1;
	for pin < 10 {
		if pin > 3 {
			fmt.Println(pin);
		}
		pin += 1
	}

    // 死循环
	for {

	}
}