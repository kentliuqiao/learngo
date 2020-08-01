package engine

// ConcurrentEngine 并发版爬虫
type ConcurrentEngine struct {
	Scheduler   Scheduler
	WorkerCount int
	ItemChan    chan Item
}

// Scheduler receives request from engine and then schedules the request to worker
type Scheduler interface {
	ReadtNotifier
	Submit(Request)
	WorkerChan() chan Request
	Run()
}

// ReadtNotifier ...
type ReadtNotifier interface {
	WorkerReady(chan Request)
}

// Run ...
func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()
	for i := 0; i < e.WorkerCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		if isDuplicate(r.URL) {
			// log.Printf("duplicate request: %s", r.URL)
			continue
		}
		e.Scheduler.Submit(r)
	}

	// profileCnt := 0
	for {
		res := <-out
		for _, item := range res.Items {
			// profileCnt++
			// log.Printf("got profile #%d: %v", profileCnt, item)
			go func(i Item) {
				e.ItemChan <- i
			}(item)
		}

		// uRL dedup
		for _, req := range res.Requests {
			if isDuplicate(req.URL) {
				continue
			}
			e.Scheduler.Submit(req)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, ready ReadtNotifier) {
	go func() {
		for {
			ready.WorkerReady(in)
			request := <-in
			result, err := worker(request)
			if err != nil {
				continue
			}
			out <- result
		}
	}()
}

var visitedURLs = make(map[string]bool)

func isDuplicate(url string) bool {
	if visitedURLs[url] {
		return true
	}
	visitedURLs[url] = true
	return false
}
