/* package main

import (
	"fmt"
	"time"
)

func test1() {
	now := time.Now()
	fmt.Printf("t: %T\n", now)
	fmt.Printf("t: %v\n", now)

	year := now.Year()     //年
	month := now.Month()   //月
	day := now.Day()       //日
	hour := now.Hour()     //小时
	minute := now.Minute() //分钟
	second := now.Second() //秒
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n", year, month, day, hour, minute, second)
	fmt.Printf("%T,%T,%T,%T,%T,%T,%T\n", now, year, month, day, hour, minute, second)

}
func test2() {
	now := time.Now()
	// 当前时间戳 TimeStamp type:int64, TimeStamp:1606832965
	fmt.Printf("TimeStamp type:%T, TimeStamp:%v", now.Unix(), now.Unix())
}

func add(h, m, s, mls, msc, ns time.Duration) {
	now := time.Now()
	fmt.Println(now.Add(time.Hour*h + time.Minute*m + time.Second*s + time.Millisecond*mls + time.Microsecond*msc + time.Nanosecond*ns))
}

func test() {
	t := time.Now()
	fmt.Println(t.Format("2006/01/02 15:04"))
}

func main() {
	// test1()
	// test2()
	test()
}
 */