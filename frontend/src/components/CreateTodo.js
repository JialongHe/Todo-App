import React, { useState } from 'react';
import { createTodo } from '../api/todoApi';
import '../styles/CreateTodo.css';

const CreateTodo = ({ onCreate }) => {
  const [title, setTitle] = useState('');
  const [description, setDescription] = useState('');
  const [dueDate, setDueDate] = useState('');

  const handleSubmit = async (e) => {
    e.preventDefault();

    try {
      const newTodo = {
        title,
        description,
        "due_date": new Date(dueDate).toISOString()
      };
      const response = await createTodo(newTodo);
      console.log('New to-do created:', response);
    } catch (error) {
      console.error('Error creating to-do:', error);
    }
    if (onCreate) onCreate();
  };

  return (
    <div className="create-todo">
      <h2>Create To-Do</h2>
      <form onSubmit={handleSubmit} className="create-todo-form">
        <div className="form-group">
          <label htmlFor="title">Title</label>
          <input
            id="title"
            type="text"
            placeholder="Enter title"
            value={title}
            onChange={(e) => setTitle(e.target.value)}
            required
            className="form-input"
          />
        </div>
        <div className="form-group">
          <label htmlFor="description">Description</label>
          <textarea
            id="description"
            placeholder="Enter description"
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            className="form-input"
          />
        </div>
        <div className="form-group">
          <label htmlFor="dueDate">Due Date</label>
          <input
            id="dueDate"
            type="date"
            value={dueDate}
            onChange={(e) => setDueDate(e.target.value)}
            required
            className="form-input"
          />
        </div>
        <button type="submit" className="submit-btn">Create</button>
      </form>
    </div>
  );
};

export default CreateTodo;
