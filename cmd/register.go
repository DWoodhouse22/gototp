package cmd

import (
	"fmt"

	"github.com/DWoodhouse22/gototp/storage"
)

func Register(secret string) {
	err := storage.SaveSecret(secret)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("Secret saved")
}
