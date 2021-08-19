package data

import "time"

type Hive struct {
	id int8
}

type Bees struct {
	id int64
	name string
	created_at time.Time
	age int
	species string
	position string
	characteristics []string
	queen bool
	hive Hive
	version int
}


type Caste int8

const (
	Worker Caste = iota
	Drone
	Queen
)

func (c Caste) String() string {
	switch c {
	case Worker:
		return "Worker"
	case Drone:
		return "Drone"
	case Queen:
		return "Queen"
	default:
		return "Little One"
	}
}





