// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package dao

import (
	"context"

	"gorm.io/gorm"
	"gorm.io/gorm/clause"
	"gorm.io/gorm/schema"

	"gorm.io/gen"
	"gorm.io/gen/field"

	"gorm.io/plugin/dbresolver"

	"github.com/MQEnergy/mqshop/internal/app/model"
)

func newCAttachment(db *gorm.DB, opts ...gen.DOOption) cAttachment {
	_cAttachment := cAttachment{}

	_cAttachment.cAttachmentDo.UseDB(db, opts...)
	_cAttachment.cAttachmentDo.UseModel(&model.CAttachment{})

	tableName := _cAttachment.cAttachmentDo.TableName()
	_cAttachment.ALL = field.NewAsterisk(tableName)
	_cAttachment.ID = field.NewInt64(tableName, "id")
	_cAttachment.UserID = field.NewInt64(tableName, "user_id")
	_cAttachment.AttachName = field.NewString(tableName, "attach_name")
	_cAttachment.AttachOriginName = field.NewString(tableName, "attach_origin_name")
	_cAttachment.AttachURL = field.NewString(tableName, "attach_url")
	_cAttachment.AttachType = field.NewInt8(tableName, "attach_type")
	_cAttachment.AttachMimetype = field.NewString(tableName, "attach_mimetype")
	_cAttachment.AttachExtension = field.NewString(tableName, "attach_extension")
	_cAttachment.AttachSize = field.NewString(tableName, "attach_size")
	_cAttachment.Status = field.NewInt8(tableName, "status")
	_cAttachment.CreatedAt = field.NewInt64(tableName, "created_at")
	_cAttachment.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_cAttachment.fillFieldMap()

	return _cAttachment
}

// cAttachment 附件表
type cAttachment struct {
	cAttachmentDo

	ALL              field.Asterisk
	ID               field.Int64
	UserID           field.Int64  // 附件上传的用户id
	AttachName       field.String // 附件新名称
	AttachOriginName field.String // 附件原名称
	AttachURL        field.String // 附件地址
	AttachType       field.Int8   // 附件类型 1：图片 2：视频 3：文件
	AttachMimetype   field.String // 附件mime类型
	AttachExtension  field.String // 附件后缀名
	AttachSize       field.String // 附件大小
	Status           field.Int8   // 状态 1：正常 0：删除
	CreatedAt        field.Int64
	UpdatedAt        field.Int64

	fieldMap map[string]field.Expr
}

func (c cAttachment) Table(newTableName string) *cAttachment {
	c.cAttachmentDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cAttachment) As(alias string) *cAttachment {
	c.cAttachmentDo.DO = *(c.cAttachmentDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cAttachment) updateTableName(table string) *cAttachment {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.UserID = field.NewInt64(table, "user_id")
	c.AttachName = field.NewString(table, "attach_name")
	c.AttachOriginName = field.NewString(table, "attach_origin_name")
	c.AttachURL = field.NewString(table, "attach_url")
	c.AttachType = field.NewInt8(table, "attach_type")
	c.AttachMimetype = field.NewString(table, "attach_mimetype")
	c.AttachExtension = field.NewString(table, "attach_extension")
	c.AttachSize = field.NewString(table, "attach_size")
	c.Status = field.NewInt8(table, "status")
	c.CreatedAt = field.NewInt64(table, "created_at")
	c.UpdatedAt = field.NewInt64(table, "updated_at")

	c.fillFieldMap()

	return c
}

func (c *cAttachment) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cAttachment) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 12)
	c.fieldMap["id"] = c.ID
	c.fieldMap["user_id"] = c.UserID
	c.fieldMap["attach_name"] = c.AttachName
	c.fieldMap["attach_origin_name"] = c.AttachOriginName
	c.fieldMap["attach_url"] = c.AttachURL
	c.fieldMap["attach_type"] = c.AttachType
	c.fieldMap["attach_mimetype"] = c.AttachMimetype
	c.fieldMap["attach_extension"] = c.AttachExtension
	c.fieldMap["attach_size"] = c.AttachSize
	c.fieldMap["status"] = c.Status
	c.fieldMap["created_at"] = c.CreatedAt
	c.fieldMap["updated_at"] = c.UpdatedAt
}

func (c cAttachment) clone(db *gorm.DB) cAttachment {
	c.cAttachmentDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cAttachment) replaceDB(db *gorm.DB) cAttachment {
	c.cAttachmentDo.ReplaceDB(db)
	return c
}

type cAttachmentDo struct{ gen.DO }

type ICAttachmentDo interface {
	gen.SubQuery
	Debug() ICAttachmentDo
	WithContext(ctx context.Context) ICAttachmentDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICAttachmentDo
	WriteDB() ICAttachmentDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICAttachmentDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICAttachmentDo
	Not(conds ...gen.Condition) ICAttachmentDo
	Or(conds ...gen.Condition) ICAttachmentDo
	Select(conds ...field.Expr) ICAttachmentDo
	Where(conds ...gen.Condition) ICAttachmentDo
	Order(conds ...field.Expr) ICAttachmentDo
	Distinct(cols ...field.Expr) ICAttachmentDo
	Omit(cols ...field.Expr) ICAttachmentDo
	Join(table schema.Tabler, on ...field.Expr) ICAttachmentDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICAttachmentDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICAttachmentDo
	Group(cols ...field.Expr) ICAttachmentDo
	Having(conds ...gen.Condition) ICAttachmentDo
	Limit(limit int) ICAttachmentDo
	Offset(offset int) ICAttachmentDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICAttachmentDo
	Unscoped() ICAttachmentDo
	Create(values ...*model.CAttachment) error
	CreateInBatches(values []*model.CAttachment, batchSize int) error
	Save(values ...*model.CAttachment) error
	First() (*model.CAttachment, error)
	Take() (*model.CAttachment, error)
	Last() (*model.CAttachment, error)
	Find() ([]*model.CAttachment, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CAttachment, err error)
	FindInBatches(result *[]*model.CAttachment, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CAttachment) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICAttachmentDo
	Assign(attrs ...field.AssignExpr) ICAttachmentDo
	Joins(fields ...field.RelationField) ICAttachmentDo
	Preload(fields ...field.RelationField) ICAttachmentDo
	FirstOrInit() (*model.CAttachment, error)
	FirstOrCreate() (*model.CAttachment, error)
	FindByPage(offset int, limit int) (result []*model.CAttachment, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICAttachmentDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cAttachmentDo) Debug() ICAttachmentDo {
	return c.withDO(c.DO.Debug())
}

func (c cAttachmentDo) WithContext(ctx context.Context) ICAttachmentDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cAttachmentDo) ReadDB() ICAttachmentDo {
	return c.Clauses(dbresolver.Read)
}

func (c cAttachmentDo) WriteDB() ICAttachmentDo {
	return c.Clauses(dbresolver.Write)
}

func (c cAttachmentDo) Session(config *gorm.Session) ICAttachmentDo {
	return c.withDO(c.DO.Session(config))
}

func (c cAttachmentDo) Clauses(conds ...clause.Expression) ICAttachmentDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cAttachmentDo) Returning(value interface{}, columns ...string) ICAttachmentDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cAttachmentDo) Not(conds ...gen.Condition) ICAttachmentDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cAttachmentDo) Or(conds ...gen.Condition) ICAttachmentDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cAttachmentDo) Select(conds ...field.Expr) ICAttachmentDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cAttachmentDo) Where(conds ...gen.Condition) ICAttachmentDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cAttachmentDo) Order(conds ...field.Expr) ICAttachmentDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cAttachmentDo) Distinct(cols ...field.Expr) ICAttachmentDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cAttachmentDo) Omit(cols ...field.Expr) ICAttachmentDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cAttachmentDo) Join(table schema.Tabler, on ...field.Expr) ICAttachmentDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cAttachmentDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICAttachmentDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cAttachmentDo) RightJoin(table schema.Tabler, on ...field.Expr) ICAttachmentDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cAttachmentDo) Group(cols ...field.Expr) ICAttachmentDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cAttachmentDo) Having(conds ...gen.Condition) ICAttachmentDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cAttachmentDo) Limit(limit int) ICAttachmentDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cAttachmentDo) Offset(offset int) ICAttachmentDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cAttachmentDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICAttachmentDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cAttachmentDo) Unscoped() ICAttachmentDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cAttachmentDo) Create(values ...*model.CAttachment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cAttachmentDo) CreateInBatches(values []*model.CAttachment, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cAttachmentDo) Save(values ...*model.CAttachment) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cAttachmentDo) First() (*model.CAttachment, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CAttachment), nil
	}
}

func (c cAttachmentDo) Take() (*model.CAttachment, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CAttachment), nil
	}
}

func (c cAttachmentDo) Last() (*model.CAttachment, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CAttachment), nil
	}
}

func (c cAttachmentDo) Find() ([]*model.CAttachment, error) {
	result, err := c.DO.Find()
	return result.([]*model.CAttachment), err
}

func (c cAttachmentDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CAttachment, err error) {
	buf := make([]*model.CAttachment, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cAttachmentDo) FindInBatches(result *[]*model.CAttachment, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cAttachmentDo) Attrs(attrs ...field.AssignExpr) ICAttachmentDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cAttachmentDo) Assign(attrs ...field.AssignExpr) ICAttachmentDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cAttachmentDo) Joins(fields ...field.RelationField) ICAttachmentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cAttachmentDo) Preload(fields ...field.RelationField) ICAttachmentDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cAttachmentDo) FirstOrInit() (*model.CAttachment, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CAttachment), nil
	}
}

func (c cAttachmentDo) FirstOrCreate() (*model.CAttachment, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CAttachment), nil
	}
}

func (c cAttachmentDo) FindByPage(offset int, limit int) (result []*model.CAttachment, count int64, err error) {
	result, err = c.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = c.Offset(-1).Limit(-1).Count()
	return
}

func (c cAttachmentDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cAttachmentDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cAttachmentDo) Delete(models ...*model.CAttachment) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cAttachmentDo) withDO(do gen.Dao) *cAttachmentDo {
	c.DO = *do.(*gen.DO)
	return c
}
