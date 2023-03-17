package entity

import "time"

type User struct {
	Id   string
	Name string
	Age  int32
	Time time.Time
}

func NewUser() *User {
	return &User{}
}

func (u *User) JavaClassName() string {
	return "com.example.dubboconsumer.User" // 如果与 Java 互通，需要与 Java 侧 User class全名对应,
}
