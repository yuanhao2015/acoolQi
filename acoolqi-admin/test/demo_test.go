/*
@Time : 2021-10-15 15:54
@Author : acool
@File : demo_test
*/
package test

import (
	"fmt"
	"github.com/satori/go.uuid"
	"strings"
	"testing"
)

func TestA(t *testing.T) {
	//u1 := uuid.Must(uuid.NewV4())
	//fmt.Printf("UUIDv4: %s\n", u1)
	must := uuid.Must(uuid.NewV4(), nil)
	fmt.Println(must)
	// or error handling
	u2 := uuid.NewV4()
	fmt.Printf("%T", u2.String())

	fmt.Printf("UUIDv4: %s\n", u2)
	fmt.Println(strings.ReplaceAll(u2.String(), "-", ""))

	// Parsing UUID from string input
	u2, err := uuid.FromString("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
	if err != nil {
		fmt.Printf("Something went wrong: %s", err)
		return
	}
	fmt.Printf("Successfully parsed: %s", u2)
	a := "users-admin-83cf67c5a7bb442bb0cc14681da10005"
	fmt.Println(strings.Split(a, "-")[2])
}
