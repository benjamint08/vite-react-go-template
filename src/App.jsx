import { useState, useEffect } from 'react'
import {clearTodos, addTodo, deleteTodo} from "./functions/todos.js";

function App() {
    const [todos, setTodos] = useState(null)
    const [newTodo, setNewTodo] = useState('')

    useEffect(() => {
        const fetchData = async () => {
            const response = await fetch('/api/todos')
            const data = await response.json()
            setTodos(data)
        }
        fetchData();
    }, []);

  return (
    <>
      <div className={"mt-10"}>
          <h1 className={"text-2xl font-bold text-center"}>React, Vite, Tailwind and Go Stack</h1>
          {todos && todos.length > 0 && (
              <div className={"flex mt-5 justify-center"}>
                  <button className={"bg-red-500 text-white p-2"} onClick={(e) => clearTodos(setTodos)}>Clear Todos</button>
              </div>
          )}
          <div className={"flex mt-5 justify-center flex-col w-[50%] mx-auto"}>
              {todos && (
                  <ul>
                      {todos.map((todo, index) => (
                        <li key={index} className={"border border-gray-300 p-2 ml-2 mt-2"}>{todo}
                            <span className={"text-red-500 ml-2 float-right font-bold hover:cursor-pointer"} key={"del-" + index} onClick={() => deleteTodo(index, todos, setTodos)}>X</span>
                        </li>
                    ))}
                  </ul>
              )}
              {todos && (
                  <>
                      <input type="text" placeholder="Add a todo" className={"border border-gray-300 p-2 ml-2 mt-2 text-black"} value={newTodo} onChange={(e) => setNewTodo(e.target.value)} />
                      <button className={"bg-blue-500 text-white p-2 ml-2 mt-2"} onClick={(e) => addTodo(todos, newTodo, setTodos, setNewTodo)}>Add</button>
                  </>
              )}
          </div>
      </div>
    </>
  )
}

export default App
