package testdao

import loglibgo "github.com/BF-Moritz/log.lib.go"

func NewDAO() *TestDao {
	return &TestDao{}
}

type TestDao struct {
}

func (dao *TestDao) Test(l *loglibgo.Logger) {
	l.LogDebug("a")
}
