import React, { useState } from 'react';
import TodoItem from './TodoItem';
import { useTodos } from '../hooks/useTodos';
import { deleteTodo } from '../api/todoApi';
import '../styles/TodoList.css';

const TodoList = () => {
  const [page, setPage] = useState(1);
  const [sortBy, setSortBy] = useState('due_date');
  const [sortOrder, setSortOrder] = useState('asc');
  const [query, setQuery] = useState('');
  const { todos, meta, loading, refetch } = useTodos(page, sortBy, sortOrder, query);

  const handleDelete = async (id) => {
    await deleteTodo(id);
    const newCount = meta.count - 1;
    const newPage = newCount <= (page - 1) * meta.limit && page > 1 ? page - 1 : page;
  
    if (newPage !== page) {
      setPage(newPage);
    } else {
      refetch();
    }
  };

  const handleSearch = (e) => {
    setQuery(e.target.value);
    setPage(1); // reset to page 1 when searching
  };

  if (loading) return <p>Loading...</p>;

  return (
    <div className="todo-list">
      <h2>To-Do List</h2>
      {(!Array.isArray(todos) || todos.length === 0) && <p>No to-dos found.</p>}
      <input
        type="text"
        placeholder="Search..."
        value={query}
        onChange={handleSearch}
      />
      <select value={sortBy} onChange={(e) => setSortBy(e.target.value)}>
        <option value="due_date">Due Date</option>
        <option value="title">Title</option>
      </select>
      <select value={sortOrder} onChange={(e) => setSortOrder(e.target.value)}>
        <option value="asc">Ascending</option>
        <option value="desc">Descending</option>
      </select>
      <ul>
        {todos.map(todo => (
          <TodoItem key={todo.id} todo={todo} onDelete={handleDelete} onUpdated={refetch} />
        ))}
      </ul>
      <button 
        onClick={() => setPage(p => p - 1)} 
        disabled={Object.keys(meta).length < 1 || meta.page <= 1}
      >
        Previous
      </button>
      <span> Page {meta.page || 1} / {Math.ceil(meta.count / meta.limit) || 1} </span>
      <button 
        onClick={() => setPage(p => p + 1)} 
        disabled={Object.keys(meta).length < 1 || meta.page * meta.limit >= meta.count}
      >
        Next
      </button>
    </div>
  );
};

export default TodoList;