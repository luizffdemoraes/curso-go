package configs

var cfg *conf

type conf struct {
	DBDriver      string
	DBHost        string
	DBPort        string
	DBUser        string
	DBPass        string
	DBName        string
	WebServerPort string
	JWTSecret     string
	JwtExperesIn  int
}

func LoadConfig(path string) (*conf, error) {
	// ...


}

