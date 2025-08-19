package src

import (
	"fmt"
	"sync"
	"testing"
)

type Wallet struct {
	accountBalance map[string]int
	mutex          sync.RWMutex
}

func (wallet *Wallet) AddBalance(wg *sync.WaitGroup, account string, amount int) {
	defer wg.Done()
	wg.Add(1)

	wallet.mutex.Lock()
	wallet.accountBalance[account] += amount
	wallet.mutex.Unlock()
}

func (wallet *Wallet) DeductBalance(wg *sync.WaitGroup, account string, amount int) {
	defer wg.Done()
	wg.Add(1)

	wallet.mutex.Lock()
	wallet.accountBalance[account] -= amount
	wallet.mutex.Unlock()
}

func (wallet *Wallet) GetBalance(wg *sync.WaitGroup, account string) int {
	defer wg.Done()
	wg.Add(1)

	wallet.mutex.RLock()
	balance := wallet.accountBalance[account]
	wallet.mutex.RUnlock()
	return balance
}

func TestMutex(t *testing.T) {
	wg := &sync.WaitGroup{}
	wallet := Wallet{}
	wallet.accountBalance = map[string]int{}

	for i := 0; i < 10; i++ {
		go func() {
			for j := 0; j < 10; j++ {
				if j%2 == 0 {
					wallet.AddBalance(wg, "Andre", j)
					fmt.Println("Current Balance Andre: ", wallet.GetBalance(wg, "Andre"))
				} else {
					wallet.AddBalance(wg, "Budi", j)
					fmt.Println("Current Balance Budi: ", wallet.GetBalance(wg, "Budi"))
				}
			}
		}()
	}

	wg.Wait()

	fmt.Println("Account Balance Andre: ", wallet.GetBalance(wg, "Andre"))
	fmt.Println("Account Balance Budi: ", wallet.GetBalance(wg, "Budi"))

	wallet.DeductBalance(wg, "Andre", 10)
	wallet.DeductBalance(wg, "Budi", 10)

	fmt.Println("Final Account Balance Andre: ", wallet.GetBalance(wg, "Andre"))
	fmt.Println("Final Account Balance Budi: ", wallet.GetBalance(wg, "Budi"))

}
