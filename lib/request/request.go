package request

import (
	"bytes"
	"fmt"
	"io"
	"net/http"

	"github.com/lspaccatrosi16/go-cli-tools/credential"
	"github.com/lspaccatrosi16/motion-tui/lib/types"
)

func MakeGetRequest(url string) ([]byte, error) {
	appData, err := types.GetAppData()
	if err != nil {
		return nil, err
	}

	cred, err := credential.GetUserAuth(types.APP_NAME)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("GET", url, nil)

	if err != nil {
		return nil, err
	}

	req.Header.Add("X-API-Key", cred.Secret)

	if appData.CanRequest() {
		appData.LogRequest()
		appData.Save()
		resp, err := http.DefaultClient.Do(req)

		if err != nil {
			return nil, err
		}

		defer resp.Body.Close()

		buf := bytes.NewBuffer(nil)
		io.Copy(buf, resp.Body)

		return buf.Bytes(), err
	} else {
		return nil, fmt.Errorf("rate limit exceded")
	}
}
