package structs

type (
	// Config setting
	Config struct {
		SQLConf Mysql `yaml:"mysql"`
	}

	// Mysql of config yml
	Mysql struct {
		User     string `yaml:"user"`
		Password string `yaml:"password"`
		Name     string `yaml:"name"`
	}
)
