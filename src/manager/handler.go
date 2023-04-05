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

type Person struct {
	Cmd    string `json:"cmd"`
	Id     string `json:"id"`
	Tunnel string `json:"tunnel"`
}

type Universal_out struct {
	Msg string `json:"msg"`
}

type New_connection_in struct {
	Cmd string `json:"cmd"`
	Id  string `json:"id"`
}
