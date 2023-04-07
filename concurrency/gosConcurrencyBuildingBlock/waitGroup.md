## WaitGroup

### Sample code
```go

var wg sync.WaitGroup

wg.Add(1)
go func() {
	defer wg.Done()
	fmt.Println("1st goroutine sleeping...")
	time.Sleep(1)
}()

wg.Add(1)
go func() {
	defer wg.Done()
	fmt.Println("2nd goroutine sleeping...")
	time.Sleep(2)
}()

wg.Wait()
fmt.Println("All goroutines complete.")
```
Addを呼び出すと引数に渡された回数分カウンターを増やし、Doneを呼び出すとカウンターを1つ減らす
Waitを呼び出すとカウンターが0になるまでブロックする

Addの呼び出しは監視対象のgoroutineの外で行うこと
そうしないと競合状態を引き起こしてしまう

goroutineはスケジュールされるタイミングに関して、何も保証がないため、goroutineを開始する前にWait()の呼び出しが起きてしまう可能性がある。
