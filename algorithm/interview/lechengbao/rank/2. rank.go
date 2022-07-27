package main

import (
	"fmt"
)

//基于以下几点要点分析
//a. 月内活动中
//1. 用户数量较大
//2. 月内更新频繁
//考虑使用hash来实现对user的快速索引
//
//b. 活动结束后
//3. 需要有序
//4. 需要前后10名玩家
//考虑使用线性结构存储，实现前后的快速索引，比如list

var (
	userList *UList		// 保存User地址
	userMap = make(map[string]*User, 10)
)
type UList struct {
	userAddr *User
	next *UList
}
type User struct {
	userID string
	score	int
	updateTime	int
}
func (u *User) UpdateSocre(delta int) bool{
	return  false
}

// 活动结束后
func sortUser(){
	// userList sort： 自定义排序算法

}


func getRank(userId string) []User{
	users := make([]User, 0)
	u := userMap[userId]
	users = append(users, *u)

	// for v = u.previow
	//   append( , *v)
	// for v = u.next
	//   append( , *v)

	return users
}

func findUserAndUpdate(userId string, delta int) *User{
	u := userMap[userId]
	u.UpdateSocre(delta)
	return nil
}






// Init --------------------------------- Test
func Init() {
	users := []User{
		{userID: "0002", score: 10, updateTime: 10001},
		{userID: "0003", score: 11, updateTime: 10000},
		{userID: "0001", score: 10, updateTime: 10000},
	}
	for _, u := range users{
		nu := new(User)
		nu.score = u.score
		nu.updateTime = u.updateTime
		nu.userID = u.userID
		userMap[nu.userID] = nu
	}
}
func print(){
	fmt.Println(userMap)
	for _, v := range userMap{
		fmt.Println(*v)
	}
}

func InitList() *UList{
	p := new(UList)
	head := p
	for _, v := range userMap{
		node := new(UList)
		node.userAddr = v
		fmt.Println("node", node)
		p.next = node
		p = p.next
	}
	//fmt.Println("p:", p)
	return head.next
}
func printList(){
	for p := userList; p != nil; p = p.next{
		fmt.Printf("%p %v\n", p.userAddr, *p.userAddr)
	}
}

func main() {
	Init()
	print()
	userList = InitList()
	printList()
}