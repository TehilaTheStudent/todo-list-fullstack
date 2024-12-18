<template>
  <div class="app-container">
    <h1>Todo List</h1>
    <form @submit.prevent="createTask" class="task-form">
      <input v-model="newTask.name" placeholder="Task name" class="task-input" />
      <button type="submit" class="add-button">Add Task</button>
    </form>

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
        <button @click="deleteTask(task.ID)" class="delete-button">Delete</button>
      </li>
    </ul>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref, onMounted } from "vue";
import api, { Task } from "./services/api";

export default defineComponent({
  setup() {
    // State variables
    const tasks = ref<Task[]>([]);
    const newTask = ref<{ name: string }>({ name: "" });

    // Fetch all tasks
    const fetchTasks = async () => {
      tasks.value = await api.getTasks();
    };

    // Create a new task
    const createTask = async () => {
      if (newTask.value.name.trim()) {
        await api.addTask({
          Name: newTask.value.name,
          IsCompleted: false,
        });
        newTask.value = { name: "" }; // Clear input
        await fetchTasks(); // Refresh task list
      }
    };

    // Toggle task completion
    const toggleCompleted = async (task: Task) => {
      const updatedTask = { ...task, isCompleted: !task.IsCompleted };
      await api.updateTask(task.ID, updatedTask);
      await fetchTasks();
    };

    // Delete a task
    const deleteTask = async (id: number) => {
      await api.deleteTask(id);
      await fetchTasks();
    };

    // Lifecycle hook
    onMounted(fetchTasks);

    return {
      tasks,
      newTask,
      createTask,
      toggleCompleted,
      deleteTask,
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
</style>
