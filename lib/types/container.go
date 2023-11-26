package types

type TaskList []*Task

func (t TaskList) Len() int {
	return len(t)
}

func (t TaskList) Swap(i, j int) {
	t[i], t[j] = t[j], t[i]
}

func (t TaskList) Less(i, j int) bool {
	return t[i].Start.UnixSecs < t[j].Start.UnixSecs
}

type AppData struct {
	Runtime *RT
	Tasks   TaskList
}

var appData *AppData

func GetAppData() (*AppData, error) {
	if appData != nil {
		return appData, nil
	}

	rt := new(RT)
	err := rt.Load()

	if err != nil {
		return nil, err
	}

	ad := AppData{Runtime: rt}
	appData = &ad
	return appData, nil
}

func (a *AppData) CanRequest() bool {
	return a.Runtime.CanRequest()
}

func (a *AppData) LogRequest() {
	a.Runtime.LogRequest()
}

func (a *AppData) Save() error {
	return a.Runtime.Save()
}
