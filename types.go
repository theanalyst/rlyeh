package main

type PerfCounter struct {
	osd_id    string
	Filestore struct {
		JournalLatency struct {
			AvgCount int     `json:"avgcount"`
			Sum      float64 `json:"sum"`
		} `json:"journal_latency"`
		Ops int `json:"ops"`
	} `json:"filestore"`
	Leveldb struct {
		LeveldbGetLatency struct {
			AvgCount int     `json:"avgcount"`
			Sum      float64 `json:"sum"`
		} `json:"leveldb_get_latency"`
	} `json:"leveldb"`
	OSD struct {
		op_in_bytes float64 `json:"op_in_bytes"`
		OpLatency   struct {
			AvgCount int     `json:"avgcount"`
			Sum      float64 `json:"sum"`
		} `json:"op_latency"`
	} `json:"osd"`
}

type MonStats struct {
	Cluster struct {
		NumObject          int `json:"num_object"`
		NumObjectDegraded  int `json:"num_object_degraded"`
		NumObjectMisplaced int `json:"num_object_misplaced"`
		NumObjectUnfound   int `json:"num_object_unfound"`
		NumOsd             int `json:"num_osd"`
		NumOsdIn           int `json:"num_osd_in"`
		NumOsdUp           int `json:"num_osd_up"`
		NumPg              int `json:"num_pg"`
		NumPgActive        int `json:"num_pg_active"`
		NumPgActiveClean   int `json:"num_pg_active_clean"`
		NumPgPeering       int `json:"num_pg_peering"`
		NumPool            int `json:"num_pool"`
		OsdBytes           int `json:"osd_bytes"`
		OsdBytesAvail      int `json:"osd_bytes_avail"`
		OsdBytesUsed       int `json:"osd_bytes_used"`
	} `json:"cluster"`
}

type Meter struct {
	Counter_type   string  `json:"counter_type"`
	Counter_name   string  `json:"counter_name"`
	Resource_id    string  `json:"resource_id"`
	Counter_unit   string  `json:"counter_unit"`
	Counter_volume float64 `json:"counter_volume"`
}
