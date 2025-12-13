package method

import "sync"

func ParallelFor2[T any](array []T, foreach func(index int, item T) error) error {
	concurrency := len(array)
	errs := make([]error, concurrency)
	var wg sync.WaitGroup
	for i := 0; i < concurrency; i++ {
		begin := len(array) / concurrency * i
		end := len(array) / concurrency * (i + 1)
		if len(array)%concurrency != 0 && i == concurrency-1 { // add the tail to the last group
			end = len(array)
		}
		wg.Add(1)
		go func(group, from, to int) {
			defer wg.Done()
			for index := from; index < to; index++ {
				if err := foreach(index, array[index]); err != nil {
					errs[group] = err
				}
			}
		}(i, begin, end)
	}
	wg.Wait()

	for _, e := range errs {
		if e != nil {
			return e
		}
	}
	return nil
}
