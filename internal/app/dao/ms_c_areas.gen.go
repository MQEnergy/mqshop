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

func newCAreas(db *gorm.DB, opts ...gen.DOOption) cAreas {
	_cAreas := cAreas{}

	_cAreas.cAreasDo.UseDB(db, opts...)
	_cAreas.cAreasDo.UseModel(&model.CAreas{})

	tableName := _cAreas.cAreasDo.TableName()
	_cAreas.ALL = field.NewAsterisk(tableName)
	_cAreas.ID = field.NewInt64(tableName, "id")
	_cAreas.Code = field.NewString(tableName, "code")
	_cAreas.Name = field.NewString(tableName, "name")
	_cAreas.CityCode = field.NewString(tableName, "city_code")
	_cAreas.ProvinceCode = field.NewString(tableName, "province_code")

	_cAreas.fillFieldMap()

	return _cAreas
}

type cAreas struct {
	cAreasDo

	ALL          field.Asterisk
	ID           field.Int64
	Code         field.String
	Name         field.String
	CityCode     field.String
	ProvinceCode field.String

	fieldMap map[string]field.Expr
}

func (c cAreas) Table(newTableName string) *cAreas {
	c.cAreasDo.UseTable(newTableName)
	return c.updateTableName(newTableName)
}

func (c cAreas) As(alias string) *cAreas {
	c.cAreasDo.DO = *(c.cAreasDo.As(alias).(*gen.DO))
	return c.updateTableName(alias)
}

func (c *cAreas) updateTableName(table string) *cAreas {
	c.ALL = field.NewAsterisk(table)
	c.ID = field.NewInt64(table, "id")
	c.Code = field.NewString(table, "code")
	c.Name = field.NewString(table, "name")
	c.CityCode = field.NewString(table, "city_code")
	c.ProvinceCode = field.NewString(table, "province_code")

	c.fillFieldMap()

	return c
}

func (c *cAreas) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := c.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (c *cAreas) fillFieldMap() {
	c.fieldMap = make(map[string]field.Expr, 5)
	c.fieldMap["id"] = c.ID
	c.fieldMap["code"] = c.Code
	c.fieldMap["name"] = c.Name
	c.fieldMap["city_code"] = c.CityCode
	c.fieldMap["province_code"] = c.ProvinceCode
}

func (c cAreas) clone(db *gorm.DB) cAreas {
	c.cAreasDo.ReplaceConnPool(db.Statement.ConnPool)
	return c
}

func (c cAreas) replaceDB(db *gorm.DB) cAreas {
	c.cAreasDo.ReplaceDB(db)
	return c
}

type cAreasDo struct{ gen.DO }

type ICAreasDo interface {
	gen.SubQuery
	Debug() ICAreasDo
	WithContext(ctx context.Context) ICAreasDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() ICAreasDo
	WriteDB() ICAreasDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) ICAreasDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) ICAreasDo
	Not(conds ...gen.Condition) ICAreasDo
	Or(conds ...gen.Condition) ICAreasDo
	Select(conds ...field.Expr) ICAreasDo
	Where(conds ...gen.Condition) ICAreasDo
	Order(conds ...field.Expr) ICAreasDo
	Distinct(cols ...field.Expr) ICAreasDo
	Omit(cols ...field.Expr) ICAreasDo
	Join(table schema.Tabler, on ...field.Expr) ICAreasDo
	LeftJoin(table schema.Tabler, on ...field.Expr) ICAreasDo
	RightJoin(table schema.Tabler, on ...field.Expr) ICAreasDo
	Group(cols ...field.Expr) ICAreasDo
	Having(conds ...gen.Condition) ICAreasDo
	Limit(limit int) ICAreasDo
	Offset(offset int) ICAreasDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) ICAreasDo
	Unscoped() ICAreasDo
	Create(values ...*model.CAreas) error
	CreateInBatches(values []*model.CAreas, batchSize int) error
	Save(values ...*model.CAreas) error
	First() (*model.CAreas, error)
	Take() (*model.CAreas, error)
	Last() (*model.CAreas, error)
	Find() ([]*model.CAreas, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CAreas, err error)
	FindInBatches(result *[]*model.CAreas, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.CAreas) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) ICAreasDo
	Assign(attrs ...field.AssignExpr) ICAreasDo
	Joins(fields ...field.RelationField) ICAreasDo
	Preload(fields ...field.RelationField) ICAreasDo
	FirstOrInit() (*model.CAreas, error)
	FirstOrCreate() (*model.CAreas, error)
	FindByPage(offset int, limit int) (result []*model.CAreas, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) ICAreasDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (c cAreasDo) Debug() ICAreasDo {
	return c.withDO(c.DO.Debug())
}

func (c cAreasDo) WithContext(ctx context.Context) ICAreasDo {
	return c.withDO(c.DO.WithContext(ctx))
}

func (c cAreasDo) ReadDB() ICAreasDo {
	return c.Clauses(dbresolver.Read)
}

func (c cAreasDo) WriteDB() ICAreasDo {
	return c.Clauses(dbresolver.Write)
}

func (c cAreasDo) Session(config *gorm.Session) ICAreasDo {
	return c.withDO(c.DO.Session(config))
}

func (c cAreasDo) Clauses(conds ...clause.Expression) ICAreasDo {
	return c.withDO(c.DO.Clauses(conds...))
}

func (c cAreasDo) Returning(value interface{}, columns ...string) ICAreasDo {
	return c.withDO(c.DO.Returning(value, columns...))
}

func (c cAreasDo) Not(conds ...gen.Condition) ICAreasDo {
	return c.withDO(c.DO.Not(conds...))
}

func (c cAreasDo) Or(conds ...gen.Condition) ICAreasDo {
	return c.withDO(c.DO.Or(conds...))
}

func (c cAreasDo) Select(conds ...field.Expr) ICAreasDo {
	return c.withDO(c.DO.Select(conds...))
}

func (c cAreasDo) Where(conds ...gen.Condition) ICAreasDo {
	return c.withDO(c.DO.Where(conds...))
}

func (c cAreasDo) Order(conds ...field.Expr) ICAreasDo {
	return c.withDO(c.DO.Order(conds...))
}

func (c cAreasDo) Distinct(cols ...field.Expr) ICAreasDo {
	return c.withDO(c.DO.Distinct(cols...))
}

func (c cAreasDo) Omit(cols ...field.Expr) ICAreasDo {
	return c.withDO(c.DO.Omit(cols...))
}

func (c cAreasDo) Join(table schema.Tabler, on ...field.Expr) ICAreasDo {
	return c.withDO(c.DO.Join(table, on...))
}

func (c cAreasDo) LeftJoin(table schema.Tabler, on ...field.Expr) ICAreasDo {
	return c.withDO(c.DO.LeftJoin(table, on...))
}

func (c cAreasDo) RightJoin(table schema.Tabler, on ...field.Expr) ICAreasDo {
	return c.withDO(c.DO.RightJoin(table, on...))
}

func (c cAreasDo) Group(cols ...field.Expr) ICAreasDo {
	return c.withDO(c.DO.Group(cols...))
}

func (c cAreasDo) Having(conds ...gen.Condition) ICAreasDo {
	return c.withDO(c.DO.Having(conds...))
}

func (c cAreasDo) Limit(limit int) ICAreasDo {
	return c.withDO(c.DO.Limit(limit))
}

func (c cAreasDo) Offset(offset int) ICAreasDo {
	return c.withDO(c.DO.Offset(offset))
}

func (c cAreasDo) Scopes(funcs ...func(gen.Dao) gen.Dao) ICAreasDo {
	return c.withDO(c.DO.Scopes(funcs...))
}

func (c cAreasDo) Unscoped() ICAreasDo {
	return c.withDO(c.DO.Unscoped())
}

func (c cAreasDo) Create(values ...*model.CAreas) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Create(values)
}

func (c cAreasDo) CreateInBatches(values []*model.CAreas, batchSize int) error {
	return c.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (c cAreasDo) Save(values ...*model.CAreas) error {
	if len(values) == 0 {
		return nil
	}
	return c.DO.Save(values)
}

func (c cAreasDo) First() (*model.CAreas, error) {
	if result, err := c.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.CAreas), nil
	}
}

func (c cAreasDo) Take() (*model.CAreas, error) {
	if result, err := c.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.CAreas), nil
	}
}

func (c cAreasDo) Last() (*model.CAreas, error) {
	if result, err := c.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.CAreas), nil
	}
}

func (c cAreasDo) Find() ([]*model.CAreas, error) {
	result, err := c.DO.Find()
	return result.([]*model.CAreas), err
}

func (c cAreasDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.CAreas, err error) {
	buf := make([]*model.CAreas, 0, batchSize)
	err = c.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (c cAreasDo) FindInBatches(result *[]*model.CAreas, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return c.DO.FindInBatches(result, batchSize, fc)
}

func (c cAreasDo) Attrs(attrs ...field.AssignExpr) ICAreasDo {
	return c.withDO(c.DO.Attrs(attrs...))
}

func (c cAreasDo) Assign(attrs ...field.AssignExpr) ICAreasDo {
	return c.withDO(c.DO.Assign(attrs...))
}

func (c cAreasDo) Joins(fields ...field.RelationField) ICAreasDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Joins(_f))
	}
	return &c
}

func (c cAreasDo) Preload(fields ...field.RelationField) ICAreasDo {
	for _, _f := range fields {
		c = *c.withDO(c.DO.Preload(_f))
	}
	return &c
}

func (c cAreasDo) FirstOrInit() (*model.CAreas, error) {
	if result, err := c.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.CAreas), nil
	}
}

func (c cAreasDo) FirstOrCreate() (*model.CAreas, error) {
	if result, err := c.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.CAreas), nil
	}
}

func (c cAreasDo) FindByPage(offset int, limit int) (result []*model.CAreas, count int64, err error) {
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

func (c cAreasDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = c.Count()
	if err != nil {
		return
	}

	err = c.Offset(offset).Limit(limit).Scan(result)
	return
}

func (c cAreasDo) Scan(result interface{}) (err error) {
	return c.DO.Scan(result)
}

func (c cAreasDo) Delete(models ...*model.CAreas) (result gen.ResultInfo, err error) {
	return c.DO.Delete(models)
}

func (c *cAreasDo) withDO(do gen.Dao) *cAreasDo {
	c.DO = *do.(*gen.DO)
	return c
}
