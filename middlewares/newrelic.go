package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/newrelic/go-agent/v3/newrelic"
)

func NewRelicMiddleware(app *newrelic.Application) gin.HandlerFunc {
	return func(c *gin.Context) {
		txn := app.StartTransaction(c.Request.URL.Path)
		defer txn.End()

		c.Request = newrelic.RequestWithTransactionContext(c.Request, txn)
		c.Next()
		txn.SetWebResponse(c.Writer)
	}
}
