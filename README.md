Para ejecutar el programa se deben tener en cuenta lo siguiente:
el programa se ejecuta con el comando go run . seguido de uno de los siguientes comandos para usar en el task CLI:

// # Agregar 1 tarea
// task-cli add "Buy groceries"
// # Output: Task added successfully (ID: 1)

// # Actualizar 1 tarea
// task-cli update 1 "Buy groceries and cook dinner"
// task-cli delete 1

// # marcar una tarea como terminada o en progreso
// task-cli mark-in-progress 1
// task-cli mark-done 1

// # listar todas las tareas
// task-cli list

// # listar tareas por estado
// task-cli list done
// task-cli list todo
// task-cli list in-progress
