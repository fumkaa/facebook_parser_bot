package configs

import (
	"log"
	"sync"

	"github.com/ilyakaznacheev/cleanenv"
)

type Configuration struct {
	Token_bot  string `env:"TOKEN_BOT"`
	UserDB     string `env:"USER_DB"`
	PasswordDB string `env:"PASSWORD_DB"`
	PortDB     string `env:"PORT_DB"`
	NameDB     string `env:"DATABASE_NAME"`
}

var (
	instance *Configuration
	once     sync.Once
)

func NewConfiguration() *Configuration {
	once.Do(func() {
		log.Print("read app config")
		instance = &Configuration{}
		parseEnv(instance)
		// if err := cleanenv.ReadConfig("./settings/config.yaml", instance); err != nil {
		// 	help, _ := cleanenv.GetDescription(instance, nil)
		// 	log.Printf("ParseConfig erorr: %v", err)
		// 	log.Fatal(help)
		// }
	})
	log.Print(*instance)
	return instance
}

func parseEnv(config *Configuration) {
	if err := cleanenv.ReadEnv(config); err != nil {
		help, _ := cleanenv.GetDescription(instance, nil)
		log.Printf("parseEnv erorr: %v", err)
		log.Fatal(help)
	}

}
