package main

import (
	"fmt"
	"strings"
	"sync"
)

func waitGroupExample(gorutines ...func() string) string {
	var res string
	var wg sync.WaitGroup
	var mu sync.Mutex

	for i, goroutine := range gorutines {
		wg.Add(1)
		go func(i int, gorutine func() string) {
			defer wg.Done()
			_ = gorutine()
			mu.Lock()
			res += fmt.Sprintf("gorutine %d done\n", i+1)
			mu.Unlock()
		}(i, goroutine)
	}

	wg.Wait()
	return res
}

func waitGroupExample2(gorutines ...func() string) string {

	var wg sync.WaitGroup
	var sb strings.Builder

	for i, goroutine := range gorutines {
		wg.Add(1)
		go func(i int, gorutine func() string) {
			defer wg.Done()
			_ = gorutine()
			sb.WriteString(fmt.Sprintf("gorutine %d done\n", i+1))
		}(i, goroutine)
	}

	wg.Wait()
	return sb.String()
}

func main() {
	count := 1000
	gorutines := make([]func() string, count)

	for i := 0; i < count; i++ {
		j := i
		gorutines[j] = func() string {
			return fmt.Sprintf("gorutine %d", j)
		}
	}
	fmt.Println(len(waitGroupExample(gorutines...)) == len(waitGroupExample2(gorutines...)))
	fmt.Println(waitGroupExample(gorutines...))

}
