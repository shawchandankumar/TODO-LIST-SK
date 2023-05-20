// stores all the task created previously
// when the application starts the first 
// made to api to get all the task
const AllTasks = [];

// console.log(getInitialTasks())

async function getInitialTasks () {
    const taskPr = await fetch('http://localhost:8080/tasks');
    const taskJson = await taskPr.json();
    return taskJson
}

let showAddTaskFormBtn = document.getElementById('show-add-task-form-btn');
let createTaskContainer = document.getElementById('create-task-container');

toggleDisplay(showAddTaskFormBtn, createTaskContainer)


$("#create-task-container").on("submit", function(e) {
    e.preventDefault()

    let form = $(this);
    let formData = extractFormData(form)
    form[0].reset()
    
    $.ajax({
        type: "POST",
        url: form.attr('action'),
        data: JSON.stringify(formData),
        success: (data) => {
            attachEventListenerToNewTodoTask(data)
        }
    })
})

function attachEventListenerToNewTodoTask (data) {
    let newTaskDom = createNewTodoTask(data)
    todoListContainer.prepend(newTaskDom)

    // attaching the event listener on a newTaskDom to toggle the button for hiding and showing the edit todo content form
    toggleDisplay(document.querySelector(`#id${data.ID} .edit-btn`), document.querySelector(`#id${data.ID} .Edit-todo-content`))

    // attaching an event listener on a delete btn
    // to delete the dom and send a ajax request to the server
    deleteTodoTask(newTaskDom, document.querySelector(`#id${data.ID} .delete-btn`), data.ID)

    editTodoTask(newTaskDom, data.ID)
}


function extractFormData(form) {
    let formData = {}
    form.serializeArray().forEach(data => {
        console.log(data)
        formData[data.name] = data.value
    })

    formData["priority"] = parseInt(formData["priority"])
    formData["userId"] = 1

    return formData
}


let todoListContainer = document.getElementById('todo-list-container')


getInitialTasks().then(task => {
    task.forEach(data => {
        attachEventListenerToNewTodoTask(data)
    })
})

function editTodoTask(newTodoTaskDom, todoTaskId) {
    $(`#${newTodoTaskDom.id} .Edit-todo-content`).on('submit', function (e) {
        e.preventDefault()

        let form = $(this);
        let formData = extractFormData(form)
        
        $.ajax({
            type: "PUT",
            url: `http://localhost:8080/tasks/${todoTaskId.toString()}`,
            data: JSON.stringify(formData),
            success: (data) => {
                console.log(data)
                let modifiedTodoContent = createNewTodoTaskDom(data)
                let modifiedTodoForm = createNewTodoTaskEditForm(data)

                // destroy the old todo content dom elements
                $(`#${newTodoTaskDom.id} .Edit-todo-content`).remove()
                $(`#${newTodoTaskDom.id} .todo-content`).remove()

                // append modified content and edit form
                newTodoTaskDom.appendChild(modifiedTodoContent)
                newTodoTaskDom.appendChild(modifiedTodoForm)

                toggleDisplay(document.querySelector(`#id${data.ID} .edit-btn`), document.querySelector(`#id${data.ID} .Edit-todo-content`))

                editTodoTask(newTodoTaskDom, data.ID)
            }
        })
    })
}

function toggleDisplay (element, elementToToggle) {
    element.addEventListener('click', () => {
        if (elementToToggle.style.display == 'none') {
            elementToToggle.style.display = 'block'
        }
        else elementToToggle.style.display = 'none'
    })
}

function deleteTodoTask(todoTaskElement, deleteBtn, todoTaskId) {
    deleteBtn.addEventListener('click', () => {
        $.ajax({
            type: "DELETE",
            url: `http://localhost:8080/tasks/${todoTaskId.toString()}`,
            success: () => {
                console.log('Successfully deleted')
            }
        })
        todoTaskElement.remove()
    })
}

function createNewTodoTask (task) {
    let newTodoTask = document.createElement('div')
    newTodoTask.classList.add('todo-task')
    newTodoTask.setAttribute('id', `id${task.ID}`)
    newTodoTask.appendChild(createActionBtn(`<i class="fa-solid fa-xmark"></i>`, `delete-btn`))
    newTodoTask.appendChild(createActionBtn(`<i class="fa-solid fa-pen"></i>`, `edit-btn`))
    newTodoTask.appendChild(createNewTodoTaskDom(task))
    newTodoTask.appendChild(createNewTodoTaskEditForm(task))
    AllTasks.push(task)
    return newTodoTask
}

function createActionBtn(itag, idVal) {
    let newBtn = document.createElement('button')
    newBtn.classList.add(`icon-btn`)
    newBtn.classList.add(`${idVal}`)
    newBtn.innerHTML = itag
    return newBtn
}

function createNewTodoTaskDom (task) {
    let newTaskEl = document.createElement('div')
    newTaskEl.classList.add('todo-content')
    
    newTaskEl.innerHTML = `<span class="todo-priority">Priority ${task.priority}</span>
    <h3 class="todo-title-content">${task.title}</h3>
    <span class="todo-task-content">${task.todo}</span>`

    return newTaskEl
}

function createNewTodoTaskEditForm (task) {
    let newTaskEditForm = document.createElement('form')
    newTaskEditForm.action = 'http://localhost:8080/tasks'
    newTaskEditForm.classList.add('Edit-todo-content')

    newTaskEditForm.innerHTML = `<label for="title">Title</label>
    <input type="text" name="title" value="${task.title}" >

    <label for="todo">Todo Task</label>
    <input type="text" name="todo" value="${task.todo}" >

    <label for="priority">Priority</label>
    <input type="text" name="priority" value="${task.priority}" >

    <input type="submit" class="btn" id="edit-btn">`

    return newTaskEditForm
}