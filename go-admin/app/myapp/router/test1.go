package router

import (
	"go-admin/app/myapp/apis"

	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerTest1Router)
}

// 需认证的路由代码
func registerTest1Router(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.Test1{}
	r := v1.Group("")
	{
		r.GET("/test1", api.GetTest1)
		r.POST("/upload_single", api.UploadFile)
		r.POST("/upload_single_definename", api.UploadFile_definename)
	}
}
