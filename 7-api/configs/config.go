package configs

import (
	"github.com/go-chi/jwtauth"
	"github.com/spf13/viper"
)

type conf struct {
	DBDriver      string `mapstructure:"DB_DRIVER"`
	DBHost        string `mapstructure:"DB_HOST"`
	DBPort        string `mapstructure:"DB_PORT"`
	DBName        string `mapstructure:"DB_NAME"`
	DBUser        string `mapstructure:"DB_USER"`
	DBPassword    string `mapstructure:"DB_PASSWORD"`
	WebServerPort string `mapstructure:"WEB_SERVER_PORT"`
	JWTSecret     string `mapstructure:"JWT_SECRET"`
	JWTExpiresIn  int    `mapstructure:"JWT_EXPIRESIN"`
	TokenAuth     *jwtauth.JWTAuth
}

func LoadConfig(path string) (*conf, error) {
	viper.SetConfigName("app_config")
	// define o tipo de origem que as configurações serão efetuadas
	viper.SetConfigType("env")
	// define o caminho que o arquivo de configurações estara
	viper.AddConfigPath(path)
	// define o nome do arquivo de configurações
	viper.SetConfigFile(".env")
	// define que se houver variaveis de ambiente elas sobrescreverão o valor do arquivo de configuração
	viper.AutomaticEnv()
	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
	var cfg *conf
	err = viper.Unmarshal(&cfg)
	if err != nil {
		panic(err)
	}
	// Essa instancia nos possibilita assinar e gerar tokens JWT
	cfg.TokenAuth = jwtauth.New("HS256", []byte(cfg.JWTSecret), nil)
	return cfg, nil
}
