package main

import (
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
	"strconv"
	"time"
)

type User struct {
	Id    int64
	Name  string
	Ctime int64
}

func main() {
	var err error
	engine, err := xorm.NewEngine("mysql", "root:@tcp(127.0.0.1:3306)/blog?charset=utf8mb4&parseTime=true&&loc=Local")
	//results, err := engine.QueryString("select * from user")
	fmt.Println(err)
	//fmt.Println(results)

	ctime := time.Now().Unix()
	var users []User
	var user User
	for i := 0; i < 4; i++ {
		user = User{Name: "jerry" + strconv.Itoa(i), Ctime: ctime}
		users = append(users, user)
	}

	affected, err := engine.Insert(&user)
	fmt.Println(affected)

	affected, err = engine.Insert(&users)
	fmt.Println(affected)

	user = User{Name: "piter", Ctime: ctime}
	affected, err = engine.Id(1).Update(&user)

	affected, err = engine.Update(&user, &User{Name: "jerry1"})

	user = User{}
	has, err := engine.Where("name = ?", "jerry0").Desc("id").Get(&user)

	user = User{Name:"jerry0"}
	has, err = engine.Desc("id").Get(&user)
	fmt.Println(has)
	fmt.Println(user)

	var name string
	has, err = engine.Table("user").Where("id = ?", 2).Cols("name").Get(&name)
	fmt.Println(has)

	name = "jerry0"
	err = engine.Where("name = ?", name).And("id > 1").Limit(10, 0).Find(&users)
	fmt.Println(users)

	affected, err = engine.Where("name=?", "piter").Delete(&user)

	affected, err = engine.ID(2).Delete(&user)
	fmt.Println(affected)

	//事务
	err = insert(engine)
}

func insert(engine *xorm.Engine) error {
	ctime := time.Now().Unix()
	session := engine.NewSession()
	defer session.Close()
	// add Begin() before any action
	if err := session.Begin(); err != nil {
		// if returned then will rollback automatically
		return err
	}
	user1 := User{Name: "xiaoxiao", Ctime: ctime}
	if _, err := session.Insert(&user1); err != nil {
		fmt.Println(err)
		return err
	}

	user2 := User{Name: "jerry0"}
	if _, err := session.Where("id = ?", 2).Update(&user2); err != nil {
		fmt.Println(err)
		return err
	}
	if _, err := session.Exec("delete from user where name = ?", user2.Name); err != nil {
		fmt.Println(err)
		return err
	}
	// add Commit() after all actions
	return session.Commit()
}
