package EraserProcessor

import (
	"fmt"
	"os"
	"sync"
)

func EraserOperator(l []string) error {
	var err error
	chDone := make(chan string)
	chErr := make(chan error)
	wg := sync.WaitGroup{}
	for i := range l {
		wg.Add(1)
		go DeleteFunc(&wg, l[i], chDone, chErr)
	}
	go func() {
		for {
			select {
			case inf := <-chDone:
				fmt.Println(inf)
			case errData := <-chErr:
				err = errData
				break
			}
		}
	}()

	wg.Wait()
	return err
}

func DeleteFunc(w *sync.WaitGroup, path string, chW chan<- string, chE chan<- error) {
	defer w.Done()
	if err := os.Remove(path); err != nil {
		chE <- err
	}
	chW <- path + " was successfully deleted"
}
