package benchmark

func calc(limit int) {
	// var by3 []int
	// var by5 []int
	// var by8 []int

	// by3 := make([]int, (limit/3)+1)
	// by5 := make([]int, (limit/5)+1)
	// by8 := make([]int, (limit/8)+1)

	by3 := make([]int, 0, (limit/3)+1)
	by5 := make([]int, 0, (limit/5)+1)
	by8 := make([]int, 0, (limit/8)+1)

	// by3 := make([]interface{}, (limit/3)+1)
	// by5 := make([]interface{}, (limit/5)+1)
	// by8 := make([]interface{}, (limit/8)+1)

	// by3 := make([]interface{}, 0, (limit/3)+1)
	// by5 := make([]interface{}, 0, (limit/5)+1)
	// by8 := make([]interface{}, 0, (limit/8)+1)

	// var by3 []interface{}
	// var by5 []interface{}
	// var by8 []interface{}

	for i := 0; i < limit; i++ {
		if i%3 == 0 {
			by3 = append(by3, i)
		}
		if i%5 == 0 {
			by5 = append(by5, i)
		}
		if i%8 == 0 {
			by8 = append(by8, i)
		}
	}

	// fmt.Println("By3:", len(by3), "\nby5:", len(by5), "\nby8:", len(by8))
}

func test() error {
	count, limit := 100, 999999
	for i := 0; i < count; i++ {
		calc(limit)
	}
	return nil
}

func ExampleBenchmark() {
	// Returns a new Benchmark pointer with all the defaults assigned
	b := New()
	// time to wait before firing the consequent request
	// WaitPerReq = time.Millisecond * 1
	// print available stats while the benchmark is running
	b.ShowProgress = true
	// Total number of requests to fire
	b.TotalRequests = 10
	// Duration in which all the requests have to be finished firing (in milliseconds).
	b.BenchDuration = 7300

	// Updates all the necessary fields according to the configuration provided
	b.Init()
	// Run the benchmark
	b.Run(test)
}
