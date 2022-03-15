package main
// xorm安装
// 源代码"github.com/go-sql-driver/mysql"
// go get github.com/go-xorm/xorm

// mysql连接示例
// ·创建引擎
engine, err := xorm.NewEngine("mysql", "用户名:密码@/数据库名称?charset=utf-8")
if err != nil {
	panic(err.Error())
}
数据库引擎关闭
defer engine.Close()



数据库引擎设置
engine.ShowSQL(true)是否显示sql语句
engine.Logger().SetLevel(core.LOG_DEBUG)设置日志级别
engine.SetMaxOpenConns(10)最大连接数

同步
engine.sync2()


映射
engine.SetMapper(core.SnakeMapper{})驼峰式
engine.SetMapper(core.SameMapper{})相同的

表的定义
type UserTable struct {
	UserId int64 `xorm:"pk autoincr"`(pk是主键，后面是自增)
	UserName string `xorm:"varchar(32)"`
	UserAge int64 `xorm:"default 1"`(某列的默认值)
}

















// 3月15日，trpc-go的客户端client成功实现，华哥安排数据库任务，自己尝试编写helloworld没成功
