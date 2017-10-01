package main

import "fmt"

type duration int64

const (
	nanosecond  duration = 1
	microsecond          = 1000 * nanosecond
	millisecond          = 1000 * microsecond
	second               = 1000 * millisecond
	minute               = 60 * second
	hour                 = 60 * minute
)

func (d *duration) setHours(h float64) {
	*d = duration(h) * hour
}

func (d duration) hours() float64 {
	hour := d / hour
	nsec := d % hour
	return float64(hour) + float64(nsec)*(1e-9/60/60)
}

func main() {

	var dur duration

	dur.setHours(5)
	fmt.Println("hours: ", dur.hours())
	fmt.Println("hours: ", dur.hours())
}
