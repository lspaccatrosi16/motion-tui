package types

import (
	"io"
	"os"
	"path/filepath"
	"time"

	"github.com/lspaccatrosi16/go-cli-tools/config"
	"github.com/lspaccatrosi16/go-cli-tools/gbin"
)

const REQ_LIMIT = 6

type RT struct {
	PrevRequests []*DT
}

func (r *RT) CanRequest() bool {
	r.CleanRequestList()
	return r.NumRequests() < REQ_LIMIT
}

func (r *RT) LogRequest() {
	t := new(DT).Now()
	r.PrevRequests = append(r.PrevRequests, t)
}

func (r *RT) CleanRequestList() {
	newList := []*DT{}

	timeThreshold := time.Now().Add(-time.Minute).Unix()

	for _, req := range r.PrevRequests {
		if req.ToUnix() >= int(timeThreshold) {
			newList = append(newList, req)
		}
	}

	r.PrevRequests = newList
}

func (r *RT) NumRequests() int {
	return len(r.PrevRequests)
}

func (r *RT) Load() error {
	path, err := cpath()
	if err != nil {
		return err
	}

	f, err := os.Open(path)

	if err != nil {
		if os.IsNotExist(err) {
			return nil
		} else {
			return err
		}
	}

	defer f.Close()

	decoder := gbin.NewDecoder[RT]()
	out, err := decoder.DecodeStream(f)
	if err != nil {
		return err
	}

	*r = *out
	return nil
}

func (r *RT) Save() error {
	path, err := cpath()
	if err != nil {
		return err
	}

	f, err := os.Create(path)
	if err != nil {
		return err
	}
	defer f.Close()

	encoder := gbin.NewEncoder[RT]()

	src, err := encoder.EncodeStream(r)
	if err != nil {
		return err
	}

	io.Copy(f, src)
	return nil
}

func cpath() (string, error) {
	base, err := config.GetConfigPath("motionTui")
	if err != nil {
		return "", err
	}

	return filepath.Join(base, "config.bin"), nil
}
