package internal

import "strconv"

type Interval struct {
	Start int
	End   int
}

func (i Interval) toString() string {
	if i.End == -1 {
		return ">" + strconv.Itoa(i.Start)
	} else {
		return strconv.Itoa(i.Start) + "-" + strconv.Itoa(i.End)
	}
}
