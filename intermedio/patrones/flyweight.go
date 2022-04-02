package main

import (
	"fmt"
	"strings"
)

type User struct {
	FullName string
}

func NewUser(fullName string) *User {
	return &User{FullName: fullName}
}

var allNames []string

type User2 struct {
	names []uint8
}

// con esto reducimos el tamaño de memoria usado para almacenar un usuario
// pero necesitamos una funcion que transforme el uint8 a string again
func NewUser2(fullName string) *User2 {
	getOrAdd := func(s string) uint8 {
		for i := range allNames {
			if allNames[i] == s {
				return uint8(i)
			}
		}
		allNames = append(allNames, s)
		return uint8(len(allNames) - 1)
	}
	result := User2{}
	parts := strings.Split(fullName, " ")
	for _, p := range parts {
		result.names = append(result.names, getOrAdd(p))
	}
	return &result
}

func (u *User2) FullName() string {
	var parts []string
	for _, id := range u.names {
		parts = append(parts, allNames[id])
	}
	return strings.Join(parts, " ")
}

func main() {
	paco := NewUser("Paco Nogales")
	maria := NewUser("Maria Morales")
	fmt.Println(paco.FullName)
	fmt.Println(maria.FullName)

	paco2 := NewUser2("Paco nogales2")
	fmt.Println(paco2.FullName())
}
