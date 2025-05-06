import React, { useState } from 'react';
import { updateTodo } from '../api/todoApi';
import '../styles/TodoItem.css';

const TodoItem = ({ todo, onDelete, onUpdated }) => {
  const [editing, setEditing] = useState(false);
  const [form, setForm] = useState({ title: todo.title, description: todo.description, due_date: todo.due_date });

  const handleSave = async () => {
    const updatedTodo = {
      title: form.title,
      description: form.description,
      due_date: new Date(form.due_date).toISOString(),
    };

    console.log('Saving todo:', updatedTodo);
    try {
      await updateTodo(todo.id, updatedTodo);
      onUpdated(); // trigger refetch or state update
      setEditing(false);
    } catch (err) {
      console.error("Failed to update todo:", err);
    }
  };

  return (
    <li className={`todo-item ${editing ? 'editing' : ''}`}>
      {editing ? (
        <>
          <div className="todo-content">
            <div className="todo-details">
              <input value={form.title} onChange={(e) => setForm({ ...form, title: e.target.value })} />
              <textarea value={form.description} onChange={(e) => setForm({ ...form, description: e.target.value })} />
              <input
                type="date"
                value={form.due_date}
                onChange={(e) => setForm({ ...form, due_date: e.target.value })}
              />
            </div>
            <div className="todo-buttons">
              <button className="edit-btn" onClick={handleSave}>Save</button>
              <button className="delete-btn" onClick={() => setEditing(false)}>Cancel</button>
            </div>
          </div>
        </>
      ) : (
        <div className="todo-content">
          <div className="todo-details">
            <strong>{todo.title}</strong>
            <p>{todo.description}</p>
            <p>{new Date(todo.due_date).toLocaleDateString()}</p>
          </div>
          <div className="todo-buttons">
            <button className="edit-btn" onClick={() => setEditing(true)}>Edit</button>
            <button className="delete-btn" onClick={() => onDelete(todo.id)}>Delete</button>
          </div>
        </div>
      )}
    </li>
  );
};

export default TodoItem;
