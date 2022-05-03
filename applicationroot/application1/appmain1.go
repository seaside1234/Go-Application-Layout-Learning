package main

import (
	"seaside1234/work/packageinapplication1"
	"seaside1234/work/internal/Innerpackage1"
	"seaside1234/work/packageinapplication1/packageinapplication1module"
)

func main() {
	var userInformation userInfoOs.UserInfoOs = CmdOs.PrintCmdOs()
	var userInformationByte []byte = CmdOs.TurningUserInformation(userInformation)
	CmdPrint.PrintUserInformation(userInformationByte)
}

