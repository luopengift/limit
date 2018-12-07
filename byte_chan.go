package limit

// BytesChan bytes chan
type BytesChan chan []byte

func worker(max int, job func() error) {
	for i := 0; i < max; i++ {
		job()
	}
}
