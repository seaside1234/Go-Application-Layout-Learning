package CmdOs

import (
	"os"
	"seaside1234/work/internal/Innerpackage1"
	"github.com/sirupsen/logrus"
)

func PrintCmdOs() userInfoOs.UserInfoOs {
	args := os.Args
	if len(args) != 4 {
		logrus.Println("you need add applicationName,name,email,company field")
	}
	var oneUserInfo userInfoOs.UserInfoOs
	oneUserInfo.ApplicationName = os.Args[0]
	oneUserInfo.Name = os.Args[1]
	oneUserInfo.Email = os.Args[2]
	oneUserInfo.Company = os.Args[3]
	return oneUserInfo
}
