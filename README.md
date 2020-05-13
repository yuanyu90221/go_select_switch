# go_select_switch

## select vs switch

## switch 
switch 可以對各種型別執行輪詢 判斷型別
重點是依序執行

```golang===
func convert(i interface{}) {
	switch t := i.(type) {
	case int:
		println("i is interger", t)
	case string:
		println("i is string", t)
	case float64:
		println("i is float64", t)
	default:
		println("type not found")
	}
}

```
## select
1 select 只能拿來做 channel相關操作
    default會直接執行
    而如果沒有default 則會blocking
    而完全沒有接受value的channel會有panic
2 select 是隨機選取的 並沒有固定順序跑 case

```golang==
// channel
	ch := make(chan int, 1)

	// ch <- 1
	select {
	case <-ch:
		fmt.Println("random 01")
	case <-ch:
		fmt.Println("random 02")
	default:
		fmt.Println("exit")
	}
```

## select的timeout機制

1 做一個timeout channel 透過一個go runtine在時間到時 寫入訊息到 timeout channel

2 使用time.After function 因為 time.After 回傳值為channel

因此可以直接讀出 time.After channel內容

```golang==
    timeout := make(chan bool, 1)
	go func() {
		time.Sleep(2 * time.Second)
		timeout <- true
	}()
	ch1 := make(chan int)
	select {
	case <-ch1:
	case <-timeout:
		fmt.Println("timeout 01")
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 02")
	}
```
## buffered channel 如果 buffer 被填滿 case <-chan 後的 statement 不會被執行

```golang===
   ch2 := make(chan int, 1)
	ch2 <- 1
	select {
	case ch2 <- 2: // if buffered filled  this will not executed
		fmt.Println("channel value is", <-ch2)
		fmt.Println("channel value is", <-ch2)
	default:
		fmt.Println("channel blocking")
	}
```

## 如果要一直讀取不同channel直到有特定條件完成 可以利用 for select 配合 break loop語法 

***Notice***  如果不希望block需要把 for loop放在一個goroutine內
***Notice***  Defer是keyword代表 最後 return 前才會執行 這邊用來執行close(ch)
```golang===
    i := 0
	ch3 := make(chan string, 0)
	defer func() {
		close(ch)
	}()

	go func() {
	LOOP:
		for {
			time.Sleep(1 * time.Second)
			fmt.Println(time.Now().Unix())
			i++

			select {
			case m := <-ch3:
				println(m)
				break LOOP
			default:
			}
		}
	}()
	time.Sleep(time.Second * 4)
	ch3 <- "stop"
```