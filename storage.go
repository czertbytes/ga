package ga

import (
	"sort"
	"sync"
)

type Storage struct {
	sync.Mutex
	Records Records
}

func NewStorage() *Storage {
	return &Storage{
		Records: Records{},
	}
}

// Inserts new record in sorted records storage
func (self *Storage) Insert(record *Record) {
	self.Lock()
	defer self.Unlock()

	self.Records = append(self.Records, record)
	length := len(self.Records)

	if length == 1 {
		return
	}

	// Built-in sort (mix of insert, heap, quick sorts) - slow
	//sort.Sort(ByTimestamp(self.Records))

	// Custom insert sort - faster
	self.sort()
}

// Simple Insert Sort implementation
func (self *Storage) sort() {
	length := len(self.Records)

	// do insert sort only for last added item
	last := self.Records[(length - 1)]
	j := length - 2

	for j >= 0 && self.Records[j].Timestamp > last.Timestamp {
		self.Records[j+1] = self.Records[j]
		j--
	}
	self.Records[j+1] = last
}

// Returns records in timestamp interval <start,end)
func (self *Storage) Find(start, end int64) Records {
	self.Lock()
	defer self.Unlock()

	length := len(self.Records)

	startIndexChan := make(chan int)
	endIndexChan := make(chan int)

	findIndex := func(value int64, resChan chan int) {
		// use binary search to find index
		index := sort.Search(length, func(i int) bool {
			return self.Records[i].Timestamp >= value
		})

		resChan <- index
	}

	// find startIndex and endIndex in parallel
	go findIndex(start, startIndexChan)
	go findIndex(end, endIndexChan)

	return self.Records[<-startIndexChan:<-endIndexChan]
}
