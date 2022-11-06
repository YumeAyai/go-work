package main

import (
	"fmt"
)

type Heroer interface {
	Say()
	Atk()
	Skill()
	Item()
}
type Hero struct {
	HP int
}

func (Y Hero) Say() {
	fmt.Println("I'm Yasuo")
}
func (Y Hero) Atk() {
	fmt.Println("你对敌人造成了1的伤害")
}
func (Y Hero) Skill() {
	fmt.Println("然而没什么卵用")
}
func (Y Hero) Item() {
	fmt.Println("然而没什么卵用")
}

func main() {
	fmt.Printf("获取英雄列表中\n")
	var hero Heroer
	var option int
	for {
		//ls()
		fmt.Printf("选择你的英雄\n\n")
		//fmt.Scanf("%d", &option)
		hero = new(Hero)
		hero.Say()
		for {
			fmt.Printf("\n\n1.发动攻击\n2.使用技能\n3.使用道具\n4.使用外挂\n5.退出程序\n")
			fmt.Scanf("%d", &option)
			switch option {
			case 1:
				hero.Atk()
			case 2:
				hero.Skill()
			case 3:
				hero.Item()
			case 4:
				//hero.Cheat()
				fmt.Printf("你在想P吃")
			case 5:
				return
			}
		}
	}
}
