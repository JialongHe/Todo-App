import axios from 'axios';

// Base URL for the API
const API_URL = process.env.REACT_APP_API_URL || 'http://localhost:8080/todos';
console.log("API URL:", API_URL);

// Fetch all todos
export const getTodos = async ({ page = 1, sortBy = 'due_date', sortOrder = "asc", query = '' }) => {
  try {
    const encodedQuery = encodeURIComponent(query || '');
    const response = await axios.get(API_URL, {params: { page, sortBy, sortOrder, q: encodedQuery }});
    return response.data;
  } catch (error) {
    console.error('Error fetching todos:', error);
    throw error;
  }
};

// Create a new todo
export const createTodo = async (todoData) => {
  try {
    const response = await axios.post(API_URL, todoData);
    return response.data;
  } catch (error) {
    console.error('Error creating todo:', error);
    throw error;
  }
};

// Update an existing todo
export const updateTodo = async (id, updatedData) => {
  try {
    const response = await axios.put(`${API_URL}/${id}`, updatedData);
    return response.data;
  } catch (error) {
    console.error('Error updating todo:', error);
    throw error;
  }
};

// Delete a todo
export const deleteTodo = async (id) => {
  try {
    const response = await axios.delete(`${API_URL}/${id}`);
    return response.data;
  } catch (error) {
    console.error('Error deleting todo:', error);
    throw error;
  }
};
