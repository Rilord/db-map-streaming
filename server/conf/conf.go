package conf

var (
	LenStackBuf = 4096

	LogLevel string
	LogPath  string
	LogFlag  int

	ConsolePort   int
	ConsolePrompt string = "Server# "
	ProfilePath   string

	ListenAddr      string
	ConnAddrs       []string
	PendingWriteNum uint
)
