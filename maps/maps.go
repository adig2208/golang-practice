package main

import "fmt"

type floatMap map[string]float64

func main() {
	websites := map[string]string{
		"Google": "http://www.google.com",
		"AWS":    "http://aws.amazon.com",
	}
	fmt.Println(websites)
	fmt.Println(websites["Google"])

	websites["Linkedin"] = "http://www.linkedin.com"
	fmt.Println(websites)

	delete(websites, "Linkedin")
	fmt.Println(websites)

	userNames := make([]string, 2)
	userNames[0] = "Julie"
	userNames = append(userNames, "Max")
	fmt.Println(userNames)

	courseRating := make(map[string]float64, 3) //can use aliases,  courseRating := make(floatMap, 3)
	courseRating["Go"] = 5.0
	courseRating["React"] = 3.50
	courseRating["Angular"] = 4.76
	fmt.Println(courseRating)

	for i, j := range userNames {
		fmt.Println(i)
		fmt.Println(j)
	}

	for key, val := range courseRating {
		fmt.Println(key)
		fmt.Println(val)
	}

}
