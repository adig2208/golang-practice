package main

import "fmt"

type product struct {
	title string
	id    string
	price float64
}

func main() {
	//1
	hobbies := []string{"sports", "movies", "science"}
	fmt.Println(hobbies)

	//2
	fmt.Println(hobbies[0])
	fmt.Println(hobbies[1:])

	//3
	hobby_slice := hobbies[:2]
	hobby_slice_2 := []string{"sports", "movies"}
	fmt.Println(hobby_slice, hobby_slice_2)

	//4
	hobby_reslice := hobby_slice[1:2]
	hobby_reslice = append(hobby_reslice, "science")
	fmt.Println(hobby_reslice)

	//5,6
	goals := []string{"Learn golang", "be proficient"}
	goals[1] = "get better"
	fmt.Println(goals)
	goals = append(goals, "master this language")
	fmt.Println(goals)

	//7
	var p1 product = product{
		title: "headphones",
		id:    "a101",
		price: 34.50}

	var p2 product = product{
		title: "headphones2",
		id:    "a102",
		price: 37.50}

	prod_list := []product{p1, p2}
	fmt.Println(prod_list)

	var p3 product = product{
		title: "headphones3",
		id:    "a103",
		price: 39.50}

	prod_list = append(prod_list, p3)
	fmt.Println(prod_list)

	abc := []float64{1.5, 2.5, 3.5}
	def := []float64{4.5, 5.5, 6.5}
	abc = append(abc, def...)
	fmt.Println(abc)
}
