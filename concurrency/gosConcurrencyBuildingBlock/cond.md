## Cond

### Sample code

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

普通、`removeFromQueue`の実行よりもmain goroutine のfor文の実行スピードの方が速い  
しかし、以下の条件文で`Wait`が呼び出されると、main goroutineは条件のシグナルが送出されるまで一時停止される  
 
```go
for len(queue) == 2 {
  c.Wait()
}
```

そのため実行結果は最初の2回のappendの後、順不同でremoveやaddingが繰り返される  

最終的にはqueueから2つの要素が取り出される前に終了する 

### Result

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

> **Note**  
> シグナルが送出されるまでは main goroutine が一時停止されるため
