package ga

type Record struct {
	Timestamp int64
	UserId    int
}

func NewRecord(timestamp int64, userId int) *Record {
	return &Record{
		Timestamp: timestamp,
		UserId:    userId,
	}
}

type ByTimestamp []*Record

func (self ByTimestamp) Len() int {
	return len(self)
}

func (self ByTimestamp) Swap(i, j int) {
	self[i], self[j] = self[j], self[i]
}

func (self ByTimestamp) Less(i, j int) bool {
	return self[i].Timestamp < self[j].Timestamp
}

type Records []*Record

// Returns number of unique userIds in records
func (self Records) UniqueUserIds() int {
	userIdsMap := make(map[int]bool)
	uniqUserIds := 0

	for i := 0; i < len(self); i++ {

		// if map does not contain result -> it's uniq, increase counter
		if _, found := userIdsMap[self[i].UserId]; !found {
			userIdsMap[self[i].UserId] = true
			uniqUserIds++
		}
	}

	return uniqUserIds
}
