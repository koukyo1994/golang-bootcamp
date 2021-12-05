// パッケージbankは一つの口座を持つ並行的に安全な銀行を提供する
package bank

var deposits = make(chan int)         // 入金額を送信する
var balances = make(chan int)         // 残高を受信する
var withdrawl = make(chan int)        // 引落額を送信する
var withdrawlResult = make(chan bool) // 引落結果を返す

func Deposit(amount int) { deposits <- amount }

func Balance() int { return <-balances }

func Withdraw(amount int) bool {
	withdrawl <- amount
	return <-withdrawlResult
}

func teller() {
	var balance int // balanceはtellerゴルーチンに閉じ込められている
	for {
		select {
		case amount := <-deposits:
			balance += amount
		case amount := <-withdrawl:
			if balance >= amount {
				balance -= amount
				withdrawlResult <- true
			} else {
				withdrawlResult <- false
			}
		case balances <- balance:
		}
	}
}

func init() {
	go teller() // モニターゴルーチンを開始する
}
