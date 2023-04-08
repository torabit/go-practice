## WaitGroup
> WaitGroupはひとまとまりの並行処理があったとき、その結果を気にしない、
> もしくは他に結果を収集する手段がある場合に、それらの処理の完了を待つ手段として非常に有効です。  

### Sample code
```go

var wg sync.WaitGroup

wg.Add(1) // 引数に渡された回数分カウンターを増やす
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

`Add`の呼び出しは監視対象のgoroutineの外で行うこと、そうしないと競合状態を引き起こしてしまう  
goroutineはスケジュールされるタイミングに関して何も保証がないため、  
goroutineを開始する前に`Wait`の呼び出しが起きてしまう可能性がある。  

### Result
> **Note**  
> 最後の`All goroutines complete.`以外は順不同で表示される
```zsh
2nd goroutine sleeping...
1st goroutine sleeping...
All goroutines complete.
```
