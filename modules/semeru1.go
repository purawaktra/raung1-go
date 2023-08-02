package modules

type Semeru1Instance struct {
	Url          string
	AuthUsername string
	AuthPassword string
}

func CreateSemeru1Instance(url string, authUsername string, authPassword string) Semeru1Instance {
	return Semeru1Instance{
		Url:          url,
		AuthUsername: authUsername,
		AuthPassword: authPassword,
	}
}
