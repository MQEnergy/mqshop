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

func newProductBrand(db *gorm.DB, opts ...gen.DOOption) productBrand {
	_productBrand := productBrand{}

	_productBrand.productBrandDo.UseDB(db, opts...)
	_productBrand.productBrandDo.UseModel(&model.ProductBrand{})

	tableName := _productBrand.productBrandDo.TableName()
	_productBrand.ALL = field.NewAsterisk(tableName)
	_productBrand.ID = field.NewInt64(tableName, "id")
	_productBrand.BrandName = field.NewString(tableName, "brand_name")
	_productBrand.LogoURL = field.NewString(tableName, "logo_url")
	_productBrand.Desc = field.NewString(tableName, "desc")
	_productBrand.SortOrder = field.NewInt(tableName, "sort_order")
	_productBrand.IsHot = field.NewInt8(tableName, "is_hot")
	_productBrand.Status = field.NewInt8(tableName, "status")
	_productBrand.CreatedAt = field.NewInt64(tableName, "created_at")
	_productBrand.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_productBrand.fillFieldMap()

	return _productBrand
}

// productBrand 商品品牌表
type productBrand struct {
	productBrandDo

	ALL       field.Asterisk
	ID        field.Int64
	BrandName field.String // 品牌名称
	LogoURL   field.String // 品牌logo
	Desc      field.String // 品牌描述
	SortOrder field.Int    // 排序
	IsHot     field.Int8   // 是否热门 1：热门 0：非热门
	Status    field.Int8   // 状态 1：正常 0：下架
	CreatedAt field.Int64
	UpdatedAt field.Int64

	fieldMap map[string]field.Expr
}

func (p productBrand) Table(newTableName string) *productBrand {
	p.productBrandDo.UseTable(newTableName)
	return p.updateTableName(newTableName)
}

func (p productBrand) As(alias string) *productBrand {
	p.productBrandDo.DO = *(p.productBrandDo.As(alias).(*gen.DO))
	return p.updateTableName(alias)
}

func (p *productBrand) updateTableName(table string) *productBrand {
	p.ALL = field.NewAsterisk(table)
	p.ID = field.NewInt64(table, "id")
	p.BrandName = field.NewString(table, "brand_name")
	p.LogoURL = field.NewString(table, "logo_url")
	p.Desc = field.NewString(table, "desc")
	p.SortOrder = field.NewInt(table, "sort_order")
	p.IsHot = field.NewInt8(table, "is_hot")
	p.Status = field.NewInt8(table, "status")
	p.CreatedAt = field.NewInt64(table, "created_at")
	p.UpdatedAt = field.NewInt64(table, "updated_at")

	p.fillFieldMap()

	return p
}

func (p *productBrand) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := p.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (p *productBrand) fillFieldMap() {
	p.fieldMap = make(map[string]field.Expr, 9)
	p.fieldMap["id"] = p.ID
	p.fieldMap["brand_name"] = p.BrandName
	p.fieldMap["logo_url"] = p.LogoURL
	p.fieldMap["desc"] = p.Desc
	p.fieldMap["sort_order"] = p.SortOrder
	p.fieldMap["is_hot"] = p.IsHot
	p.fieldMap["status"] = p.Status
	p.fieldMap["created_at"] = p.CreatedAt
	p.fieldMap["updated_at"] = p.UpdatedAt
}

func (p productBrand) clone(db *gorm.DB) productBrand {
	p.productBrandDo.ReplaceConnPool(db.Statement.ConnPool)
	return p
}

func (p productBrand) replaceDB(db *gorm.DB) productBrand {
	p.productBrandDo.ReplaceDB(db)
	return p
}

type productBrandDo struct{ gen.DO }

type IProductBrandDo interface {
	gen.SubQuery
	Debug() IProductBrandDo
	WithContext(ctx context.Context) IProductBrandDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IProductBrandDo
	WriteDB() IProductBrandDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IProductBrandDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IProductBrandDo
	Not(conds ...gen.Condition) IProductBrandDo
	Or(conds ...gen.Condition) IProductBrandDo
	Select(conds ...field.Expr) IProductBrandDo
	Where(conds ...gen.Condition) IProductBrandDo
	Order(conds ...field.Expr) IProductBrandDo
	Distinct(cols ...field.Expr) IProductBrandDo
	Omit(cols ...field.Expr) IProductBrandDo
	Join(table schema.Tabler, on ...field.Expr) IProductBrandDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IProductBrandDo
	RightJoin(table schema.Tabler, on ...field.Expr) IProductBrandDo
	Group(cols ...field.Expr) IProductBrandDo
	Having(conds ...gen.Condition) IProductBrandDo
	Limit(limit int) IProductBrandDo
	Offset(offset int) IProductBrandDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IProductBrandDo
	Unscoped() IProductBrandDo
	Create(values ...*model.ProductBrand) error
	CreateInBatches(values []*model.ProductBrand, batchSize int) error
	Save(values ...*model.ProductBrand) error
	First() (*model.ProductBrand, error)
	Take() (*model.ProductBrand, error)
	Last() (*model.ProductBrand, error)
	Find() ([]*model.ProductBrand, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProductBrand, err error)
	FindInBatches(result *[]*model.ProductBrand, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.ProductBrand) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IProductBrandDo
	Assign(attrs ...field.AssignExpr) IProductBrandDo
	Joins(fields ...field.RelationField) IProductBrandDo
	Preload(fields ...field.RelationField) IProductBrandDo
	FirstOrInit() (*model.ProductBrand, error)
	FirstOrCreate() (*model.ProductBrand, error)
	FindByPage(offset int, limit int) (result []*model.ProductBrand, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IProductBrandDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (p productBrandDo) Debug() IProductBrandDo {
	return p.withDO(p.DO.Debug())
}

func (p productBrandDo) WithContext(ctx context.Context) IProductBrandDo {
	return p.withDO(p.DO.WithContext(ctx))
}

func (p productBrandDo) ReadDB() IProductBrandDo {
	return p.Clauses(dbresolver.Read)
}

func (p productBrandDo) WriteDB() IProductBrandDo {
	return p.Clauses(dbresolver.Write)
}

func (p productBrandDo) Session(config *gorm.Session) IProductBrandDo {
	return p.withDO(p.DO.Session(config))
}

func (p productBrandDo) Clauses(conds ...clause.Expression) IProductBrandDo {
	return p.withDO(p.DO.Clauses(conds...))
}

func (p productBrandDo) Returning(value interface{}, columns ...string) IProductBrandDo {
	return p.withDO(p.DO.Returning(value, columns...))
}

func (p productBrandDo) Not(conds ...gen.Condition) IProductBrandDo {
	return p.withDO(p.DO.Not(conds...))
}

func (p productBrandDo) Or(conds ...gen.Condition) IProductBrandDo {
	return p.withDO(p.DO.Or(conds...))
}

func (p productBrandDo) Select(conds ...field.Expr) IProductBrandDo {
	return p.withDO(p.DO.Select(conds...))
}

func (p productBrandDo) Where(conds ...gen.Condition) IProductBrandDo {
	return p.withDO(p.DO.Where(conds...))
}

func (p productBrandDo) Order(conds ...field.Expr) IProductBrandDo {
	return p.withDO(p.DO.Order(conds...))
}

func (p productBrandDo) Distinct(cols ...field.Expr) IProductBrandDo {
	return p.withDO(p.DO.Distinct(cols...))
}

func (p productBrandDo) Omit(cols ...field.Expr) IProductBrandDo {
	return p.withDO(p.DO.Omit(cols...))
}

func (p productBrandDo) Join(table schema.Tabler, on ...field.Expr) IProductBrandDo {
	return p.withDO(p.DO.Join(table, on...))
}

func (p productBrandDo) LeftJoin(table schema.Tabler, on ...field.Expr) IProductBrandDo {
	return p.withDO(p.DO.LeftJoin(table, on...))
}

func (p productBrandDo) RightJoin(table schema.Tabler, on ...field.Expr) IProductBrandDo {
	return p.withDO(p.DO.RightJoin(table, on...))
}

func (p productBrandDo) Group(cols ...field.Expr) IProductBrandDo {
	return p.withDO(p.DO.Group(cols...))
}

func (p productBrandDo) Having(conds ...gen.Condition) IProductBrandDo {
	return p.withDO(p.DO.Having(conds...))
}

func (p productBrandDo) Limit(limit int) IProductBrandDo {
	return p.withDO(p.DO.Limit(limit))
}

func (p productBrandDo) Offset(offset int) IProductBrandDo {
	return p.withDO(p.DO.Offset(offset))
}

func (p productBrandDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IProductBrandDo {
	return p.withDO(p.DO.Scopes(funcs...))
}

func (p productBrandDo) Unscoped() IProductBrandDo {
	return p.withDO(p.DO.Unscoped())
}

func (p productBrandDo) Create(values ...*model.ProductBrand) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Create(values)
}

func (p productBrandDo) CreateInBatches(values []*model.ProductBrand, batchSize int) error {
	return p.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (p productBrandDo) Save(values ...*model.ProductBrand) error {
	if len(values) == 0 {
		return nil
	}
	return p.DO.Save(values)
}

func (p productBrandDo) First() (*model.ProductBrand, error) {
	if result, err := p.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductBrand), nil
	}
}

func (p productBrandDo) Take() (*model.ProductBrand, error) {
	if result, err := p.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductBrand), nil
	}
}

func (p productBrandDo) Last() (*model.ProductBrand, error) {
	if result, err := p.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductBrand), nil
	}
}

func (p productBrandDo) Find() ([]*model.ProductBrand, error) {
	result, err := p.DO.Find()
	return result.([]*model.ProductBrand), err
}

func (p productBrandDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.ProductBrand, err error) {
	buf := make([]*model.ProductBrand, 0, batchSize)
	err = p.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (p productBrandDo) FindInBatches(result *[]*model.ProductBrand, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return p.DO.FindInBatches(result, batchSize, fc)
}

func (p productBrandDo) Attrs(attrs ...field.AssignExpr) IProductBrandDo {
	return p.withDO(p.DO.Attrs(attrs...))
}

func (p productBrandDo) Assign(attrs ...field.AssignExpr) IProductBrandDo {
	return p.withDO(p.DO.Assign(attrs...))
}

func (p productBrandDo) Joins(fields ...field.RelationField) IProductBrandDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Joins(_f))
	}
	return &p
}

func (p productBrandDo) Preload(fields ...field.RelationField) IProductBrandDo {
	for _, _f := range fields {
		p = *p.withDO(p.DO.Preload(_f))
	}
	return &p
}

func (p productBrandDo) FirstOrInit() (*model.ProductBrand, error) {
	if result, err := p.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductBrand), nil
	}
}

func (p productBrandDo) FirstOrCreate() (*model.ProductBrand, error) {
	if result, err := p.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.ProductBrand), nil
	}
}

func (p productBrandDo) FindByPage(offset int, limit int) (result []*model.ProductBrand, count int64, err error) {
	result, err = p.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = p.Offset(-1).Limit(-1).Count()
	return
}

func (p productBrandDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = p.Count()
	if err != nil {
		return
	}

	err = p.Offset(offset).Limit(limit).Scan(result)
	return
}

func (p productBrandDo) Scan(result interface{}) (err error) {
	return p.DO.Scan(result)
}

func (p productBrandDo) Delete(models ...*model.ProductBrand) (result gen.ResultInfo, err error) {
	return p.DO.Delete(models)
}

func (p *productBrandDo) withDO(do gen.Dao) *productBrandDo {
	p.DO = *do.(*gen.DO)
	return p
}
