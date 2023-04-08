# The sync Package
* [1. WaitGroup](#anchor1)
* [2. Mutex and RWMutex](#anchor2)
* [3. Cond](#anchor3)
* [4. Once](#anchor4)
* [5. Pool](#anchor5)

<a id="anchor1"></a>
## 1. WaitGroup
> WaitGroupはひとまとまりの並行処理があったとき、その結果を気にしない、
> もしくは他に結果を収集する手段がある場合に、それらの処理の完了を待つ手段として非常に有効です。  

### Sample code
```go
var wg sync.WaitGroup

wg.Add(1) // 引数に渡された整数だけカウンターを増やす
go func() {
	defer wg.Done() // カウンターを1つ減らす
	fmt.Println("1st goroutine sleeping...")
	time.Sleep(1)
}()

wg.Add(1)
go func() {
	defer wg.Done()
	fmt.Println("2nd goroutine sleeping...")
	time.Sleep(2)
}()

wg.Wait() // カウンターが0になるまでブロックする
fmt.Println("All goroutines complete.")
```
```go
wg.Add(n)
```
n個のgoroutineが起動したことを表している。  
```go
defer wg.Done()
```
Doneをdeferキーワードを使って呼び出す。  
これによって、たとえpanicになったとしても確実にWaitGroupに終了することを伝える。  
***
`Add`の呼び出しは監視対象のgoroutineの外で行うこと。そうしないと競合状態を引き起こしてしまう。  
goroutineはスケジュールされるタイミングに関して何も保証がないため、goroutineを開始する前に`Wait`の呼び出しが起きてしまう可能性がある。  

### Result
> **Note**  
> 最後の`All goroutines complete.`以外は順不同で表示される。
```zsh
2nd goroutine sleeping...
1st goroutine sleeping...
All goroutines complete.
```

<a id="anchor3"></a>
## 3. Cond

> ゴルーチンが待機したりイベントの発生を知らせるためのランデブーポイントです。  

### Sample code

10個の要素をqueueに追加し、最後の2要素をqueueから取り出す前に終了するプログラム  

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

しかし、以下の条件で`Wait`が呼び出されると、main goroutineは条件のシグナルが送出されるまで一時停止される  

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
