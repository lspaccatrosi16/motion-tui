package types

import "time"

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

func (d *DT) String() string {
	if d.HasError {
		return "INVALID TIME"
	}
	return d.asTime().Format(time.DateTime)
}
