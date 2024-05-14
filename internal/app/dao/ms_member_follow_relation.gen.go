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

func newMemberFollowRelation(db *gorm.DB, opts ...gen.DOOption) memberFollowRelation {
	_memberFollowRelation := memberFollowRelation{}

	_memberFollowRelation.memberFollowRelationDo.UseDB(db, opts...)
	_memberFollowRelation.memberFollowRelationDo.UseModel(&model.MemberFollowRelation{})

	tableName := _memberFollowRelation.memberFollowRelationDo.TableName()
	_memberFollowRelation.ALL = field.NewAsterisk(tableName)
	_memberFollowRelation.ID = field.NewInt64(tableName, "id")
	_memberFollowRelation.FollowMemberID = field.NewInt64(tableName, "follow_member_id")
	_memberFollowRelation.MemberID = field.NewInt64(tableName, "member_id")
	_memberFollowRelation.CreatedAt = field.NewInt64(tableName, "created_at")
	_memberFollowRelation.UpdatedAt = field.NewInt64(tableName, "updated_at")

	_memberFollowRelation.fillFieldMap()

	return _memberFollowRelation
}

// memberFollowRelation 用户关注关联表
type memberFollowRelation struct {
	memberFollowRelationDo

	ALL            field.Asterisk
	ID             field.Int64
	FollowMemberID field.Int64 // 被关注者ID
	MemberID       field.Int64 // 关注者ID
	CreatedAt      field.Int64
	UpdatedAt      field.Int64

	fieldMap map[string]field.Expr
}

func (m memberFollowRelation) Table(newTableName string) *memberFollowRelation {
	m.memberFollowRelationDo.UseTable(newTableName)
	return m.updateTableName(newTableName)
}

func (m memberFollowRelation) As(alias string) *memberFollowRelation {
	m.memberFollowRelationDo.DO = *(m.memberFollowRelationDo.As(alias).(*gen.DO))
	return m.updateTableName(alias)
}

func (m *memberFollowRelation) updateTableName(table string) *memberFollowRelation {
	m.ALL = field.NewAsterisk(table)
	m.ID = field.NewInt64(table, "id")
	m.FollowMemberID = field.NewInt64(table, "follow_member_id")
	m.MemberID = field.NewInt64(table, "member_id")
	m.CreatedAt = field.NewInt64(table, "created_at")
	m.UpdatedAt = field.NewInt64(table, "updated_at")

	m.fillFieldMap()

	return m
}

func (m *memberFollowRelation) GetFieldByName(fieldName string) (field.OrderExpr, bool) {
	_f, ok := m.fieldMap[fieldName]
	if !ok || _f == nil {
		return nil, false
	}
	_oe, ok := _f.(field.OrderExpr)
	return _oe, ok
}

func (m *memberFollowRelation) fillFieldMap() {
	m.fieldMap = make(map[string]field.Expr, 5)
	m.fieldMap["id"] = m.ID
	m.fieldMap["follow_member_id"] = m.FollowMemberID
	m.fieldMap["member_id"] = m.MemberID
	m.fieldMap["created_at"] = m.CreatedAt
	m.fieldMap["updated_at"] = m.UpdatedAt
}

func (m memberFollowRelation) clone(db *gorm.DB) memberFollowRelation {
	m.memberFollowRelationDo.ReplaceConnPool(db.Statement.ConnPool)
	return m
}

func (m memberFollowRelation) replaceDB(db *gorm.DB) memberFollowRelation {
	m.memberFollowRelationDo.ReplaceDB(db)
	return m
}

type memberFollowRelationDo struct{ gen.DO }

type IMemberFollowRelationDo interface {
	gen.SubQuery
	Debug() IMemberFollowRelationDo
	WithContext(ctx context.Context) IMemberFollowRelationDo
	WithResult(fc func(tx gen.Dao)) gen.ResultInfo
	ReplaceDB(db *gorm.DB)
	ReadDB() IMemberFollowRelationDo
	WriteDB() IMemberFollowRelationDo
	As(alias string) gen.Dao
	Session(config *gorm.Session) IMemberFollowRelationDo
	Columns(cols ...field.Expr) gen.Columns
	Clauses(conds ...clause.Expression) IMemberFollowRelationDo
	Not(conds ...gen.Condition) IMemberFollowRelationDo
	Or(conds ...gen.Condition) IMemberFollowRelationDo
	Select(conds ...field.Expr) IMemberFollowRelationDo
	Where(conds ...gen.Condition) IMemberFollowRelationDo
	Order(conds ...field.Expr) IMemberFollowRelationDo
	Distinct(cols ...field.Expr) IMemberFollowRelationDo
	Omit(cols ...field.Expr) IMemberFollowRelationDo
	Join(table schema.Tabler, on ...field.Expr) IMemberFollowRelationDo
	LeftJoin(table schema.Tabler, on ...field.Expr) IMemberFollowRelationDo
	RightJoin(table schema.Tabler, on ...field.Expr) IMemberFollowRelationDo
	Group(cols ...field.Expr) IMemberFollowRelationDo
	Having(conds ...gen.Condition) IMemberFollowRelationDo
	Limit(limit int) IMemberFollowRelationDo
	Offset(offset int) IMemberFollowRelationDo
	Count() (count int64, err error)
	Scopes(funcs ...func(gen.Dao) gen.Dao) IMemberFollowRelationDo
	Unscoped() IMemberFollowRelationDo
	Create(values ...*model.MemberFollowRelation) error
	CreateInBatches(values []*model.MemberFollowRelation, batchSize int) error
	Save(values ...*model.MemberFollowRelation) error
	First() (*model.MemberFollowRelation, error)
	Take() (*model.MemberFollowRelation, error)
	Last() (*model.MemberFollowRelation, error)
	Find() ([]*model.MemberFollowRelation, error)
	FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MemberFollowRelation, err error)
	FindInBatches(result *[]*model.MemberFollowRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error
	Pluck(column field.Expr, dest interface{}) error
	Delete(...*model.MemberFollowRelation) (info gen.ResultInfo, err error)
	Update(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	Updates(value interface{}) (info gen.ResultInfo, err error)
	UpdateColumn(column field.Expr, value interface{}) (info gen.ResultInfo, err error)
	UpdateColumnSimple(columns ...field.AssignExpr) (info gen.ResultInfo, err error)
	UpdateColumns(value interface{}) (info gen.ResultInfo, err error)
	UpdateFrom(q gen.SubQuery) gen.Dao
	Attrs(attrs ...field.AssignExpr) IMemberFollowRelationDo
	Assign(attrs ...field.AssignExpr) IMemberFollowRelationDo
	Joins(fields ...field.RelationField) IMemberFollowRelationDo
	Preload(fields ...field.RelationField) IMemberFollowRelationDo
	FirstOrInit() (*model.MemberFollowRelation, error)
	FirstOrCreate() (*model.MemberFollowRelation, error)
	FindByPage(offset int, limit int) (result []*model.MemberFollowRelation, count int64, err error)
	ScanByPage(result interface{}, offset int, limit int) (count int64, err error)
	Scan(result interface{}) (err error)
	Returning(value interface{}, columns ...string) IMemberFollowRelationDo
	UnderlyingDB() *gorm.DB
	schema.Tabler
}

func (m memberFollowRelationDo) Debug() IMemberFollowRelationDo {
	return m.withDO(m.DO.Debug())
}

func (m memberFollowRelationDo) WithContext(ctx context.Context) IMemberFollowRelationDo {
	return m.withDO(m.DO.WithContext(ctx))
}

func (m memberFollowRelationDo) ReadDB() IMemberFollowRelationDo {
	return m.Clauses(dbresolver.Read)
}

func (m memberFollowRelationDo) WriteDB() IMemberFollowRelationDo {
	return m.Clauses(dbresolver.Write)
}

func (m memberFollowRelationDo) Session(config *gorm.Session) IMemberFollowRelationDo {
	return m.withDO(m.DO.Session(config))
}

func (m memberFollowRelationDo) Clauses(conds ...clause.Expression) IMemberFollowRelationDo {
	return m.withDO(m.DO.Clauses(conds...))
}

func (m memberFollowRelationDo) Returning(value interface{}, columns ...string) IMemberFollowRelationDo {
	return m.withDO(m.DO.Returning(value, columns...))
}

func (m memberFollowRelationDo) Not(conds ...gen.Condition) IMemberFollowRelationDo {
	return m.withDO(m.DO.Not(conds...))
}

func (m memberFollowRelationDo) Or(conds ...gen.Condition) IMemberFollowRelationDo {
	return m.withDO(m.DO.Or(conds...))
}

func (m memberFollowRelationDo) Select(conds ...field.Expr) IMemberFollowRelationDo {
	return m.withDO(m.DO.Select(conds...))
}

func (m memberFollowRelationDo) Where(conds ...gen.Condition) IMemberFollowRelationDo {
	return m.withDO(m.DO.Where(conds...))
}

func (m memberFollowRelationDo) Order(conds ...field.Expr) IMemberFollowRelationDo {
	return m.withDO(m.DO.Order(conds...))
}

func (m memberFollowRelationDo) Distinct(cols ...field.Expr) IMemberFollowRelationDo {
	return m.withDO(m.DO.Distinct(cols...))
}

func (m memberFollowRelationDo) Omit(cols ...field.Expr) IMemberFollowRelationDo {
	return m.withDO(m.DO.Omit(cols...))
}

func (m memberFollowRelationDo) Join(table schema.Tabler, on ...field.Expr) IMemberFollowRelationDo {
	return m.withDO(m.DO.Join(table, on...))
}

func (m memberFollowRelationDo) LeftJoin(table schema.Tabler, on ...field.Expr) IMemberFollowRelationDo {
	return m.withDO(m.DO.LeftJoin(table, on...))
}

func (m memberFollowRelationDo) RightJoin(table schema.Tabler, on ...field.Expr) IMemberFollowRelationDo {
	return m.withDO(m.DO.RightJoin(table, on...))
}

func (m memberFollowRelationDo) Group(cols ...field.Expr) IMemberFollowRelationDo {
	return m.withDO(m.DO.Group(cols...))
}

func (m memberFollowRelationDo) Having(conds ...gen.Condition) IMemberFollowRelationDo {
	return m.withDO(m.DO.Having(conds...))
}

func (m memberFollowRelationDo) Limit(limit int) IMemberFollowRelationDo {
	return m.withDO(m.DO.Limit(limit))
}

func (m memberFollowRelationDo) Offset(offset int) IMemberFollowRelationDo {
	return m.withDO(m.DO.Offset(offset))
}

func (m memberFollowRelationDo) Scopes(funcs ...func(gen.Dao) gen.Dao) IMemberFollowRelationDo {
	return m.withDO(m.DO.Scopes(funcs...))
}

func (m memberFollowRelationDo) Unscoped() IMemberFollowRelationDo {
	return m.withDO(m.DO.Unscoped())
}

func (m memberFollowRelationDo) Create(values ...*model.MemberFollowRelation) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Create(values)
}

func (m memberFollowRelationDo) CreateInBatches(values []*model.MemberFollowRelation, batchSize int) error {
	return m.DO.CreateInBatches(values, batchSize)
}

// Save : !!! underlying implementation is different with GORM
// The method is equivalent to executing the statement: db.Clauses(clause.OnConflict{UpdateAll: true}).Create(values)
func (m memberFollowRelationDo) Save(values ...*model.MemberFollowRelation) error {
	if len(values) == 0 {
		return nil
	}
	return m.DO.Save(values)
}

func (m memberFollowRelationDo) First() (*model.MemberFollowRelation, error) {
	if result, err := m.DO.First(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberFollowRelation), nil
	}
}

func (m memberFollowRelationDo) Take() (*model.MemberFollowRelation, error) {
	if result, err := m.DO.Take(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberFollowRelation), nil
	}
}

func (m memberFollowRelationDo) Last() (*model.MemberFollowRelation, error) {
	if result, err := m.DO.Last(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberFollowRelation), nil
	}
}

func (m memberFollowRelationDo) Find() ([]*model.MemberFollowRelation, error) {
	result, err := m.DO.Find()
	return result.([]*model.MemberFollowRelation), err
}

func (m memberFollowRelationDo) FindInBatch(batchSize int, fc func(tx gen.Dao, batch int) error) (results []*model.MemberFollowRelation, err error) {
	buf := make([]*model.MemberFollowRelation, 0, batchSize)
	err = m.DO.FindInBatches(&buf, batchSize, func(tx gen.Dao, batch int) error {
		defer func() { results = append(results, buf...) }()
		return fc(tx, batch)
	})
	return results, err
}

func (m memberFollowRelationDo) FindInBatches(result *[]*model.MemberFollowRelation, batchSize int, fc func(tx gen.Dao, batch int) error) error {
	return m.DO.FindInBatches(result, batchSize, fc)
}

func (m memberFollowRelationDo) Attrs(attrs ...field.AssignExpr) IMemberFollowRelationDo {
	return m.withDO(m.DO.Attrs(attrs...))
}

func (m memberFollowRelationDo) Assign(attrs ...field.AssignExpr) IMemberFollowRelationDo {
	return m.withDO(m.DO.Assign(attrs...))
}

func (m memberFollowRelationDo) Joins(fields ...field.RelationField) IMemberFollowRelationDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Joins(_f))
	}
	return &m
}

func (m memberFollowRelationDo) Preload(fields ...field.RelationField) IMemberFollowRelationDo {
	for _, _f := range fields {
		m = *m.withDO(m.DO.Preload(_f))
	}
	return &m
}

func (m memberFollowRelationDo) FirstOrInit() (*model.MemberFollowRelation, error) {
	if result, err := m.DO.FirstOrInit(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberFollowRelation), nil
	}
}

func (m memberFollowRelationDo) FirstOrCreate() (*model.MemberFollowRelation, error) {
	if result, err := m.DO.FirstOrCreate(); err != nil {
		return nil, err
	} else {
		return result.(*model.MemberFollowRelation), nil
	}
}

func (m memberFollowRelationDo) FindByPage(offset int, limit int) (result []*model.MemberFollowRelation, count int64, err error) {
	result, err = m.Offset(offset).Limit(limit).Find()
	if err != nil {
		return
	}

	if size := len(result); 0 < limit && 0 < size && size < limit {
		count = int64(size + offset)
		return
	}

	count, err = m.Offset(-1).Limit(-1).Count()
	return
}

func (m memberFollowRelationDo) ScanByPage(result interface{}, offset int, limit int) (count int64, err error) {
	count, err = m.Count()
	if err != nil {
		return
	}

	err = m.Offset(offset).Limit(limit).Scan(result)
	return
}

func (m memberFollowRelationDo) Scan(result interface{}) (err error) {
	return m.DO.Scan(result)
}

func (m memberFollowRelationDo) Delete(models ...*model.MemberFollowRelation) (result gen.ResultInfo, err error) {
	return m.DO.Delete(models)
}

func (m *memberFollowRelationDo) withDO(do gen.Dao) *memberFollowRelationDo {
	m.DO = *do.(*gen.DO)
	return m
}
