# Package

## Cond

```go
	c := sync.NewCond(&sync.Mutex{})
	queue := make([]interface{}, 0, 10)

	removeFromQueue := func(delay time.Duration) {
		time.Sleep(delay)
		c.L.Lock()
		queue = queue[1:]
		fmt.Println("Removed from queue: " + strconv.Itoa(len(queue)))
		c.L.Unlock()
		c.Signal()
	}

	for i := 0; i < 10; i++ {
		c.L.Lock()
		for len(queue) == 2 {
			c.Wait()
		}
		fmt.Println("Adding to queue: " + strconv.Itoa(len(queue)))
		queue = append(queue, struct{}{})
		go removeFromQueue(1 * time.Second)
		c.L.Unlock()
	}
```
```go
for len(queue) == 2 {
  c.Wait()
}
```
こちらの条件式が真になるまでは `removeFromQueue` 関数は実行されない
`Wait()`を呼び出し、条件のシグナルが送出されるまでmain goroutineは一時停止される
そのため実行結果は最初の2回のappendの後、順不同でremoveやaddingが繰り返され、最終的にはqueueから2つの要素が取り出される前に終了する
```zsh
Adding to queue: 0
Adding to queue: 1
Removed from queue: 1
Adding to queue: 1
Removed from queue: 1
Adding to queue: 1
Removed from queue: 1
Adding to queue: 1
Removed from queue: 1
Adding to queue: 1
Removed from queue: 1
Removed from queue: 0
Adding to queue: 0
Adding to queue: 1
Removed from queue: 1
Adding to queue: 1
Removed from queue: 1
Adding to queue: 1
```
> また、新たに1つのアイテムをキューに追加する前に少なくとも1つの要素がキューから取り出されるのを待ちます。
:::note info
シグナルが送出されるまでは main goroutine が一時停止されるため
:::
