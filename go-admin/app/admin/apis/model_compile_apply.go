package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type ModelCompileApply struct {
	api.Api
}

// GetPage 获取模型编译申请列表
// @Summary 获取模型编译申请列表
// @Description 获取模型编译申请列表
// @Tags 模型编译申请
// @Param applicant query string false "申请人"
// @Param status query string false "状态"
// @Param name query string false "模型名称"
// @Param type query string false "模型类别"
// @Param md5 query string false "模型md5"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ModelCompileApply}} "{"code": 200, "data": [...]}"
// @Router /api/v1/model-compile-apply [get]
// @Security Bearer
func (e ModelCompileApply) GetPage(c *gin.Context) {
	req := dto.ModelCompileApplyGetPageReq{}
	s := service.ModelCompileApply{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.ModelCompileApply, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取模型编译申请失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

func (e ModelCompileApply) GetUserRole(c *gin.Context) {
	err := e.MakeContext(c).Errors
	if err != nil {
		e.Logger.Error(err)
		return
	}
	e.OK(user.GetRoleName(c), "查询成功")
}

// Get 获取模型编译申请
// @Summary 获取模型编译申请
// @Description 获取模型编译申请
// @Tags 模型编译申请
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.ModelCompileApply} "{"code": 200, "data": [...]}"
// @Router /api/v1/model-compile-apply/{id} [get]
// @Security Bearer
func (e ModelCompileApply) Get(c *gin.Context) {
	req := dto.ModelCompileApplyGetReq{}
	s := service.ModelCompileApply{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.ModelCompileApply

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取模型编译申请失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建模型编译申请
// @Summary 创建模型编译申请
// @Description 创建模型编译申请
// @Tags 模型编译申请
// @Accept application/json
// @Product application/json
// @Param data body dto.ModelCompileApplyInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/model-compile-apply [post]
// @Security Bearer
func (e ModelCompileApply) Insert(c *gin.Context) {
	req := dto.ModelCompileApplyInsertReq{}
	s := service.ModelCompileApply{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建模型编译申请失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改模型编译申请
// @Summary 修改模型编译申请
// @Description 修改模型编译申请
// @Tags 模型编译申请
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ModelCompileApplyUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/model-compile-apply/{id} [put]
// @Security Bearer
func (e ModelCompileApply) Update(c *gin.Context) {
	req := dto.ModelCompileApplyUpdateReq{}
	s := service.ModelCompileApply{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改模型编译申请失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除模型编译申请
// @Summary 删除模型编译申请
// @Description 删除模型编译申请
// @Tags 模型编译申请
// @Param data body dto.ModelCompileApplyDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/model-compile-apply [delete]
// @Security Bearer
func (e ModelCompileApply) Delete(c *gin.Context) {
	s := service.ModelCompileApply{}
	req := dto.ModelCompileApplyDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除模型编译申请失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
