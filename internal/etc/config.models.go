package etc

type Configuration struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
	Web     web    `toml:"web"`
	Token   token  `toml:"token"`
	Redis   redis  `toml:"redis"`
}

type web struct {
	Listen string `toml:"listen"`
}

type token struct {
	Issuer string `toml:"issuer"`
	Expire uint16 `toml:"expire"`
	Secret string `toml:"secret"`
}

type redis struct {
	Enable bool `toml:"enable"`
	Addr string `toml:"addr"`
	Password string `toml:"password"`
	Db int `toml:"db"`
}
