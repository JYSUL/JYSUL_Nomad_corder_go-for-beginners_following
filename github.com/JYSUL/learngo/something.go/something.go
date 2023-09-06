package something

import "fmt"

func sayBye() { // 함수를 export할때는 가장 앞의 알파벳을 대문자로 하면 자동으로 export
	fmt.Println("Bye")
}

func SayHello() {
	fmt.Println("Hello")
}
