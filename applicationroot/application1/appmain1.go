package main

import (
	"seaside1234/work/packageinapplication1"
	"seaside1234/work/internal/Innerpackage1"
)

func main() {
	var userInformation userInfoOs.UserInfoOs = CmdOs.PrintCmdOs()
	CmdOs.TurningUserInfomation(userInformation)
}

