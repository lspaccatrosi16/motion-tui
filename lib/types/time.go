package types

import (
	"time"
)

type DT struct {
	UnixSecs int
	HasError bool
}

func (d *DT) asTime() time.Time {
	return time.Unix(int64(d.UnixSecs), 0)
}

func (d *DT) ToISO() string {
	return d.asTime().Format(time.RFC3339)
}

func (d *DT) FromISO(str string) *DT {
	t, err := time.Parse(time.RFC3339, str)
	if err != nil {
		d.HasError = true
	} else {
		d.UnixSecs = int(t.Unix())
	}
	return d
}

func (d *DT) FromUnix(i int) *DT {
	d.UnixSecs = i
	return d
}

func (d *DT) ToUnix() int {
	return d.UnixSecs
}

func (d *DT) Now() *DT {
	t := time.Now()
	d.UnixSecs = int(t.Unix())
	return d
}

func (d *DT) InDays(days int) bool {
	t := d.asTime()
	tStartOfDay := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, time.UTC)

	today := time.Now()
	todayStartOfDay := time.Date(today.Year(), today.Month(), today.Day(), 0, 0, 0, 0, time.UTC)

	provTime := todayStartOfDay.Add(time.Duration(days * 24 * int(time.Hour)))

	// fmt.Printf("%s %s\n", todayStartOfDay.Format(time.DateTime), provTime.Format(time.DateTime))

	if provTime.After(tStartOfDay) || provTime.Equal(tStartOfDay) {
		return true
	}

	return false
}

func (d *DT) ShortString() string {
	if d.HasError {
		return "INVALID TIME"
	}
	return d.asTime().Format("Mon 02 Jan")
}

func (d *DT) LongString() string {
	if d.HasError {
		return "INVALID TIME"
	}
	return d.asTime().Format("Mon 02 Jan at 15:04")

}
