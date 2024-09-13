package common

type Config struct {
	Url          string
	Mode         string
	Check        bool
	Domain       string
	Username     string
	Password     string
	UserFile     string
	PassFile     string
	UserPassFile string
	UserAsPass   bool
	OutputFile   string
	PthMode      bool
	Proxy        string
	Threads      int
	NoColor      bool
	Verbose      bool
	Delay        int
}

type TaskInfo struct {
	Target  string
	Domain  string
	Mode    string
	Task    chan []string
	Result  *Result
	PthMode bool
	Delay   int
	Proxy   string
}

type BruteWorker func(info *TaskInfo)

var ExchangeEndpoints = map[string]string{
	"autodiscover": "/autodiscover",
	"ews":          "/ews",
	"mapi":         "/mapi",
	"activesync":   "/Microsoft-Server-ActiveSync",
	"oab":          "/oab/global.asax",
	"rpc":          "/rpc",
	"owa":          "/owa/auth.owa",
	"powershell":   "/powershell",
	"ecp":          "/owa/auth.owa",
}

var (
	UserPassDict map[string]string
	UserDict     []string
	PassDict     []string
	Worker       BruteWorker
	Log          Logger
)
