package manager

type Request struct {
	Cmd     string `json:"cmd"`
	Id      string `json:"id"`
	Ip      string `json:"ip"`
	Port    string `json:"port"`
	Rating  string `json:"rating"`
	Hs      bool   `json:"false"`
	Payload string `json:"payload"`
}

var cmdHandlers = map[string]func(int) string{
	"1": nil,
}
