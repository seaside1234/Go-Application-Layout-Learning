package ownCommandPrint
import (
	"flag"
	"encoding/json"
	"fmt"
	"seaside1234/work/internal/Innerpackage2"
)
var (
         name    string
         url     string
         email   string
         company string
     )
     
     func PrintCommandFlag() {
     
         flag.StringVar(&name, "n", "zzz", "show user name")
         flag.StringVar(&url, "u", "sea.work", "show user url")
         flag.StringVar(&email, "e", "102@qq.com", "show user email")
         flag.StringVar(&company, "c", "gongsihao", "show user company")
         flag.Parse()
     
         var oneUser userinformantionFlag.FlagUserInformation
         oneUser.Name = name
         oneUser.Url = url
         oneUser.Email = email
         oneUser.Company = company
         jsonByte, _ := json.MarshalIndent(oneUser, " ", " ")
        fmt.Println(string(jsonByte))
        
}