package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	url := "https://httpbin.org/get"
	parallelRequest := 5
	requestCount := 50
	result := benchRequest(url, parallelRequest, requestCount)

	for i := 0; i < requestCount; i++ {
		statusCode := <-result
		if statusCode != 200 {
			panic(fmt.Sprintf("Ошибка при отправке запроса: %d", statusCode))
		}
	}
	fmt.Println("Все горутины завершили работу")
}
func benchRequest(url string, parallelRequest int, requestCount int) <-chan int {
	limiter := make(chan struct{}, parallelRequest)
	result := make(chan int, requestCount)
	var wg sync.WaitGroup

	// Запускаем горутины
	for i := 0; i < requestCount; i++ {
		limiter <- struct{}{}
		wg.Add(1)
		go func() {
			defer wg.Done()
			statusCode, err := httpRequest(url)
			if err != nil {
				fmt.Println(err)
			}
			result <- statusCode
			<-limiter
		}()
	}

	// Закрываем канал результатов после завершения всех горутин
	go func() {
		wg.Wait()
		close(result)
	}()

	return result
}

func httpRequest(url string) (int, error) {
	resp, err := http.Get(url)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()
	return resp.StatusCode, nil
}
