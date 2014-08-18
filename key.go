package main

import (
    "fmt"
    "flag"
    "math/rand"
    "strconv"
    "time"
    "strings"
)

var chars string = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"
var secretnumber int = 1200

func main() {
    key := flag.String("key","aaaa-aaaa-aaaa-aaaa", "Sets the key")
    flag.Parse()
    rand.Seed(time.Now().Unix())
    if *key == "aaaa-aaaa-aaaa-aaaa" {
        genKey()
    } else {
        if isKey(*key) {
            fmt.Println("This is a key!")
        } else {
            fmt.Println("This isn't a key!")
        }
    }
}

func genKey() {
    var keyLength int
    for keyLength != secretnumber {
       keyLength = generateKey()
    }
}

func generateKey() int {
    var count int
    var i int = 0
    var newstring = [4]string{}
    for split1 := range newstring {
        for i < 4 {
            tempcount := chars[rand.Intn(len(chars))]
            count += int(tempcount)
            tempstring := string(tempcount)
            newstring[split1] += tempstring
            i++
        }
        i = 0
        fmt.Printf(newstring[split1])
        if split1 != 3 {
            fmt.Printf("-")
        }
    }
    fmt.Printf(" = " + strconv.Itoa(count) + "\n")
    return count
}

func isKey(key string) bool {
    var secret int
    split1 := strings.Split(key,"-")
    for _, split2 := range split1 {
        for _, split3 := range split2 {
            secret += int(split3)
        }
    }
    fmt.Println(secret)
    if secret == secretnumber {
        return true
    }
    return false
}