package main

import "fmt"

func main() {
	var name string = "Muhammad Rizki Pradana"
	fmt.Println(name)
	photo := "https://myblog.com/photo.png"
	fmt.Println(photo)
	var email string = "mrpradana2@gmail.com"
	fmt.Println(email)
	var age int8 = 22
	fmt.Println(age)
	var telp string = "085952805111"
	fmt.Println(telp)
	var isMarried bool = false 
	fmt.Println(isMarried) 
	school := map[string]string{
		"name": "Universitas Pancasakti Tegal",
		"major": "manajemen",
	}
	fmt.Println(school)
}