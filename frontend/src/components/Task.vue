<template>
  <li class="task-item flex justify-between items-center p-2 border-b">
    <label class="task-label flex items-center">
      <input
        type="checkbox"
        :checked="task.IsCompleted"
        @change="toggleCompleted(task)"
        class="task-checkbox mr-2"
      />
      {{ task.Name }}
    </label>
    <button @click="deleteTask(task.ID)" class="delete-button bg-red-500 text-white px-2 py-1 rounded">Delete</button>
  </li>
</template>

<script lang="ts" setup>
import { Task } from "../services/api";

const props = defineProps<{
  task: Task;
  toggleCompleted: (task: Task) => void;
  deleteTask: (id: number) => void;
}>();
</script>

<style scoped>
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
