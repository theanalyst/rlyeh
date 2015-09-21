package main

type PerfCounter struct {
	Filestore struct {
		JournalLatency struct {
			AvgCount int     `json:"avgcount"`
			Sum      float64 `json:"sum"`
		}
		Ops int `json:"ops"`
	} `json:"filestore"`
	Leveldb struct {
		LeveldbGetLatency struct {
			AvgCount int     `json:"avgcount"`
			Sum      float64 `json:"sum"`
		}
	} `json:"leveldb"`
}
