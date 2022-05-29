package main

import (
	"context"

	"github.com/Zaba505/tblconv/cmd/tblconv-plugin-sqlite/cmd"
)

func main() {
	cmd.Execute(context.Background())
}
