package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"time"
)

// Propiedades de Tarea
// Cada tarea debe tener las siguientes propiedades:

// id: Un identificador único para la tarea
// description: Una breve descripción de la tarea
// status: El estado de la tarea (todo, in-progress, done)
// createdAt: La fecha y hora en que se creó la tarea
// updatedAt: La fecha y hora en que se actualizó la tarea por última vez

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

//Variables necesarias para el manejo de las task:

var tasks []Task
var fileName = "tasks.json"
var i int

// # Agregar 1 tarea
// task-cli add "Buy groceries"
// # Output: Task added successfully (ID: 1)

// # Updating and deleting tasks
// task-cli update 1 "Buy groceries and cook dinner"
// task-cli delete 1

// # Marking a task as in progress or done
// task-cli mark-in-progress 1
// task-cli mark-done 1

// # Listing all tasks
// task-cli list

// # Listing tasks by status
// task-cli list done
// task-cli list todo
// task-cli list in-progress

func loadTask() error {
	//:= es un operador que asigna los valores que se retornen de la derecha, a las variables de la izquierda, en el orden que se devuelva,
	//Es por esto que se debe tener mucho cuidado en el orden que se hacen.
	//ioUtil.ReadFile devuelve primero el contenido del archivo si lo hay, y segundo error si hay algún error.
	file, err := ioutil.ReadFile(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			//Si el archivo no existe se debe crear y vacio.
			tasks = []Task{}
			return saveTasks()
		}
		return err
	}
	err = json.Unmarshal(file, &tasks)
	return err
}

func saveTasks() error {
	file, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		return err
	}
	return ioutil.WriteFile(fileName, file, 0644)
}

func addTask(description string) {
	id := len(tasks) + 1
	newTask := Task{
		ID:          id,
		Description: description,
		Status:      "todo",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
	tasks = append(tasks, newTask)
	saveTasks()
	fmt.Printf("Tarea añadida correctamente (ID: %d)\n", id)
}

func listTask(filter string) {
	for _, task := range tasks {
		if filter == "" || task.Status == filter {
			fmt.Printf("%d: %s (Status: %s)\n", task.ID, task.Description, task.Status)
		}
	}
}

func updateTask(id int, description string) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Description = description
			tasks[i].UpdatedAt = time.Now()
			saveTasks()
			fmt.Println("Tarea Actualizada correctamente")
			return
		}
	}
	fmt.Println("Tarea no encontrada")
}

func deleteTask(id int) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			fmt.Println("Tarea Borrada correctamente")
			return
		}
	}
	fmt.Println("Tarea no encontrada")
}

func markTask(id int, status string) {
	for i := range tasks {
		if tasks[i].ID == id {
			tasks[i].Status = status
			tasks[i].UpdatedAt = time.Now()
			saveTasks()
			fmt.Printf("Tarea marcada como %s\n", status)
			return
		}
	}
	fmt.Println("Tarea no encontrada")
}

func main() {

	//Validamos rápidamente si existe el archivo json para almacenar los tareas ya creado, para crearlo si no existe.
	loadTask()

	//Estamos cuidando de que los argumentos que pasa si cumplan con el minimo requerido para ejecutar algún comando.
	if len(os.Args) < 2 {
		fmt.Println("Used argumentos correctos, por ejemplo: task-cli <command> [arguments]")
		return
	}

	err := loadTask()
	if err != nil {
		fmt.Printf("Error cargando tareas: %v\n", err)
		return
	}

	//Ahora, como ya recibimos el comando por argumentos posicionales, vamos a empezar tomar todos los casos posibles
	switch os.Args[2] {
	case "add":
		if len(os.Args) != 4 {
			fmt.Printf("Para agregar tareas el comando correcto es: task-cli add <description>")
			return
		}
		addTask(os.Args[2])
	case "list":
		if len(os.Args) > 3 {
			listTask(os.Args[3])
		} else {
			listTask("")
		}

	case "update":
		if len(os.Args) != 4 {
			fmt.Println("Para actualizar tareas el comando correcto es: task-cli update <description>")
			return
		}
		id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("ID invalido")
			return
		}
		updateTask(id, os.Args[4])
	case "delete":
		if len(os.Args) != 4 {
			fmt.Println("Usage: task-cli delete <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("ID invalido")
			return
		}
		deleteTask(id)
	case "mark-in-progress":
		if len(os.Args) != 4 {
			fmt.Println("Usage: task-cli mark-in-progress <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("ID invalido")
			return
		}
		markTask(id, "in-progress")
	case "mark-done":
		if len(os.Args) != 4 {
			fmt.Println("Usage: task-cli mark-done <id>")
			return
		}
		id, err := strconv.Atoi(os.Args[3])
		if err != nil {
			fmt.Println("ID invalido")
			return
		}
		markTask(id, "done")
	default:
		fmt.Println("Comando desconocido:", os.Args[1])
	}

}
