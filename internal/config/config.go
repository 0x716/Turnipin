package config

var GlobalConfig *AppConfig

type AppConfig struct {
	App struct {
		Host string
		Port string
		Name string
	}

	URL struct {
		Root string
	}

	Database struct {
		Title    string
		Name     string
		User     string
		Password string
		Engine   string
	}

	Development struct {
		Mode string
	}
}
