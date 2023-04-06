# Go's Concurrency Building Blocks

## Goroutine

### Goroutineの中でClosureを実行するとどうなるか
```go
	var wg sync.WaitGroup
	salutationList := []string{}
	for _, salutation := range []string{"hello", "greeting", "good day"} {
		wg.Add(1)
		// Goroutineが起動する前にLoop処理が終了されてしまう
		go func() {
			defer wg.Done()
			// そのため、変数salutationはスコープ外になってしまう
			// しかし、GoのランタイムがGoroutineがメモリにアクセスし続けられるように、メモリをヒープへ移す
			// その結果、最後の値である"good day"を参照し、3度表示されるようになる
			salutationList = append(salutationList, salutation)
		}()
	}
	wg.Wait()
```
```go
	var wg sync.WaitGroup
	salutationList := []string{}
	for _, salutation := range []string{"hello", "greeting", "good day"} {
		wg.Add(1)

		go func(salutation string) {
			defer wg.Done()
			salutationList = append(salutationList, salutation)
		}(salutation)
	}
	wg.Wait()
	return salutationList
```
