package util

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"sync"
	"time"
)

const saltBit = uint(8)                  // 随机因子二进制位数
const saltShift = uint(8)                // 随机因子移位数
const increasShift = saltBit + saltShift // 自增数移位数

type Mist struct {
	sync.Mutex       // 互斥锁
	increas    int64 // 自增数
	saltA      int64 // 随机因子一
	saltB      int64 // 随机因子二
	saltC      int64 // 随机因子二

}

/* 初始化 Mist 结构体*/
func newMist() *Mist {
	mist := Mist{increas: 1}
	return &mist
}

/* 生成唯一编号 */
func (c *Mist) generate() uint64 {
	c.Lock()
	c.increas++
	// 获取随机因子数值 ｜ 使用真随机函数提高性能
	randA, _ := rand.Int(rand.Reader, big.NewInt(1000))
	c.saltA = randA.Int64()
	randB, _ := rand.Int(rand.Reader, big.NewInt(2000))
	c.saltB = randB.Int64()
	c.saltC = time.Now().Unix()
	//fmt.Println(c)
	// 通过位运算实现自动占位
	mist := uint64((c.increas << increasShift) | (c.saltA << saltShift) | c.saltB | c.saltC)
	c.Unlock()
	return mist
}

func GetMistID() {
	// 使用方法
	//mist := newMist()
	j := 1667235309
	for i := 0; i < 20; i++ {
		j = j + 1
		fmt.Println(j)
	}
}
