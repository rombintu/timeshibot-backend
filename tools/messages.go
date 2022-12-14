package tools

const (
	DefaultPortServer  string = ":7000"
	DefaultDevDatabase string = "sqlite.db" // file::memory:?cache=shared --- in memory
	// DefaultDevDatabase  string = "file::memory:?cache=shared"
	DefaultHost         string = "localhost"
	DefaultUser         string = "postgres"
	DefaultPass         string = "password"
	DefaultMode         string = "disable"
	DefaultPortDatabase string = "5432"
	DefaultNameDatabase string = "example"
)

const (
	ServerIsStarting string = "Server is starting"
	Create           string = "create"
	Update           string = "update"
	Delete           string = "delete"
	NotFound         string = "not found"
	Get              string = "get"
	All              string = "all"
)

const (
	Mondey    string = "mon"
	Tuesday   string = "tue"
	Wednesday string = "wed"
	Thursday  string = "thu"
	Friday    string = "fri"
	Saturday  string = "sat"
	Sunday    string = "sun"
)
