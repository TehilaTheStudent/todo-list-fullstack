<template>
  <div class="app-container">
    <h1>Todo List</h1>
    <form @submit.prevent="createTask" class="task-form">
      <input
        v-model="newTask.name"
        placeholder="Task name"
        class="task-input"
      />
      <button type="submit" class="add-button">Add Task</button>
    </form>

    <div v-if="error" class="error-message">{{ error }}</div>

    <div class="auth-container">
      <input v-model="username" placeholder="Username" class="auth-input" />
      <input
        type="password"
        v-model="password"
        placeholder="Password"
        class="auth-input"
      />
      <button @click="register" class="auth-button">Register</button>
      <button @click="login" class="auth-button">Login</button>
      <button @click="logout" class="auth-button">Logout</button>
    </div>

    <ul class="task-list">
      <li v-for="task in tasks" :key="task.ID" class="task-item">
        <label class="task-label">
          <input
            type="checkbox"
            :checked="task.IsCompleted"
            @change="toggleCompleted(task)"
            class="task-checkbox"
          />
          {{ task.Name }}
        </label>
        <button @click="deleteTask(task.ID)" class="delete-button">
          Delete
        </button>
      </li>
    </ul>

    <div class="user-list-container">
      <h2>Users</h2>
      <ul class="user-list">
        <li v-for="user in users" :key="user.ID" class="user-item">
          {{ user.Username }} (ID: {{ user.ID }}, Password: {{ user.Password }})
        </li>
      </ul>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from "vue";
import api, { Task, User } from "./services/api";

export default defineComponent({
  setup() {
    // State variables
    const tasks = ref<Task[]>([]);
    const newTask = ref<{ name: string }>({ name: "" });
    const username = ref<string>("");
    const password = ref<string>("");
    const error = ref<string | null>(null);
    const users = ref<User[]>([]);

    // Fetch all tasks
    const fetchTasks = async () => {
      try {
        tasks.value = await api.getTasks();
        error.value = null;
      } catch (err) {
        error.value = `Failed to fetch tasks. Error: ${err.message}`;
      }
    };

    // Fetch all users
    const fetchUsers = async () => {
      try {
        users.value = await api.getUsers();
        error.value = null;
      } catch (err) {
        error.value = `Failed to fetch users. Error: ${err.message}`;
      }
    };

    // Create a new task
    const createTask = async () => {
      if (newTask.value.name.trim()) {
        try {
          await api.addTask({
            Name: newTask.value.name,
            IsCompleted: false,
          });
          newTask.value = { name: "" }; // Clear input
          await fetchTasks(); // Refresh task list
          error.value = null;
        } catch (err) {
          error.value = `Failed to create task. Error: ${err.message}`;
        }
      }
    };

    // Toggle task completion
    const toggleCompleted = async (task: Task) => {
      try {
        const updatedTask = { ...task, isCompleted: !task.IsCompleted };
        await api.updateTask(task.ID, updatedTask);
        await fetchTasks();
        error.value = null;
      } catch (err) {
        error.value = `Failed to update task. Error: ${err.message}`;
      }
    };

    // Delete a task
    const deleteTask = async (id: number) => {
      try {
        await api.deleteTask(id);
        await fetchTasks();
        error.value = null;
      } catch (err) {
        error.value = `Failed to delete task. Error: ${err.message}`;
      }
    };

    // Register a new user
    const register = async () => {
      try {
        await api.register(username.value, password.value);
        error.value = null;
      } catch (err) {
        error.value = `Failed to register. Error: ${err.message}`;
      }
    };

    // Login a user
    const login = async () => {
      try {
        await api.login(username.value, password.value);
        error.value = null;
      } catch (err) {
        error.value = `Failed to login. Error: ${err.message}`;
      }
    };

    // Logout a user
    const logout = async () => {
      try {
        api.logout();
        error.value = null;
      } catch (err) {
        error.value = `Failed to logout. Error: ${err.message}`;
      }
    };

    // Lifecycle hook
    onMounted(() => {
      fetchTasks();
      fetchUsers();
    });

    return {
      tasks,
      newTask,
      username,
      password,
      createTask,
      toggleCompleted,
      deleteTask,
      register,
      login,
      logout,
      error,
      users,
    };
  },
});
</script>

<style>
.app-container {
  max-width: 600px;
  margin: 0 auto;
  padding: 20px;
  font-family: Arial, sans-serif;
  background-color: #f9f9f9;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.1);
}

h1 {
  text-align: center;
  color: #333;
}

.task-form {
  display: flex;
  justify-content: space-between;
  margin-bottom: 20px;
}

.task-input {
  flex: 1;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
  margin-right: 10px;
}

.add-button {
  padding: 10px 20px;
  background-color: #28a745;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.add-button:hover {
  background-color: #218838;
}

.error-message {
  color: red;
  margin-bottom: 20px;
  text-align: center;
}

.auth-container {
  display: flex;
  flex-direction: column;
  margin-bottom: 20px;
}

.auth-input {
  padding: 10px;
  margin-bottom: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
}

.auth-button {
  padding: 10px;
  margin-bottom: 10px;
  background-color: #007bff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.auth-button:hover {
  background-color: #0056b3;
}

.task-list {
  list-style: none;
  padding: 0;
}

.task-item {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 10px;
  border-bottom: 1px solid #ccc;
}

.task-label {
  display: flex;
  align-items: center;
  color: black;
}

.task-checkbox {
  margin-right: 10px;
}

.delete-button {
  padding: 5px 10px;
  background-color: #dc3545;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}

.delete-button:hover {
  background-color: #c82333;
}

.user-list-container {
  margin-top: 20px;
}

.user-list {
  list-style: none;
  padding: 0;
}

.user-item {
  padding: 10px;
  border-bottom: 1px solid #ccc;
  color: black;
}
</style>
