package CmdOs
import (
	"encoding/json"
	"seaside1234/work/internal/Innerpackage1"
)
 
func TurningUserInformation(userInformation userInfoOs.UserInfoOs)([]byte){
	jsonByte, _ := json.MarshalIndent(userInformation, "", "")
	return jsonByte
}