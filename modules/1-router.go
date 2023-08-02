package modules

import "github.com/gin-gonic/gin"

type Raung1Router struct {
	engine *gin.Engine
	rh     Raung1RequestHandler
}

func CreateRaung1Router(engine *gin.Engine, rh Raung1RequestHandler) Raung1Router {
	return Raung1Router{
		engine: engine,
		rh:     rh,
	}
}
