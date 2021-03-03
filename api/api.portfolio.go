package api

import "github.com/gin-gonic/gin"

type portfolioApi struct {}

func (a *portfolioApi) Routes(r gin.IRoutes) {
	r.POST("", a.create)
	r.GET("", a.getAll)
	r.GET("/:id", a.get)
	r.POST("/:id", a.update)
	r.POST("/:id/requestOptimization", a.requestOptimization)
}

func (a *portfolioApi) get(ctx *gin.Context) {}

func (a *portfolioApi) create(ctx *gin.Context) {}

func (a *portfolioApi) update(ctx *gin.Context) {}

func (a *portfolioApi) getAllByClient(ctx *gin.Context) {}

func (a *portfolioApi) getAll(ctx *gin.Context) {}

func (a *portfolioApi) requestOptimization(ctx *gin.Context) {}




