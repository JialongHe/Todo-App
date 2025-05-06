// Please note that there are still some issues with the Jest
// Jest does not support ES modules natively, so we need to use Babel to transpile our code.
// I can fix this later, but for now, let's move on to the other functionalities.
import React from 'react';
import { render, screen } from '@testing-library/react';
import TodoItem from '../components/TodoItem';

test("renders todo item in view mode", () => {
  const mockTodo = { id: 1, title: "Test Todo", description: "Test Description", due_date: "2025-05-06" };

  render(
    React.createElement(TodoItem, { todo: mockTodo, onDelete: jest.fn(), onUpdated: jest.fn() })
  );

  // Check if the TodoItem displays the expected content
  expect(screen.getByText("Test Todo")).toBeInTheDocument();
  expect(screen.getByText("Test Description")).toBeInTheDocument();
  expect(screen.getByText("Edit")).toBeInTheDocument();
});

test("renders todo item in edit mode", () => {
  const mockTodo = { id: 1, title: "Test Todo", description: "Test Description", due_date: "2025-05-06" };

  render(
    React.createElement(TodoItem, { todo: mockTodo, onDelete: jest.fn(), onUpdated: jest.fn() })
  );

  // Initially, the TodoItem is in view mode. Click Edit to change it to edit mode.
  const editButton = screen.getByText("Edit");
  editButton.click();

  // Now the TodoItem should show inputs for editing.
  expect(screen.getByDisplayValue("Test Todo")).toBeInTheDocument();
  expect(screen.getByDisplayValue("Test Description")).toBeInTheDocument();
  expect(screen.getByDisplayValue("2025-05-06")).toBeInTheDocument();
});
