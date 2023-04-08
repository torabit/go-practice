<a id="anchor0"></a>
# The sync Package
* [1. WaitGroup](#anchor1)
* [2. Mutex and RWMutex](#anchor2)
* [3. Cond](#anchor3)
* [4. Once](#anchor4)
* [5. Pool](#anchor5)

<a id="anchor1"></a>
## 1. WaitGroup
> WaitGroupはひとまとまりの並行処理があったとき、その結果を気にしない、
> もしくは他に結果を収集する手段がある場合に、それらの処理の完了を待つ手段として非常に有効です。[[1]](#quote1)  



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

1個のgoroutineが起動したことを表している。  

```go
wg.Add(1)
```

Doneをdeferキーワードを使って呼び出す。  
これによって、たとえpanicになったとしても確実にWaitGroupに終了することを伝える。  

```go
defer wg.Done()
```

***

`Add`の呼び出しは監視対象のgoroutineの外で行うこと。そうしないと競合状態を引き起こしてしまう。  
goroutineはスケジュールされるタイミングに関して何も保証がないため、goroutineを開始する前に`Wait`の呼び出しが起きてしまう可能性がある。  

### Result

```zsh
2nd goroutine sleeping...
1st goroutine sleeping...
All goroutines complete.
```
> **Note**  
> 最後の`All goroutines complete.`以外は順不同で表示される。  

[Back to Top](#anchor0)  

<a id="anchor2"></a>
## 2. Mutex and RWMutex

> プログラム内のクリティカルセクションを保護する方法の1つです。...  
> Mutexは並行処理で安全な方法でこれらの共有リソースに対する排他的アクセスを提供しています。[[1]](#quote1)

### Mutex
### Sample code
```go
var count int
var lock sync.Mutex

increment := func() {
  lock.Lock() // クリティカルセクションの占有を要求
  defer lock.Unlock() // クリティカルセクションの処理が終了したことを通知
  count++
  fmt.Printf("Incrementing: %d\n", count)
}

decrement := func() {
  lock.Lock()
  defer lock.Unlock()
  count--
  fmt.Printf("Decrementing: %d\n", count)
}

var wg sync.WaitGroup

const numIncrements = 5
wg.Add(numIncrements)
for i := 0; i <= 5; i++ {
  go func() {
    defer wg.Done()
    increment()
  }()
}

const numDecrements = 5
wg.Add(numDecrements)
for i := 0; i <= 5; i++ {
  go func() {
    defer wg.Done() 
    decrement()
  }()
}

wg.Wait()

fmt.Println("Arithmetic complete")
```

`Mutex`は、排他的なアクセス(クリティカルセクション)を制御するために使用される。  
つまり、同時に1つのスレッドだけが`Mutex`をロックし、そのリソースにアクセスすることができる。  
ほかのスレッドは`Mutex`が解放されるのを待たなければならない。これにより競合状態を回避することを実現している。  

### RWMutex
> **Note**  
> RWMutexの説明では記述量を減らすためSample codeを省略します。

`RWMutex`は、複数のスレッドが同時に読み込み(共有アクセス)できるようにするための`Mutex`。  
読み込みアクセスはクリティカルセクションではないため、複数のスレッドが同時に`RWMutex`を読み込むことが可能。  
しかし、書き込みアクセスは`Mutex`同様クリティカルセクションの占有が必要のため、`RWMutex`を書き込むスレッドがロックを取得するまで、ほかのスレッドを待つ必要がある。  
つまり、`RWMutex`は読み込み多重化を可能にし、並列性を向上させる。  
`RWMutex`は読み込みが頻繁に行われるが、書き込みが比較的に稀なプログラムに有用。  
例えば、大量の読み込み処理を行うバックエンドとか……(経験がなさ過ぎて例が思い浮かびません！)

### 結論
`RWMutex`は、読み込み操作に対しては複数のスレッドが同時にアクセスできるようにし、
書き込み操作に対してはクリティカルセクションの占有を要求することで、並列性を向上させる。
書き込み操作が頻繁に行われる場合には、`Mutex`などの排他的ロックを使用するべき。
(クリティカルセクションがどのような処理を行っているかによる)  


[Back to Top](#anchor0)

<a id="anchor3"></a>
## 3. Cond

> ゴルーチンが待機したりイベントの発生を知らせるためのランデブーポイントです。[[1]](#quote1)  

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
> また、新たに1つのアイテムをキューに追加する前に少なくとも1つの要素がキューから取り出されるのを待ちます。[[1]](#quote1)  

> **Note**  
> シグナルが送出されるまでは main goroutine が一時停止されるため  

[Back to Top](#anchor0)

## 参考文献
<a id="quote1">[1]</a>
<cite>[Go言語による並行処理](https://www.oreilly.co.jp/books/9784873118468/)</cite>
