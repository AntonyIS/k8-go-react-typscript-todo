import React, { useEffect } from 'react';


interface  Todo {
    id :string
    title : string
    description:string
    state:string
}
  
interface TodosListProps {
    todos : Todo[];
}
  
  
  
const TodoList:React.FC<TodosListProps> = ({todos}) =>{

  
  return (
    <div className="card mt-3">
    <div className="card-body">
      <div className="row " >
        {
          todos.map((value:Todo)=>{
            return (
              <div className="col-4 mb-2">
                  <div  className='card'key={value.id}>
                      <div className="card-body">
                        <h4>{value.title}</h4>
                        <p>
                          {value.description}
                        </p>
                        <div className="card-footer">
                          <span>
                            {value.state}
                          </span>
                        </div>
                      </div>
                  </div>
              </div>
            )
          })
        }
      </div>
      
    </div>
</div>
  );
}

export default TodoList;
