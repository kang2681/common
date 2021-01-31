package drivers

import "github.com/jinzhu/gorm"

//NewMysqlConn 创建MYSQL连接
func NewMysqlConn(user, password, host, dbname string, maxConn, minConn int) (*gorm.DB, error) {
	url := user + ":" + password + "@tcp" + "(" + host + ")/" + dbname + "?&charset=" + charSet + "&parseTime=True&loc=Asia%2FChongqing"
	db, err := gorm.Open("mysql", url)
	if err != nil {
		return nil, err
	}
	db.DB().SetMaxOpenConns(maxConn)
	db.DB().SetMaxIdleConns(minConn)
	return db, err
}
