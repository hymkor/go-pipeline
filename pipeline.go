package pipeline

import "sync"

func sub[T any](in, out chan T, f func(in, out chan T), wg *sync.WaitGroup) {
	f(in, out)
	wg.Done()
	close(out)
}

func Run[T any](funcs ...func(in, out chan T)) {
	var wg sync.WaitGroup
	out := make(chan T)
	for i := len(funcs) - 1; i >= 0; i-- {
		in := make(chan T)
		wg.Add(1)
		go sub(in, out, funcs[i], &wg)
		out = in
	}
	close(out)
	wg.Wait()
}
