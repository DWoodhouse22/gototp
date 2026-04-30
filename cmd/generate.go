package cmd

import (
	"fmt"

	"github.com/DWoodhouse22/gototp/storage"
	"github.com/DWoodhouse22/gototp/totp"
)

func Generate() {
	secret, err := storage.LoadSecret()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	code, err := totp.Generate(secret)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println(code)
}
