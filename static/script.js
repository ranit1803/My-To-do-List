const apiUrl = "http://localhost:8080/tasks";
const todoInput = document.querySelector(".todo-input");
const todoButton = document.querySelector(".todo-button");
const todoList = document.querySelector(".todo-list");
const filterOption = document.querySelector(".filter-todo");

document.addEventListener("DOMContentLoaded", fetchTasks);
todoButton.addEventListener("click", addTask);
todoList.addEventListener("click", handleTaskActions);
filterOption.addEventListener("change", filterTasks);

async function fetchTasks() {
    let response = await fetch(apiUrl);
    let tasks = await response.json();
    todoList.innerHTML = "";
    tasks.forEach(task => renderTask(task));
}

async function addTask(event) {
    event.preventDefault();
    let taskTitle = todoInput.value.trim();
    if (taskTitle === "") return;
    
    let response = await fetch(apiUrl, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ title: taskTitle, completed: false })
    });
    let newTask = await response.json();
    renderTask(newTask);
    todoInput.value = "";
}

async function completeTask(id) {
    await fetch(`${apiUrl}/${id}`, { method: "PUT" });
    fetchTasks();
}

async function deleteTask(id) {
    await fetch(`${apiUrl}/${id}`, { method: "DELETE" });
    fetchTasks();
}

function renderTask(task) {
    const todoDiv = document.createElement("div");
    todoDiv.classList.add("todo");
    if (task.completed) todoDiv.classList.add("completed");

    const newTodo = document.createElement("li");
    newTodo.innerText = task.title;
    newTodo.classList.add("todo-item");
    todoDiv.appendChild(newTodo);

    const completeButton = document.createElement("button");
    completeButton.innerHTML = '<i class="fas fa-check-circle"></i>';
    completeButton.classList.add("complete-btn");
    completeButton.setAttribute("data-id", task.id);
    todoDiv.appendChild(completeButton);

    const trashButton = document.createElement("button");
    trashButton.innerHTML = '<i class="fas fa-trash"></i>';
    trashButton.classList.add("trash-btn");
    trashButton.setAttribute("data-id", task.id);
    todoDiv.appendChild(trashButton);

    todoList.appendChild(todoDiv);
}

function handleTaskActions(e) {
    const item = e.target;
    if (item.classList.contains("complete-btn")) {
        const taskId = item.getAttribute("data-id");
        completeTask(taskId);
    } else if (item.classList.contains("trash-btn")) {
        const taskId = item.getAttribute("data-id");
        deleteTask(taskId);
    }
}

function filterTasks() {
    const filterValue = filterOption.value;
    const todos = todoList.childNodes;
    todos.forEach(todo => {
        switch (filterValue) {
            case "all":
                todo.style.display = "flex";
                break;
            case "completed":
                todo.classList.contains("completed") ? 
                    todo.style.display = "flex" : 
                    todo.style.display = "none";
                break;
            case "incomplete":
                !todo.classList.contains("completed") ? 
                    todo.style.display = "flex" : 
                    todo.style.display = "none";
                break;
        }
    });
}
