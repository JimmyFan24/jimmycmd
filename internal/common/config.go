package common

type Config struct {
	//Common string
	name string
	path string
}

func NewConfig(name string, path string) *Config {
	conf := &Config{
		name: name,
		path: path,
	}
	return conf
}
