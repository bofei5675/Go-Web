package middleware
import (
	"net/http"
	"time"
	"github.com/gin-gonic/gin"
)
/*
NoCache is a middleware function that appends headers
to prevent the client from caching the HTTP Response.
 */
func NoCache(c *gin.Context) {
	c.Header("Cache-Control","no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires","Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

/*
Options is middleware function that appends headers
for options requests and aborts then exists the middleware
chain and ends the request
The HTTP OPTIONS method is used to describe the communication options for the target resource.
The client can specify a URL for the OPTIONS method, or an asterisk (*) to refer to the entire server.
 */
func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header("Access-Control-Allow-Origin","*")
		c.Header("Access-Control-Allow-Methods","GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Access-Control-Allow-Headers", "authorization, origin, content-type, accept")
		c.Header("Allow","HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header("Content-Type","application/json")
		c.AbortWithStatus(200)
	}
}

/*
Secure is a middleware function that appends security and
resource access headers
 */

func Secure(c *gin.Context)  {
	c.Header("Access-Control-Allow-Origin","*")
	c.Header("X-Frame-Options","DENY")
	c.Header("X-Content-Type-Options","nosniff")
	c.Header("X-XSS-Protection","1; mode=block")
	// Also consider adding Content-Security-Policy headers
	// c.Header("Content-Security-Policy", "script-src 'self' https://cdnjs.cloudflare.com")

}

