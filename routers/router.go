package routers

import (
	"mygramapi/config"
	"mygramapi/controllers"
	"mygramapi/middleware"
	"mygramapi/repository"
	"mygramapi/service"

	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/gin-gonic/gin"
)

func StartServer() *gin.Engine {
	config.DBInit()
	db := config.GetDB()

	userRepository := repository.NewUserRepository(db)
	userService := service.NewUserService(userRepository)
	userController := controllers.NewUserController(userService)

	photoRepository := repository.NewPhotoRepository(db)
	photoService := service.NewPhotoService(photoRepository)
	photoController := controllers.NewPhotoController(photoService)

	commentRepository := repository.NewPhotoRepository(db)
	commentService := service.NewPhotoService(commentRepository)
	commentController := controllers.NewPhotoController(commentService)

	socialMediaRepository := repository.NewSocialMediaRepository(db)
	socialMediaService := service.NewSocialMediaService(socialMediaRepository)
	socialMediaController := controllers.NewSocialMediaController(socialMediaService)

	r := gin.Default()

	public := r.Group("/users")
	{
		public.POST("/register", userController.RegisterUser)
		public.POST("/login", userController.LoginUser)

		public.Use(middleware.Auth())
		public.GET("/", userController.GetUsers)
		public.GET("/:id", userController.GetUserById)
		public.PUT("/:id", userController.UpdateUser)
		public.DELETE("/:id", userController.DeleteUser)
	}

	photoRouter := r.Group("/photos")
	{
		photoRouter.Use(middleware.Auth())
		photoRouter.POST("/", photoController.PostPhoto)
		photoRouter.GET("/", photoController.GetPhotos)
		photoRouter.GET("/:id", photoController.GetPhotoById)
		photoRouter.PUT("/:id", photoController.UpdatePhoto)
		photoRouter.DELETE("/:id", photoController.DeletePhoto)
	}

	commentRouter := r.Group("/comments")
	{
		commentRouter.Use(middleware.Auth())
		commentRouter.POST("/", commentController.PostPhoto)
		commentRouter.GET("/", commentController.GetPhotos)
		commentRouter.GET("/:id", commentController.GetPhotoById)
		commentRouter.PUT("/:id", commentController.UpdatePhoto)
		commentRouter.DELETE("/:id", commentController.DeletePhoto)
	}

	socialMediaRouter := r.Group("/socialmedias")
	{
		socialMediaRouter.Use(middleware.Auth())
		socialMediaRouter.POST("/", socialMediaController.PostSocialMedia)
		socialMediaRouter.GET("/", socialMediaController.GetSocialMedia)
		socialMediaRouter.GET("/:id", socialMediaController.GetSocialMediaById)
		socialMediaRouter.PUT("/:id", socialMediaController.UpdateSocialMedia)
		socialMediaRouter.DELETE("/:id", socialMediaController.DeleteSocialMedia)
	}

	return r
}
