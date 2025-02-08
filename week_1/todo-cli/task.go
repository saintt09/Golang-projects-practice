/* Estructuras de datos
- Una estructura Task con los campos:
		ID (int).
		Description (string).
		Completed (bool)
- Una estructura TaskList que contenga un slice de Task.

Las dos estructuras deben tener tags json para que puedan ser serializadas y deserializadas correctamente.
*/

package main

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

/*
{
	"tasks": [
			{
					"id": 1,
					"description": "Estudiando Go",
					"completed": true
			}
	]
}
*/
