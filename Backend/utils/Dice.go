package utils

import (
	"math/rand"
	"time"
)

const (
	Dice4   = 4
	Dice8   = 8
	Dice10  = 10
	Dice12  = 12
	Dice20  = 20
	Dice24  = 24
	Dice30  = 30
	Dice100 = 100
)

// 4 8 10 12 20 24 30 100面的骰子
func Dice(num, side int) []int {
	// 初始化随机数种子
	rand.Seed(time.Now().UnixNano())

	// 创建一个长度为num的切片用于存储结果
	result := make([]int, num)

	// 投掷num次骰子，每次的结果是1到side之间的随机数
	for i := 0; i < num; i++ {
		result[i] = rand.Intn(side) + 1
	}
	return result
}
func Chaos(dice []int) (sum int) {
	//如果dice里面有三个三 sum为0
	trible := 0
	for i := 0; i < len(dice); i++ {
		if dice[i] != 3 {
			sum += dice[i]
		}
		if dice[i] == 3 {
			trible++
		}
	}
	if trible == 3 { //三重升华
		return 0
	}
	return sum
}
