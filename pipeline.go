package pipeline

import "sync"

func sub(in, out chan interface{}, f func(in, out chan interface{}), wg *sync.WaitGroup) {
	f(in, out)
	wg.Done()
	close(out)
}

func Run(funcs ...func(in, out chan interface{})) {
	var wg sync.WaitGroup
	out := make(chan interface{})
	for i := len(funcs) - 1; i >= 0; i-- {
		in := make(chan interface{})
		wg.Add(1)
		go sub(in, out, funcs[i], &wg)
		out = in
	}
	close(out)
	wg.Wait()
}
