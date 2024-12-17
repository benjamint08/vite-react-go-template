async function deleteTodo(index, todos, setTodos) {
    const response = await fetch('/api/delete-todo', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ todo: todos[index].todo })
    });
    if(response.ok) {
        const oneWeDontWant = todos[index].todo
        setTodos(todos.filter(todo => todo.todo !== oneWeDontWant))
    }
}

async function addTodo(todos, newTodo, setTodos, setNewTodo) {
    const response = await fetch('/api/add-todo', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify({ todo: newTodo })
    });
    if(response.ok) {
        setTodos([...todos, {
            todo: newTodo
        }])
        setNewTodo('')
    } else {
        const res = await response.text()
        window.alert('Failed to add todo - ' + res)
        window.location.reload()
    }
}

async function clearTodos(setTodos) {
    const response = await fetch('/api/clear-todos', {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        }
    });
    if(response.ok) {
        setTodos([])
    } else {
        const res = await response.text()
        window.alert('Failed to clear todos - ' + res)
        window.location.reload()
    }
}

export { clearTodos, addTodo, deleteTodo }