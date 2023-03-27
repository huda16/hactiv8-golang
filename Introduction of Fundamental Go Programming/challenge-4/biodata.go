package main

import (
	"fmt"
	"os"
	"strconv"
)

type Friend struct {
	Name     string
	Address  string
	Job      string
	ClassNum int
	Reason   string
}

func main() {
	friends := []Friend{
		{"Mochamad Nurul Huda", "Jl. Laswi", "Software Engineer", 1, "Ingin belajar bahasa Go untuk memperluas skill programming"},
		{"Irfansyah", "Jl. Merdeka", "Data Analyst", 2, "Tertarik dengan kemampuan bahasa Go dalam menangani concurrency"},
		{"Andreas", "Jl. Buah Batu", "DevOps Engineer", 3, "Menggunakan bahasa Go untuk membangun sistem infrastruktur di perusahaannya"},
	}

	if len(os.Args) < 2 {
		fmt.Println("Usage: go run biodata.go [class number]")
		os.Exit(1)
	}
	
	classNum := os.Args[1]
	num, err := strconv.Atoi(classNum)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	if num > len(friends) {
		fmt.Println("There's no data")
	}
	
	for _, friend := range friends {
		if fmt.Sprintf("%d", friend.ClassNum) == classNum {
			fmt.Printf("Nama      : %s\n", friend.Name)
			fmt.Printf("Alamat    : %s\n", friend.Address)
			fmt.Printf("Pekerjaan : %s\n", friend.Job)
			fmt.Printf("Alasan    : %s\n", friend.Reason)
		}
	}
}