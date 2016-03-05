package chaospeddler

import "database/sql"

//Where - wraps where to return the interface type
func (s *GormDBWrapper) Where(query interface{}, args ...interface{}) GormDB {
	db := s.DBWrapper.DB.Where(query, args...)
	s.DBWrapper = DBWrapper{db}
	return s
}

//DB - a wrapper to delegate down to the gorm.DB
func (s *GormDBWrapper) DB() *sql.DB {
	return s.DBWrapper.DB.DB()
}
