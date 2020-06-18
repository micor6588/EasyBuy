package commons

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

//数据库操作的三个对象
var (
	db   *sql.DB
	stmt *sql.Stmt
	rows *sql.Rows
)

//打开数据库连接,不要忘记导入驱动包
func openConn() (err error) {
	//此处为等号,否则创建局部变量
	db, err = sql.Open("mysql", "root:371871@tcp(localhost:3306)/easybuy?charset=utf8")
	if err != nil {
		fmt.Println("连接失败", err)
		return
	}
	return nil
}

//关闭连接,首字母大写,需要跨包访问的
func CloseConn() {
	if rows != nil {
		rows.Close()
	}
	if stmt != nil {
		stmt.Close()
	}
	if db != nil {
		db.Close()
	}
}

//执行DML新增,删除,修改操作
func Dml(sql string, args ...interface{}) (int64, error) {
	err := openConn()
	if err != nil {
		fmt.Println("执行DML时出现错误,打开连接失败")
		return 0, err
	}
	//此处也是等号
	stmt, err = db.Prepare(sql)
	if err != nil {
		fmt.Println("执行DML时出现错误,预处理出现错误")
		return 0, err
	}
	//此处要有...表示切片,如果没有表示数组,会报错
	result, err := stmt.Exec(args...)
	if err != nil {
		fmt.Println("执行DML出现错误,执行错误")
		return 0, err
	}
	count, err := result.RowsAffected()
	if err != nil {
		fmt.Println("执行DML出现错误,获取受影响行数错误")
		return 0, err
	}
	CloseConn() //关闭连接
	return count, err
}

//执行DQL查询
func Dql(sql string, args ...interface{}) (*sql.Rows, error) {
	err := openConn()
	if err != nil {
		fmt.Println("执行DQL出现错误,打开连接失败")
		return nil, err
	}
	//此处是等号
	stmt, err = db.Prepare(sql)
	if err != nil {
		fmt.Println("执行DQL出现错误,预处理实现")
		return nil, err
	}
	//此处参数是切片
	rows, err = stmt.Query(args...)
	if err != nil {
		fmt.Println("执行DQL出现错误,执行错误")
		return nil, err
	}
	//此处没有关闭,调用此函数要记得关闭连接
	return rows, nil
}
