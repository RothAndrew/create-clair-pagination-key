package main

import (
	"fmt"
	"github.com/coreos/clair/pkg/pagination"
)

func main() {
	fmt.Println(pagination.Must(pagination.NewKey()).String())
}
