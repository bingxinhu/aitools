package dto

import (
	"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ModelCompileApplyGetPageReq struct {
	dto.Pagination `search:"-"`
	Applicant      string `form:"applicant"  search:"type:exact;column:applicant;table:model_compile_apply" comment:"申请人"`
	Status         string `form:"status"  search:"type:exact;column:status;table:model_compile_apply" comment:"状态"`
	Name           string `form:"name"  search:"type:contains;column:name;table:model_compile_apply" comment:"模型名称"`
	Type           string `form:"type"  search:"type:contains;column:type;table:model_compile_apply" comment:"模型类别"`
	Md5            string `form:"md5"  search:"type:exact;column:md5;table:model_compile_apply" comment:"模型md5"`
	ModelCompileApplyOrder
}

type ModelCompileApplyOrder struct {
	Id           string `form:"idOrder"  search:"type:order;column:id;table:model_compile_apply"`
	ApplyTime    string `form:"applyTimeOrder"  search:"type:order;column:apply_time;table:model_compile_apply"`
	Applicant    string `form:"applicantOrder"  search:"type:order;column:applicant;table:model_compile_apply"`
	Status       string `form:"statusOrder"  search:"type:order;column:status;table:model_compile_apply"`
	Name         string `form:"nameOrder"  search:"type:order;column:name;table:model_compile_apply"`
	Type         string `form:"typeOrder"  search:"type:order;column:type;table:model_compile_apply"`
	Onnx_local   string `form:"onnxOrder"  search:"type:order;column:onnx_local;table:model_compile_apply"`
	Md5          string `form:"md5Order"  search:"type:order;column:md5;table:model_compile_apply"`
	NspCnt       string `form:"nspCntOrder"  search:"type:order;column:nsp_cnt;table:model_compile_apply"`
	CompileType  string `form:"compileTypeOrder"  search:"type:order;column:compile_type;table:model_compile_apply"`
	ChannelOrder string `form:"channelOrderOrder"  search:"type:order;column:channel_order;table:model_compile_apply"`
	CalibDataset string `form:"calibDatasetOrder"  search:"type:order;column:calib_dataset;table:model_compile_apply"`
	PostCode     string `form:"postCodeOrder"  search:"type:order;column:post_code;table:model_compile_apply"`
	Platform     string `form:"platformOrder"  search:"type:order;column:platform;table:model_compile_apply"`
	CreatedAt    string `form:"createdAtOrder"  search:"type:order;column:created_at;table:model_compile_apply"`
	UpdatedAt    string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:model_compile_apply"`
	DeletedAt    string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:model_compile_apply"`
	CreateBy     string `form:"createByOrder"  search:"type:order;column:create_by;table:model_compile_apply"`
	UpdateBy     string `form:"updateByOrder"  search:"type:order;column:update_by;table:model_compile_apply"`
}

func (m *ModelCompileApplyGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ModelCompileApplyInsertReq struct {
	Id           int       `json:"-" comment:"序号"` // 序号
	ApplyTime    time.Time `json:"applyTime" comment:"申请时间"`
	Applicant    string    `json:"applicant" comment:"申请人"`
	Status       string    `json:"status" comment:"状态"`
	Name         string    `json:"name" comment:"模型名称"`
	Type         string    `json:"type" comment:"模型类别"`
	Onnx         string    `json:"onnx" comment:"ONNX_md5.onnx"`
	Onnx_local   string    `json:"onnx_local" comment:"ONNX"`
	Md5          string    `json:"md5" comment:"模型md5"`
	NspCnt       string    `json:"nspCnt" comment:"NSP核数"`
	CompileType  string    `json:"compileType" comment:"编译类型"`
	ChannelOrder string    `json:"channelOrder" comment:"通道顺序"`
	CalibDataset string    `json:"calibDataset" comment:"量化数据集"`
	PostCode     string    `json:"postCode" comment:"后处理代码"`
	Platform     string    `json:"platform" comment:"platform"`
	common.ControlBy
}

func (s *ModelCompileApplyInsertReq) Generate(model *models.ModelCompileApply) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ApplyTime = s.ApplyTime
	model.Applicant = s.Applicant
	model.Status = s.Status
	model.Name = s.Name
	model.Type = s.Type
	model.Onnx = s.Onnx
	model.Onnx_local = s.Onnx_local
	model.Platform = s.Platform
	model.Md5 = s.Md5
	model.NspCnt = s.NspCnt
	model.CompileType = s.CompileType
	model.ChannelOrder = s.ChannelOrder
	model.CalibDataset = s.CalibDataset
	model.PostCode = s.PostCode
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *ModelCompileApplyInsertReq) GetId() interface{} {
	return s.Id
}

type ModelCompileApplyUpdateReq struct {
	Id           int       `uri:"id" comment:"序号"` // 序号
	ApplyTime    time.Time `json:"applyTime" comment:"申请时间"`
	Applicant    string    `json:"applicant" comment:"申请人"`
	Status       string    `json:"status" comment:"状态"`
	Name         string    `json:"name" comment:"模型名称"`
	Type         string    `json:"type" comment:"模型类别"`
	Onnx         string    `json:"onnx" comment:"ONNX_md5.onnx"`
	Onnx_local   string    `json:"onnx_local" comment:"ONNX"`
	Md5          string    `json:"md5" comment:"模型md5"`
	NspCnt       string    `json:"nspCnt" comment:"NSP核数"`
	CompileType  string    `json:"compileType" comment:"编译类型"`
	ChannelOrder string    `json:"channelOrder" comment:"通道顺序"`
	CalibDataset string    `json:"calibDataset" comment:"量化数据集"`
	PostCode     string    `json:"postCode" comment:"后处理代码"`
	Platform     string    `json:"platform" comment:"platform"`
	common.ControlBy
}

func (s *ModelCompileApplyUpdateReq) Generate(model *models.ModelCompileApply) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.ApplyTime = s.ApplyTime
	model.Applicant = s.Applicant
	model.Status = s.Status
	model.Name = s.Name
	model.Type = s.Type
	model.Onnx = s.Onnx
	model.Onnx_local = s.Onnx_local
	model.Platform = s.Platform
	model.Md5 = s.Md5
	model.NspCnt = s.NspCnt
	model.CompileType = s.CompileType
	model.ChannelOrder = s.ChannelOrder
	model.CalibDataset = s.CalibDataset
	model.PostCode = s.PostCode
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *ModelCompileApplyUpdateReq) GetId() interface{} {
	return s.Id
}

// ModelCompileApplyGetReq 功能获取请求参数
type ModelCompileApplyGetReq struct {
	Id int `uri:"id"`
}

func (s *ModelCompileApplyGetReq) GetId() interface{} {
	return s.Id
}

// ModelCompileApplyDeleteReq 功能删除请求参数
type ModelCompileApplyDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ModelCompileApplyDeleteReq) GetId() interface{} {
	return s.Ids
}
