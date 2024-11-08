// specifies the data types needed for the application
package types

import "fmt"

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Status      string `json:"status"`
	CreatedAt   string `json:"createdAt"`
	UpdatedAt   string `json:"updatedAt"`
}

type Tasks []Task

func (t Tasks) Output() {
	var task Task

	if len(t) == 1 {
		task = t[0]
		fmt.Printf("Task: %v, added successfully (ID: %v)\n", task.Description, task.Id)
	} else {
		for _, task = range t {
			fmt.Printf("Task: %v, added successfully (ID: %v)\n", task.Description, task.Id)
		}
	}
}

const JsonFile string = "tasks.json"
