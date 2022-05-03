package CmdOs
import (
	"fmt"
	"encoding/json"
	"seaside1234/work/internal/Innerpackage1"
)
 
func TurningUserInfomation(userInformation userInfoOs.UserInfoOs){
	jsonByte, _ := json.MarshalIndent(userInformation, "", "")
	fmt.Println(string(jsonByte))
}