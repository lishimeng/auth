package etc

type Configuration struct {
	Name    string `toml:"name"`
	Version string `toml:"version"`
	Web     web    `toml:"web"`
	User    user   `toml:"user"`
	Token   token  `toml:"token"`
}

type web struct {
	Listen string `toml:"listen"`
}

type user struct {
	Name     string `toml:"name"`
	Password string `toml:"password"`
}

type token struct {
	Issuer string `toml:"issuer"`

	Expire uint16 `toml:"expire"`

	Secret string `toml:"secret"`
}
