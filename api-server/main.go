package main

import (
	docs "ds/api-server/docs"
	"ds/api-server/pkg/api"
	//"net/url"
  //"os"
	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
	//"github.com/maximRnback/gin-oidc"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {

	// @title Data Collector API
	// @version 1.0
	// @description This is a sample data collection server.
	// @termsOfService http://swagger.io/terms/

	// @contact.name API Support
	// @contact.url http://www.swagger.io/support
	// @contact.email support@swagger.io

	// @license.name Apache 2.0
	// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

	// @BasePath /api/v1
	r := gin.Default()
	store := cookie.NewStore([]byte("secret"))
	r.Use(sessions.Sessions("mysession", store))

	// issuer, _ := url.Parse("http://localhost:9998/")
	// clientUrl, _ := url.Parse("http://localhost:8080/")
	// postLogoutUrl, _ := url.Parse("http://localhost:8080/swagger/")
	//
	// initParams := gin_oidc.InitParams{
	// 	Router:       r,
	// 	ClientId:     "web",
	// 	ClientSecret: "secret",
	// 	Issuer:       *issuer,    //add '.well-known/openid-configuration' to see it's a good link
	// 	ClientUrl:    *clientUrl, //your website's url
	// 	Scopes:       []string{"openid"},
	// 	ErrorHandler: func(c *gin.Context) {
	// 		//gin_oidc pushes a new error before any "ErrorHandler" invocation
	// 		message := c.Errors.Last().Error()
	// 		//redirect to ErrorEndpoint with error message
	// 		//redirectToErrorPage(c, "http://example2.domain/error", message)
	// 		c.JSON(401, message)
	// 		//when "ErrorHandler" ends "c.Abort()" is invoked - no further handlers will be invoked
	// 	},
	// 	PostLogoutUrl: *postLogoutUrl,
	// }
	//
	// //protect all endpoint below this line
	// middleware := gin_oidc.Init(initParams)
	// if os.Getenv("OIDC")!="" {
	//   r.Use(middleware)
  // }

	// auth, err := authenticator.New()
	// if err != nil {
	// 	log.Fatalf("Failed to initialize the authenticator: %v", err)
	// }
	//
	// rtr := router.New(auth)
	// log.Print("Server listening on http://localhost:3000/")
	// if err := http.ListenAndServe("0.0.0.0:3000", rtr); err != nil {
	// 	log.Fatalf("There was an error with the http server: %v", err)
	// }

	docs.SwaggerInfo.BasePath = "/api/v1"
	v1 := r.Group("/api/v1")
	{
		auth := v1.Group("/auth")
		{
			auth.POST("", api.Login)
			//auth.POST("", middleware, api.Login)
		}
		data := v1.Group("/data")
		{
			data.POST("", api.PostData)
			//data.POST("", middleware, api.PostData)
		}
		pipeline := v1.Group("/pipeline")
     {
			 pipeline.POST("/:id", api.CreatePipeline)
			 pipeline.GET("/:id", api.GetPipeline)
			 pipeline.DELETE("/:id", api.DeletePipeline)
		 }
	}
	r.GET("/", api.Home)
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))
	r.Run()

}
