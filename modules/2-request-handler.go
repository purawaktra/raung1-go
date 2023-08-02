package modules

type Raung1RequestHandler struct {
	ctrl Raung1Controller
}

func CreateRaung1RequestHandler(ctrl Raung1Controller) Raung1RequestHandler {
	return Raung1RequestHandler{
		ctrl: ctrl,
	}
}
