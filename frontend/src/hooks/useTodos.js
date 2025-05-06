import { useEffect, useState, useCallback } from 'react';
import { getTodos } from '../api/todoApi';

export const useTodos = (page, sortBy, sortOrder, query) => {
  const [todos, setTodos] = useState([]);
  const [meta, setMeta] = useState({});
  const [loading, setLoading] = useState(true);

  const fetchTodos = useCallback(() => {
    setLoading(true);
    getTodos({ page, sortBy, sortOrder, query })
      .then(res => {
        if (res.results && Array.isArray(res.results)) {
          setTodos(res.results);
          setMeta({ count: res.count, page: res.page, limit: res.limit });
        } else {
          setTodos([]);
          console.log('Todo-list is empty', res);
        }
      })
      .catch(err => console.error(err))
      .finally(() => setLoading(false));
  }, [page, sortBy, sortOrder, query]);

  useEffect(() => {
    fetchTodos();
  }, [fetchTodos]);

  return { todos, meta, loading, refetch: fetchTodos};
};
