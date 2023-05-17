package routers

import (
	"campus_pharmacy_share/controllers"
	"campus_pharmacy_share/middlewares"
	"github.com/gin-gonic/gin"
)

func InitRouter() *gin.Engine {
	router := gin.Default()

	// 用户相关路由
	user := router.Group("/user")
	{
		user.POST("/register", controllers.Register)
		user.POST("/login", controllers.Login)
	}

	// 需要认证的路由
	auth := router.Group("/")
	auth.Use(middlewares.JWTAuth())
	{
		// 区域相关路由
		auth.GET("/regions", controllers.GetRegions)
		auth.POST("/region", controllers.CreateRegion)

		// 药品相关路由
		auth.POST("/medicine", controllers.CreateMedicine)
		auth.GET("/medicines/:region_id", controllers.GetMedicinesByRegion)
		auth.PUT("/medicine", controllers.UpdateMedicine)
		auth.DELETE("/medicine/:id", controllers.DeleteMedicine)
	}

	return router
}
func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.Use(middlewares.CORSMiddleware())

	v1 := router.Group("/api/v1")
	{

		// 使用JWTAuth中间件
		auth := v1.Group("/")
		v1.POST("/register", controllers.Register)
		v1.POST("/", controllers.Login)
		auth.Use(middlewares.JWTAuth())
		{ // 用户相关路由
			auth.GET("/user/:id", controllers.GetUser)

			auth.PUT("/user", controllers.UpdateUser)

			// 生活园区(区域)相关路由
			//auth.GET("/regions", controllers.GetRegions)
			//auth.POST("/region", controllers.CreateRegion)
			// 生活园区相关路由
			auth.POST("/living_area", controllers.CreateLivingArea)
			auth.GET("/living_area/:id", controllers.GetLivingAreaByID)
			auth.PUT("/living_area", controllers.UpdateLivingArea)
			auth.DELETE("/living_area/:id", controllers.DeleteLivingArea)
			// 药品相关路由
			auth.POST("/medicine", controllers.CreateMedicine)
			auth.GET("/regions/:region_id/medicines", controllers.GetMedicinesByRegion)
			auth.PUT("/medicine", controllers.UpdateMedicine)
			auth.DELETE("/medicine/:id", controllers.DeleteMedicine)
			auth.GET("/search/medicines", controllers.SearchMedicines) // 添加搜索药品路由
			auth.GET("/medicines/:medicine_id/images", controllers.GetMedicineImages)

			// 帖子相关路由
			auth.POST("/post", controllers.CreatePost)
			auth.GET("/post", controllers.GetPostsByUserID)
			auth.PUT("/post/:id", controllers.UpdatePost)
			auth.DELETE("/post/:id", controllers.DeletePost)

			// 药品图片相关路由
			auth.POST("/medicine_image", controllers.CreateMedicineImage)
			auth.GET("/medicine_images/:medicine_id", controllers.GetMedicineImagesByMedicineID)
			auth.DELETE("/medicine_image/:id", controllers.DeleteMedicineImage)
		}
	}
	superAdmin := auth.Group("/")
	superAdmin.Use(middlewares.RequireRole("超级管理员"))
	{
		superAdmin.GET("/search/regions", controllers.SearchRegions)
		superAdmin.GET("/regions/:region_id/users", controllers.GetUsersByRegion)
		superAdmin.GET("/search/users", controllers.SearchUsers)
		superAdmin.PUT("/user/:id/promote", controllers.PromoteUserToManager)
		superAdmin.GET("/all_medicines", controllers.GetAllMedicines)
		superAdmin.GET("/search/medicines", controllers.SearchMedicines)
		superAdmin.POST("/announcement", controllers.CreateAnnouncement)
	}

	return router
}
