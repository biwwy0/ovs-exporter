package ovs

//This file contains the JSON response object (structs)

//Written by Megyo @ LeanNet

type Flow struct {
	Cookie      string  `json:"cookie"`
	Duration    float64 `json:"duration"`
	Table       string  `json:"table"`
	Packets     int     `json:"packets"`
	Bytes       int     `json:"bytes"`
	IdleAge		int     `json:"idle_age"`
	HardAge		int     `json:"hard_age"`
	Priority    string  `json:"priority"`
	Match       string  `json:"match"`
	Action      string  `json:"action"`
}

type Port struct {
	PortNumber   string `json:"portnumber"`
	RxPackets    int    `json:"rxpackets"`
	TxPackets    int    `json:"txpackets"`
	RxBytes      int    `json:"rxbytes"`
	TxBytes      int    `json:"txbytes"`
	RxDrops      int    `json:"rxdrops"`
	TxDrops      int    `json:"txdrops"`
	RxErrors     string `json:"rxerrors"`
	TxErrors     string `json:"txerrors"`
	RxFrameErr   string `json:"rxframeerr"`
	RxOverruns   string `json:"rxovverruns"`
	RxCrcErrors  string `json:"rxcrcerrors"`
	TxCollisions string `json:"txcollisions"`
}

type Group struct {
	GroupId   string   `json:"groupid"`
	GroupType string   `json:"grouptype"`
	Buckets   []Bucket `json:"buckets"`
	Duration  int      `json:"duration"`
	Bytes     int      `json:"bytes"`
	Packets   int      `json:"packets"`
}

type Bucket struct {
	//    BucketId  string    `json:"bucketid"` // for now I see no real usage of BucketID
	Actions string `json:"actions"`
	Bytes   int    `json:"bytes"`
	Packets int    `json:"packets"`
}
