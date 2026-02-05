import { useState, useEffect } from 'react'
import {clearTodos, addTodo, deleteTodo} from "./functions/todos.js";

function App() {
    const [todos, setTodos] = useState(null)
    const [newTodo, setNewTodo] = useState('')
    const [lastReqTime, setLastReqTime] = useState("")
    const [files, setFiles] = useState([])
    const [selectedFile, setSelectedFile] = useState(null)
    const [uploadStatus, setUploadStatus] = useState("")

    function updateLastReqTime(reqTime, resTime) {
        let timeTookInMs = resTime.getTime() - reqTime.getTime();
        setLastReqTime(`${timeTookInMs} ms`)
    }

    async function fetchTodos() {
        const nowTime = new Date();
        const response = await fetch('/api/todos')
        const endTime = new Date();
        updateLastReqTime(nowTime, endTime)
        const data = await response.json()
        setTodos(data)
    }

    async function fetchFiles() {
        const nowTime = new Date();
        const response = await fetch('/api/files')
        const endTime = new Date();
        updateLastReqTime(nowTime, endTime)
        if (response.ok) {
            const data = await response.json()
            setFiles(data)
        }
    }

    async function handleUpload() {
        if (!selectedFile) {
            setUploadStatus("Please select a file to upload.")
            return
        }
        const formData = new FormData()
        formData.append('file', selectedFile)
        const nowTime = new Date();
        const response = await fetch('/api/upload', {
            method: 'POST',
            body: formData
        })
        const endTime = new Date();
        updateLastReqTime(nowTime, endTime)
        if (response.ok) {
            setUploadStatus("Upload successful.")
            setSelectedFile(null)
            await fetchFiles()
        } else {
            const res = await response.text()
            setUploadStatus(`Upload failed: ${res}`)
        }
    }

    useEffect(() => {
        fetchTodos();
        fetchFiles();
    }, []);

  return (
    <>
      <div className={"mt-10"}>
          <h1 className={"text-2xl font-bold text-center"}>React, Vite, Tailwind and Go Stack</h1>
          {todos && todos.length > 0 && (
              <div className={"flex mt-5 justify-center"}>
                  <button className={"bg-red-500 text-white p-2"} onClick={(e) => {
                      const nowTime = new Date();
                      clearTodos(setTodos)
                      const endTime = new Date();
                      updateLastReqTime(nowTime, endTime)
                  }}>Clear Todos</button>
              </div>
          )}
          <div className={"flex mt-5 justify-center flex-col w-[50%] mx-auto"}>
              <h2 className={"text-xl font-bold"}>Todos</h2>
              {todos && (
                  <ul>
                      {todos.map((todo, index) => (
                        <li key={index} className={"border border-gray-300 p-2 ml-2 mt-2"}>{todo}
                            <span className={"text-red-500 ml-2 float-right font-bold hover:cursor-pointer"} key={"del-" + index} onClick={() => {
                                const nowTime = new Date();
                                deleteTodo(index, todos, setTodos)
                                const endTime = new Date();
                                updateLastReqTime(nowTime, endTime)
                            }}>X</span>
                        </li>
                    ))}
                  </ul>
              )}
              {todos && (
                  <>
                      <input type="text" placeholder="Add a todo" className={"border border-gray-300 p-2 ml-2 mt-2 text-black"} value={newTodo} onChange={(e) => setNewTodo(e.target.value)} />
                      <button className={"bg-blue-500 text-white p-2 ml-2 mt-2"} onClick={(e) => {
                          const nowTime = new Date();
                          addTodo(todos, newTodo, setTodos, setNewTodo)
                          const endTime = new Date();
                          updateLastReqTime(nowTime, endTime)
                      }}>Add</button>
                  </>
              )}
          </div>
          <div className={"flex mt-8 justify-center flex-col w-[50%] mx-auto"}>
              <h2 className={"text-xl font-bold"}>Files</h2>
              <div className={"flex mt-2"}>
                  <input
                      type="file"
                      className={"border border-gray-300 p-2 text-white w-full"}
                      onChange={(e) => setSelectedFile(e.target.files?.[0] || null)}
                  />
                  <button className={"bg-green-600 text-white p-2 ml-2"} onClick={handleUpload}>
                      Upload
                  </button>
              </div>
              {uploadStatus && (
                  <p className={"text-sm text-gray-600 mt-2"}>{uploadStatus}</p>
              )}
              <ul className={"mt-3"}>
                  {files.map((file, index) => (
                      <li key={index} className={"border border-gray-300 p-2 mt-2 flex items-center justify-between"}>
                          <span>{file}</span>
                          <a
                              className={"bg-blue-600 text-white px-3 py-1"}
                              href={`/api/download?file=${encodeURIComponent(file)}`}
                          >
                              Download
                          </a>
                      </li>
                  ))}
              </ul>
          </div>
          <footer>
                <p className={"text-center mt-10 text-gray-500"}>Last request time: {lastReqTime}</p>
          </footer>
      </div>
    </>
  )
}

export default App
