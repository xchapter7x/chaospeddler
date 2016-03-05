package chaospeddler

import "database/sql"

//Where - wraps where to return the interface type
func (s *GormDBWrapper) Where(query interface{}, args ...interface{}) GormDB {
	return s.Where(query, args...)
}

//DB - a wrapper to delegate down to the gorm.DB
func (s *GormDBWrapper) DB() *sql.DB {
	return s.DBWrapper.DB.DB()
}
