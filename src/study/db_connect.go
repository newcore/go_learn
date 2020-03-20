package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	userName string = "root"
	passWord string = "xxxxxxx"
	ipAddress string = "xxxxxxxx"
	port int = 3306
	dbName string = "test"
	charset string = "utf8"
)

func ConnectMysql() (*sqlx.DB){
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s",userName,passWord,ipAddress,port,dbName,charset)
	Db,err := sqlx.Open("mysql",dsn)
	if err != nil{
		fmt.Printf("mysql connect failed ,detail is [%v]",err.Error())
	}
	return Db
}

func addRecord(Db *sqlx.DB){
	result ,err := Db.Exec("insert into article (title,content,keywords,come) values (?,?,?,?)","标题测试","内容222","keyword","aobama")
	if err != nil{
		fmt.Printf("insert data failed ,detail is %v",err.Error())
		return
	}
	id,_ := result.LastInsertId()
	fmt.Printf("last insert id is:[%d]\n",id)
	return
}

func updateRecord(Db *sqlx.DB){
	result,err := Db.Exec("update article set ctime = '2020-03-20 23:51:08' where aid = 16")
	if err != nil{
		fmt.Printf("update err ,detail is [%v]",err.Error())
		return
	}
	num,_ := result.RowsAffected()
	fmt.Printf("update success,affected rows %d",num)
}

func queryData(Db *sqlx.DB)  {
	rows,err := Db.Query("select aid,title,content,keywords from article")
	if err != nil{
		fmt.Printf("query data error,detail is [%v]",err.Error())
		return
	}
	for rows.Next(){
		var aid int
		var title,content,keywords string
		err = rows.Scan(&aid,&title,&content,&keywords)
		if err != nil {
			fmt.Println("get data failed, error:[%v]", err.Error())
		}
		fmt.Println("id:",aid,title,keywords)
	}
	rows.Close()
}

func main() {
	Db := ConnectMysql()
	defer Db.Close()
	//addRecord(Db)
	//updateRecord(Db)
	queryData(Db)
}
