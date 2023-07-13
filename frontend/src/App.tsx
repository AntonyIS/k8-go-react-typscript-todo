import React, { useEffect } from 'react';
import AddTodo from './components/AddTodo';
import TodoList from './components/TodoList';

interface  Todo {
  id :string
  title : string
  description:string
  state:string
}

interface TodosListProps {
  todos : Todo[];
}



const App: React.FC = () => {

  const [todos, setTodos] = React.useState<Todo[]>([])
  const [error, setError] = React.useState<string>("")
  const todosURL = "http://127.0.0.1:8080/api/v1/todos/"

  useEffect(() => {
    const fetchData = async () => {
    try {
        const response = await fetch(todosURL);
        const data = await response.json();
        setTodos(data);
    } catch (error) {
      setError("Backend internal server error!")
    }
    };

    fetchData(); // Call the function to fetch data when the component mounts
}, []);



return (
  <div className="App">
      <nav className="navbar bg-body-tertiary">
        <div className="container">
          <span className="navbar-brand mb-0 h1">K8, Go ReactJS + TypeScript App</span>
        </div>
      </nav>
      <div className='container'>
        <h4 className="text-capitalize">
          Todo List
        </h4>
        { error && 
            <div className="alert alert-danger">
                <span>
                    {error}
                </span>
            </div>
        }
        {
          !error && <>
              <AddTodo />
              <TodoList todos={todos} />
          </>
        }
       
      </div>
  </div>
);
}

export default App;
