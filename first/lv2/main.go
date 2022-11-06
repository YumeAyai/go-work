package main

import "fmt"

type Movie struct {
	Name string
	Star float32
	Age  string
	Lang string
}

func main() {
	m := Movie{Name: "西线无战事", Star: 8.8, Age: "All ages", Lang: "English"}
	fmt.Printf(`请输入你的命令
	1.获得名字
	2.查看评分
	3.查看分级
	4.查看语言
	5.退出程序
	`)
	var option int
	for {
		fmt.Scanf("%d", &option)
		switch option {
		case 1:
			fmt.Println(m.Name)
		case 2:
			fmt.Println(m.Star)
		case 3:
			fmt.Println(m.Age)
		case 4:
			fmt.Println(m.Lang)
		case 5:
			fmt.Print("Exit")
			return
		}
	}
}
