package queries

import (
	"fmt"
	"strconv"
)

func QueryUpdateNews(id string, title string) string {
	num, err := strconv.Atoi(id)

	if err != nil {
		fmt.Println("Error")

	}

	q := fmt.Sprintf("UPDATE news SET title = '%s' WHERE id = %d", title, num)
	fmt.Println(q)
	fmt.Println(q)

	return q
}

func QueryListNews() string {
	q := "SELECT * FROM news"
	return q
}
