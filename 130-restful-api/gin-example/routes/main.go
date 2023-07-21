package routes

import (
	"html/template"
	"log"

	"github.com/gin-gonic/gin"
)

var router = gin.Default()
var html = template.Must(template.New("http").Parse(`
<html>
<head>
  <title>Http Test</title>
  <script src="/assets/app.js"></script>
</head>
<body>
  <h1 style="color:red;">Welcome, Ginner!</h1>
</body>
</html>
`))

func Run() {
	getRoutes()
	router.Static("/assets", "./assets")
	router.SetHTMLTemplate(html)

	router.GET("/push", func(c *gin.Context) {
		if pusher := c.Writer.Pusher(); pusher != nil {
			if err := pusher.Push("/assets/app.js", nil); err != nil {
				log.Printf("Failed to push: %v", err)
			}
		}

		c.HTML(200, "http", gin.H{
			"status": "success",
		})

	})

	router.Run(":5009")
}

func getRoutes() {
	v1 := router.Group("/v1")
	addUserRoutes(v1)
	addPingRoutes(v1)

	v2 := router.Group("/v2")
	addPingRoutes(v2)
}
