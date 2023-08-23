package service

import (
	"errors"

    "github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type ModelCompileApply struct {
	service.Service
}

// GetPage 获取ModelCompileApply列表
func (e *ModelCompileApply) GetPage(c *dto.ModelCompileApplyGetPageReq, p *actions.DataPermission, list *[]models.ModelCompileApply, count *int64) error {
	var err error
	var data models.ModelCompileApply

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ModelCompileApplyService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取ModelCompileApply对象
func (e *ModelCompileApply) Get(d *dto.ModelCompileApplyGetReq, p *actions.DataPermission, model *models.ModelCompileApply) error {
	var data models.ModelCompileApply

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetModelCompileApply error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建ModelCompileApply对象
func (e *ModelCompileApply) Insert(c *dto.ModelCompileApplyInsertReq) error {
    var err error
    var data models.ModelCompileApply
    c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ModelCompileApplyService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改ModelCompileApply对象
func (e *ModelCompileApply) Update(c *dto.ModelCompileApplyUpdateReq, p *actions.DataPermission) error {
    var err error
    var data = models.ModelCompileApply{}
    e.Orm.Scopes(
            actions.Permission(data.TableName(), p),
        ).First(&data, c.GetId())
    c.Generate(&data)

    db := e.Orm.Save(&data)
    if err = db.Error; err != nil {
        e.Log.Errorf("ModelCompileApplyService Save error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权更新该数据")
    }
    return nil
}

// Remove 删除ModelCompileApply
func (e *ModelCompileApply) Remove(d *dto.ModelCompileApplyDeleteReq, p *actions.DataPermission) error {
	var data models.ModelCompileApply

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
        e.Log.Errorf("Service RemoveModelCompileApply error:%s \r\n", err)
        return err
    }
    if db.RowsAffected == 0 {
        return errors.New("无权删除该数据")
    }
	return nil
}
