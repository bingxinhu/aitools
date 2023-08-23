package router

import (
	"github.com/gin-gonic/gin"
	jwt "github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth"

	"go-admin/app/admin/apis"
	"go-admin/common/actions"
	"go-admin/common/middleware"
)

func init() {
	routerCheckRole = append(routerCheckRole, registerModelCompileApplyRouter)
}

// registerModelCompileApplyRouter
func registerModelCompileApplyRouter(v1 *gin.RouterGroup, authMiddleware *jwt.GinJWTMiddleware) {
	api := apis.ModelCompileApply{}
	r := v1.Group("/model-compile-apply").Use(authMiddleware.MiddlewareFunc()).Use(middleware.AuthCheckRole())
	{
		r.GET("", actions.PermissionAction(), api.GetPage)
		r.GET("/:id", actions.PermissionAction(), api.Get)
		r.POST("", api.Insert)
		r.PUT("/:id", actions.PermissionAction(), api.Update)
		r.DELETE("", api.Delete)
		r.GET("/userrole", actions.PermissionAction(), api.GetUserRole)
	}
	/// QUERY:: 		curl http://127.0.0.1:8000/api/v1/model-compile-apply/query_noauth?status=待编译
	/// download onnx:: wget http://10.40.29.152:8000/static/myuploadfiles/e2b186d17a4ad4fed6a5ef62f36ab340.png
	/// update status:: curl -X PUT http://127.0.0.1:8000/api/v1/model-compile-apply/update_noauth/16 --data '{"id":16,"applyTime":"2022-11-14T15:28:20+08:00","applicant":"刘方盛","status":"编译中","name":"啊啊啊","type":"啊啊","onnx":"e2b186d17a4ad4fed6a5ef62f36ab340.png","onnx_local":"住宿2.png","md5":"啊","nspCnt":"2","compileType":"INT8","channelOrder":"NHWC","calibDataset":"","postCode":"","createdAt":"2022-11-14T15:28:42+08:00","updatedAt":"2022-11-14T15:36:05+08:00","createBy":1,"updateBy":1}' -H "Content-Type:application/json"
	r_noAuth := v1.Group("/model-compile-apply")
	{
		r_noAuth.GET("/query_noauth", api.GetPage)
		r_noAuth.PUT("/update_noauth/:id", api.Update)
	}
}
