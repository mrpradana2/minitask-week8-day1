package main

import (
	"fmt"
	"strings"
)

func main() {
    fakeLogin("user123", "User123#")
}

func fakeLogin(username, password string)  {
    user := map[string]string{
		"username": "user123",
		"password": "User123#",
	}

    resultUsername, messageUsername :=  validationUsername(username)
    resultPassword, messagePassword := validationPass(password)

    if resultUsername && resultPassword {
        if username == user["username"] && password == user["password"] {
            fmt.Println("Success: Login Berhasil")
        } else {
            if username != user["username"] {
                fmt.Println("Failed: Username anda salah")
            } else {
                fmt.Println("Failed: Password anda salah")
            }
        }
    } else {
        if !resultUsername && !resultPassword {
            fmt.Println("Failed:", messageUsername, "dan", messagePassword)
        } else if !resultUsername {
            fmt.Println("Failed:", messageUsername)
        } else if !resultPassword {
            fmt.Println("Failed:", messagePassword)
        }
    }

}

func validationPass(password string) (result bool, message string)  {
    if len(password) == 0 {
        return false, "Password harus diisi"
    }
    
    if len(password) <= 7 {
        return false, "Jumlah karakter password harus diatas 7"
    }

    splitPass := strings.Split(password, "")

    if !checkUpperCaseLetters(splitPass) || !checkLowerCaseLetters(splitPass) || !checkNumbers(splitPass) || !checkCharacters(splitPass) {
        return false, "Format password harus mengandung huruf besar, huruf kecil, angka, dan karakter khusus"
    }

    return true, "Format password benar"
}

func checkUpperCaseLetters(pass []string) (result bool) {
    upperCaseLetter := "ABCDEFGHIJKLMNOPQRSTUVWXYZ"

    for _, char := range pass {
        if strings.Contains(upperCaseLetter, char) {
            return true
        }
    }
    return false    
}

func checkLowerCaseLetters(pass []string) (result bool) {
    lowerCaseLetter := "abcdefghijklmnopqrstuvwxyz"

    for _, char := range pass {
        if strings.Contains(lowerCaseLetter, char) {
            return true
        }
    }
    return false    
}

func checkNumbers(pass []string) (result bool) {
    numbers := "1234567890"

    for _, char := range pass {
        if strings.Contains(numbers, char) {
            return true
        }
    }
    return false    
}

func checkCharacters(pass []string) (result bool) {
    characters := "~!@#$%^&*()`_+-={}[]|;':<>?/.,"

    for _, char := range pass {
        if strings.Contains(characters, char) {
            return true
        }
    }
    return false    
}

func validationUsername(username string) (result bool, message string)  {
    if len(username) == 0 {
        return false, "Username harus diisi"
    }

    return true, "Username sudah diisi"
}