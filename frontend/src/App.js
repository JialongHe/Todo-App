import React, { useState } from 'react';
import CreateTodo from './components/CreateTodo';
import TodoList from './components/TodoList';
import './App.css';

const App = () => {
  const [refreshKey, setRefreshKey] = useState(0);

  return (
    <div className="App">
      <h1>To-Do App</h1>
      <div className="todo-container">
        <CreateTodo onCreate={() => setRefreshKey(prev => prev + 1)} />
        <TodoList key={refreshKey} />
      </div>
    </div>
  );
};

export default App;