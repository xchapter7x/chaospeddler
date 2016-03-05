package chaospeddler

//Where - wraps where to return the interface type
func (s *GormDBWrapper) Where(query interface{}, args ...interface{}) GormDB {
	return s.Where(query, args...)
}
