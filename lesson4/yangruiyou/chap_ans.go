package main

import (
	"fmt"
)

func main() {

	s := []int{9, 1, 9, 5, 4, 4, 2, 1, 5, 4, 8, 8, 4, 3, 6, 9, 5, 7, 5}
	fmt.Println(UniqueInts(s))

	irregularMatrix := [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11}, {12, 13, 14, 15}, {16, 17, 18, 19, 20}}
	slice := Flatten(irregularMatrix)
	fmt.Printf("1x%d : %v \n", len(slice), slice)

	fmt.Printf(" 3x%d: %v\n", neededRows(slice, 3), Make2D(slice, 3))
	fmt.Printf(" 4x%d: %v\n", neededRows(slice, 4), Make2D(slice, 4))
	fmt.Printf(" 5x%d: %v\n", neededRows(slice, 5), Make2D(slice, 5))
	fmt.Printf(" 6x%d: %v\n", neededRows(slice, 6), Make2D(slice, 6))

	iniData := []string{
		"; Cut down copy of Mozilla application.ini file",
		"",
		"[App]",
		"Vendor=Mozilla",
		"Name=Iceweasel",
		"Profile=mozilla/firefox",
		"Version=3.5.16",
		"[Gecko]",
		"MinVersion=1.9.1",
		"MaxVersion=1.9.1.*",
		"[XRE]",
		"EnableProfileMigrator=0",
		"EnableExtensionManager=1",
	}
	ini := ParseIni(iniData)
	PrintIni(ini)

}

//创建一个函数以接受一个[]int切片并返回一个[]int切片，其中返回的切片为传入切片的副本，只是将其中重复的内容删除了。
func UniqueInts(s []int) []int {

	seen := map[int]bool{}
	unique := []int{}
	for _, x := range s {
		if _, ok := seen[x]; !ok {
			unique = append(unique, x)
			seen[x] = true
		}

	}
	return unique

}

//创建一个函数接受一个[][]int切片,然后返回一个[]int切片，其中包含二维切片的第一个切片，接着是二维切片中的第二个切片等
func Flatten(s [][]int) []int {
	slice := []int{}
	for _, interSlice := range s {
		for _, x := range interSlice {
			slice = append(slice, x)
		}
	}
	return slice
}

//创建一个接受[]int切片和一个列数量参数的函数，然后返回一个[][]int切片，其内部切片的长度与给定的列数量参数相同。

func neededRows(slice []int, columns int) int {
	rows := len(slice) / columns
	if len(slice)%columns != 0 {
		rows++
	}
	return rows
}

func Make2D(s []int, columns int) [][]int {
	matrix := make([][]int, neededRows(s, columns))
	for i, x := range s {

		row := i / columns
		column := i % columns
		if matrix[row] == nil {
			matrix[row] = make([]int, columns)

		}
		matrix[row][column] = x

	}
	return matrix

}

//创建一个接受[]string切片参数的函数，其中该切片包含一个.ini文件格式的内容，并返回一个map[string]map[string]类型，其键为"键-值"映射组成的映射组。
//空行与以;开头的行需忽略。每一组在其所在的行中以一个方括号保卫的名字标识，同时每一组的键和值以一行或者更多行的"键-值"的形式给出。
func ParseIni() {}

func PrintIni() {}
