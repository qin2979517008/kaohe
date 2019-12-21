package main

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"

)

func main(){
	db, err := sql.Open("mysql", "root:123456@tcp(localhost:3306)/food?charset=utf8")
	defer db.Close()
	if(err != nil){
		fmt.Println("打开数据库失败")
		fmt.Println(err)
	}
 insert(db)
	 xuanzhe(db)

}
func insert(db *sql.DB)  {
	stmt,err := db.Prepare("INSERT delicious (name,price) values (?,?)")
	if (err !=nil){
	 fmt.Println(err)
	}
	stmt.Exec("小面",6)
	stmt.Exec("饭团",7)
	stmt.Exec("香菇滑鸡",12)
	stmt.Exec("小炒肉",15)
	stmt.Exec("黄焖鸡",16)
	stmt.Exec("冒菜",18)
	fmt.Println("插入食物菜单成功")
}

func xuanzhe(db *sql.DB)  {

		stmt, err := db.Query("select * from delicious;")
		if err != nil {
			fmt.Println(err)
		}
		shiwu := make(map[string]int)
		defer stmt.Close()
		for stmt.Next(){
			var name string
			var price int
			err := stmt.Scan(&name,&price)
			if err != nil {
				fmt.Println(err)
			}
			//fmt.Println(name,price)
			//fmt.Println()
			 shiwu[name] =price

				 }
//遍历食物数组，挑选两样价格不超过二十元的食物
	for i,j :=range shiwu{
		for a,b :=range shiwu{
			if j+b <20&&j!=b{
				fmt.Println("学长今天中午的食物是")
				fmt.Println(i+"和"+a)
			  }

		  }
		}
	}

