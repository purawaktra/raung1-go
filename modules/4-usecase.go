package modules

type Raung1Usecase struct {
	repo Raung1Repo
}

func CreateRaung1Usecase(repo Raung1Repo) Raung1Usecase {
	return Raung1Usecase{
		repo: repo,
	}
}
