# Go's Concurrency Building Blocks

## Goroutine

### Goroutineの中でClosureを実行するとどうなるか
```go
	var wg sync.WaitGroup
	salutationList := []string{}
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		// Goroutineが起動する前にLoop処理が終了されてしまう
		go func() {
			defer wg.Done()
			// そのため、変数salutationはスコープ外になってしまう
			// しかし、GoのランタイムがGoroutineがメモリにアクセスし続けられるように、メモリをヒープへ移す
			salutationList = append(salutationList, salutation)
		}()
	}
	wg.Wait()
```
実行結果は以下のようになる
```zsh
good day
good day
good day
```
このループを正しく動作させるためには、salutationのコピーをclosureに渡す必要がある
```go
	var wg sync.WaitGroup
	salutationList := []string{}
	for _, salutation := range []string{"hello", "greetings", "good day"} {
		wg.Add(1)
		go func(salutation string) {
			defer wg.Done()
			salutationList = append(salutationList, salutation)
		}(salutation)
	}
	wg.Wait()
	return salutationList
```
この結果、次のように正しい出力を得る
```zhs
good day
hello
greetings
```
