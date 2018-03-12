package engine

type ConcurrentEngine struct {
	Scheduler Scheduler
	WorkCount int
	ItemChan  chan Item
}

type Scheduler interface {
	ReadyNotifier
	Submit(Request)
	//ConfigureMasterWorkerChan(chan Request)
	WorkerChan() chan Request
	Run()
}

type ReadyNotifier interface {
	WorkerReady(chan Request)
}

func (e *ConcurrentEngine) Run(seeds ...Request) {

	out := make(chan ParseResult)
	e.Scheduler.Run()

	for i := 0; i < e.WorkCount; i++ {
		createWorker(e.Scheduler.WorkerChan(), out, e.Scheduler)
	}

	for _, r := range seeds {
		e.Scheduler.Submit(r)
	}

	//itemCount := 0
	for {
		result := <-out
		//log.Printf("receiveing parse result %v", result)
		for _, item := range result.Items {
			//itemCount++
			//log.Printf("Got item %d: %v\n", itemCount, item)
			go func() { e.ItemChan <- item }()
		}
		for _, request := range result.Requests {
			if isVisitedRequest(&request) {
				continue
			}
			e.Scheduler.Submit(request)
		}
	}
}

func createWorker(in chan Request, out chan ParseResult, s ReadyNotifier) {
	go func() {
		for {
			s.WorkerReady(in)
			request := <-in
			parseResult, err := worker(request)
			if err != nil {
				continue
			}
			//log.Printf("pushing parse result %v", parseResult)
			out <- parseResult
			//go func() { out <- parseResult }()
		}
	}()
}

var visitedUrls = make(map[string]bool)

func isVisitedRequest(request *Request) bool {
	if visitedUrls[request.Url] {
		return true
	}
	visitedUrls[request.Url] = true
	return false
}
