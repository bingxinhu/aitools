package models

import (
	"time"

	"go-admin/common/models"
)

type ModelCompileApply struct {
	models.Model

	ApplyTime    time.Time `json:"applyTime" gorm:"type:timestamp;comment:申请时间"`
	Applicant    string    `json:"applicant" gorm:"type:varchar(128);comment:申请人"`
	Status       string    `json:"status" gorm:"type:varchar(128);comment:状态"`
	Name         string    `json:"name" gorm:"type:varchar(128);comment:模型名称"`
	Type         string    `json:"type" gorm:"type:varchar(128);comment:模型类别"`
	Onnx         string    `json:"onnx" gorm:"type:varchar(128);comment:ONNX_md5.onnx"`
	Onnx_local   string    `json:"onnx_local" gorm:"type:varchar(128);comment:ONNX"`
	Md5          string    `json:"md5" gorm:"type:varchar(128);comment:模型md5"`
	NspCnt       string    `json:"nspCnt" gorm:"type:varchar(128);comment:NSP核数"`
	CompileType  string    `json:"compileType" gorm:"type:varchar(128);comment:编译类型"`
	ChannelOrder string    `json:"channelOrder" gorm:"type:varchar(45);comment:通道顺序"`
	CalibDataset string    `json:"calibDataset" gorm:"type:varchar(128);comment:量化数据集"`
	PostCode     string    `json:"postCode" gorm:"type:varchar(128);comment:后处理代码"`
	Platform     string    `json:"platform" gorm:"type:varchar(45);comment:platform"`
	models.ModelTime
	models.ControlBy
}

func (ModelCompileApply) TableName() string {
	return "model_compile_apply"
}

func (e *ModelCompileApply) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ModelCompileApply) GetId() interface{} {
	return e.Id
}
