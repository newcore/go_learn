package handler

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"strconv"
	"strings"
	"sync"
	"time"
)

var (
	userName  string = "root"
	password  string = "123456"
	ipAddrees string = "localhost"
	port      int    = 3306
	dbName    string = "test"
	charset   string = "utf8"
)

type Szlc struct {
	Id int
	LightProductCode string
	LightProductName string
	TrendDESProductId string
	Create_time string
}

func connect_mysql()  *sqlx.DB{
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s", userName, password, ipAddrees, port, dbName, charset)
	Db, err := sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("mysql connect failed, detail is [%v]", err.Error())
	}
	return Db
}
func Run(){
	var Db *sqlx.DB = connect_mysql()
	defer Db.Close()
	rows,_ := Db.Query("select create_time from t_data_szlc_go_all GROUP BY create_time;")
	for rows.Next(){
		var create_time string
		err := rows.Scan(&create_time)
		if err != nil{
			fmt.Println("查询create_time失败")
		}
		fmt.Println(create_time)
		createTable(Db,create_time)
	}
	rows.Close()
	row := Db.QueryRow("select max(id) as max_id from t_data_szlc_go_all")
	var maxId int
	row.Scan(&maxId)
	fmt.Println(maxId)
	var start int
	var wg sync.WaitGroup
	for {
		if start> maxId{
			break
		} else {
			wg.Add(1)
			go doWork(start,Db,&wg)
			start += 10000
		}
	}
	wg.Wait()
	fmt.Println("over")

}

func doWork(start int,Db *sqlx.DB,wg *sync.WaitGroup)  {
	fmt.Println("start",start)
	rows,err := Db.Query("select * from t_data_szlc_go_all where id >= "+strconv.Itoa(start)+" limit 10000")
	if err != nil{
		fmt.Println("query err",err.Error())
		return
	}
	for rows.Next(){
		var szlc Szlc
		err := rows.Scan(&szlc.Id,&szlc.LightProductCode,&szlc.LightProductName,&szlc.TrendDESProductId,&szlc.Create_time)
		if err != nil{
			fmt.Println("read err:",err.Error())
		}
		//fmt.Println(szlc.Id,szlc.LightProductCode,szlc.LightProductName,szlc.Create_time)
		insert_one(Db,&szlc)
	}
	rows.Close()
	wg.Done()
}

func insert_one(db *sqlx.DB ,szlc *Szlc)  {
	date := szlc.Create_time
	tableName := getTable(date)
	fmt.Println(tableName)
	_,err := db.Exec("insert into "+tableName+"  values(?,?,?,?,?)",szlc.Id,szlc.LightProductCode,szlc.LightProductName,szlc.TrendDESProductId,szlc.Create_time)
	if err != nil{
		fmt.Println("insert err:",err.Error())
		panic(err)
	}
	fmt.Println("保存成功")
}

func getTable(date string) string {
	loc,_ := time.LoadLocation("Local")
	create_date, _ := time.ParseInLocation("2006-01-02", date, loc) //使用模板在对应时区转化为time.time类型
	year := create_date.Year()
	month := int(create_date.Month())
	var month_str string
	if month < 10{
		month_str = "0"+strconv.Itoa(month)
	} else {
		month_str = strconv.Itoa(month)
	}
	table_name := "t_data_szlc_" + strconv.Itoa(year)+ month_str
	return table_name
}

func createTable(db *sqlx.DB,date string){
	tableName := getTable(date)
	sql := `create table if not EXISTS table_name(
	id int(11) unsigned NOT NULL DEFAULT 0,
		lightProductCode varchar(255) DEFAULT NULL,
		lightProductName varchar(255) DEFAULT NULL,
		trendDESProductId varchar(255) DEFAULT NULL,
		create_time date DEFAULT NULL
	) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	sql = strings.Replace(sql,"table_name",tableName,1)
	fmt.Println(sql)
	_,err := db.Exec(sql)
	if err != nil{
		fmt.Println("创建表失败")
		panic(err.Error())
	}
	fmt.Println("创建表成功")
}