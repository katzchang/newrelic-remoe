package main

import (
	"fmt"
	"github.com/reeve0930/go-remoe"
	"os"
)

func main() {
	token := os.Getenv("NATURE_API_TOKEN")
	client := remoe.NewClient(token)
	data, _ := client.GetRawData()

	fmt.Println(data)
}
