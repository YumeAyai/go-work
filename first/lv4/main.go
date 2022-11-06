package main

import (
	"fmt"
	"html/template"
	"io"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

var value string
var numin int
var secretNumber int
var mes string

func gennum() {
	maxNum := 100
	rand.Seed(time.Now().UnixNano())
	secretNumber = rand.Intn(maxNum)
	// 作弊模式
	fmt.Println("The secret number is ", secretNumber)

	fmt.Println("Please input your guess")
	// 通过一个 for 循环实现一直猜数，直到猜中
	// _, err := fmt.Scanf("%s", &guess)
	// if err != nil {
	// 	fmt.Println("Invalid input. Please enter an integer value")
	// 	continue
	// }
}

func guessnum(guess int) {
	fmt.Println("You guess is", guess)
	if guess > secretNumber {
		fmt.Println("Your guess is bigger than the secret number. Please try again")
		mes = "Your guess is bigger than the secret number. Please try again!"
	} else if guess < secretNumber {
		fmt.Println("Your guess is smaller than the secret number. Please try again")
		mes = "Your guess is smaller than the secret number. Please try again!"
	} else {
		fmt.Println("Correct, you Legend!")
		mes = "Correct, you Legend!"
	}
}

func post() {
	tmpl := template.Must(template.ParseFiles("index.html"))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method != http.MethodPost {
			tmpl.Execute(w, nil)
			return
		}
		value = r.FormValue("value")
		if numin, err := strconv.Atoi(value); err == nil {
			guessnum(numin)
		}
		io.WriteString(w, mes)
	})
}

func Request(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, string(mes))
}

func main() {
	gennum()
	post()
	http.ListenAndServe("127.0.0.1:8080", nil)
}
