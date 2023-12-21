package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// В момент старта запускает 10 горутин,
// каждая из которых с промежутком от 0.5 секунд до 1 секунды
// зачисляет на счёт клиента случайную сумму от 1 до 10.

// Так же запускается 5 горутин,
// которые с промежутком 0.5 секунд до 1 секунды
// снимают с клиента случайную сумму от 1 до 5.
// Если снятие невозможно, в консоль выводится сообщение об ошибке, и приложение продолжает работу.
func main() {
	var wg sync.WaitGroup
	run := true
	runLock := sync.RWMutex{}
	cl := NewBankClient()
	wg.Add(1)
	go func() {
		defer wg.Done()
		innerRun := true
		for innerRun {
			var s string
			_, err := fmt.Scanln(&s)
			if err != nil {
				fmt.Println(err)
				continue
			}
			switch s {
			case "balance":
				fmt.Println(cl.Balance())

			case "exit":
				fmt.Println("breaking operations...")
				runLock.Lock()
				run = false
				innerRun = false
				runLock.Unlock()

			case "deposit":
				fmt.Println("please enter deposit")
				var deposit int
				_, err = fmt.Scanln(&deposit)
				if err != nil {
					fmt.Println(err)
					continue
				}
				cl.Deposit(deposit)
			case "withdrawal":
				fmt.Println("please enter widtdrawal summ")
				var amount int
				_, err = fmt.Scanln(&amount)
				if err != nil {
					fmt.Println(err)
					continue
				}
				err = cl.Withdrawal(amount)
				if err != nil {
					fmt.Println(err)
					continue
				}
			default:
				fmt.Println("Unsupported command. You can use commands: balance, deposit, withdrawal, exit")
			}
		}
		fmt.Println("closing account...")
	}()

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				vr := rand.Intn(500) + 500
				time.Sleep(time.Duration(vr) * time.Millisecond)
				ir := rand.Intn(9) + 1
				cl.Deposit(ir)
				runLock.RLock()
				breaking := !run
				runLock.RUnlock()
				if breaking {
					break
				}
			}

		}()
	}

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				vr := rand.Intn(500) + 500
				time.Sleep(time.Duration(vr) * time.Millisecond)
				ir := rand.Intn(4) + 1
				res := cl.Withdrawal(ir)
				if res != nil {
					fmt.Println(res)
				}

				runLock.RLock()
				breaking := !run
				runLock.RUnlock()
				if breaking {
					break
				}
			}

		}()
	}

	wg.Wait()

	fmt.Println("Done.")
}
