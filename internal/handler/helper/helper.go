package helper

type Response struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
}

type Error struct {
	Error string `json:"error"`
}

type rWallet interface {
	GetByToken(token string) (string, string, error)
}

type Handler struct {
	repoWallet rWallet
}

func New(repoWallet rWallet) (*Handler, error) {
	return &Handler{
		repoWallet: repoWallet,
	}, nil
}
