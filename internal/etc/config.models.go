package etc

type Configuration struct {
	Name       string `toml:"name"`
	Version    string `toml:"version"`
	Web        web `toml:"web"`
	Token        token `toml:"token"`
}

type web struct {
	Listen string `toml:"listen"`
}

type token struct {
	Issuer string `toml:"issuer"`

	Expire uint16 `toml:"expire"`

	Secret string `toml:"secret"`
}
