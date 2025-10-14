package main

import (
	"fmt"
	"os"

	"golang.org/x/crypto/bcrypt"
)

func main() {
	passwords := map[string]string{
		"admin":   "admin123",
		"user1":   "user123",
		"user2":   "user123",
		"manager": "manager123",
	}

	fmt.Println("生成 bcrypt 密码哈希：")
	fmt.Println("====================")

	for username, password := range passwords {
		hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
		if err != nil {
			fmt.Printf("生成 %s 的哈希失败: %v\n", username, err)
			os.Exit(1)
		}
		fmt.Printf("%s (密码: %s):\n%s\n\n", username, password, string(hash))
	}
}
