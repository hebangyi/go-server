package module

type Config struct {
	Name    string    `yaml:"name"`
	Modules []Modules `yaml:"moudles"`
}

type Modules struct {
	Game string `yaml:"name"`
	User string `yaml:"user"`
}
