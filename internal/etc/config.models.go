package etc

type Configuration struct {
	Name    string
	Version string
	Web     web
	Token   token
	Redis   redis
	Db      db
	Mail    mail
}

type web struct {
	Listen string
}

type token struct {
	Issuer string
	Expire uint16
	Secret string
}

type redis struct {
	Enable   bool
	Addr     string
	Password string
	Db       int
}

type db struct {
	User     string
	Password string
	Host     string
	Port     int
	Database string
	Ssl      string
}

type mail struct {
	Host   string
	Sender string
	Debug  bool
}
