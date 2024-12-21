<template>
  <div class="container mx-auto p-4">
    <h1 class="text-2xl font-bold mb-4">Tasks</h1>
    <div v-if="error" class="error text-red-500 mb-4">{{ error }}</div>
    <div class="flex mb-4">
      <input v-model="newTaskName" placeholder="New task name" class="flex-grow p-2 border rounded mr-2 bg-transparent" />
      <button @click="addTask" class="bg-blue-500 text-white px-4 py-2 rounded">Add Task</button>
    </div>
    <ul class="space-y-2">
      <Task
        v-for="task in tasks"
        :key="task.ID"
        :task="task"
        :toggleCompleted="toggleCompleted"
        :deleteTask="deleteTask"
      />
    </ul>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import Task from "../components/Task.vue";
import api from "../services/api";
import { Task as TaskType } from "../services/api";

const tasks = ref<TaskType[]>([]);
const error = ref<string | null>(null);
const newTaskName = ref<string>("");

const fetchTasks = async () => {
  try {
    error.value = null;
    tasks.value = await api.getTasks();
  } catch (err: any) {
    error.value = `Error ${err.response?.status}: ${err.response?.data?.message || "Failed to fetch tasks."}`;
  }
};

const toggleCompleted = async (task: TaskType) => {
  try {
    error.value = null;
    await api.updateTask(task.ID, { IsCompleted: !task.IsCompleted });
    fetchTasks();
  } catch (err: any) {
    error.value = `Error ${err.response?.status}: ${err.response?.data?.message || "Failed to update task."}`;
  }
};

const deleteTask = async (id: number) => {
  try {
    error.value = null;
    await api.deleteTask(id);
    fetchTasks();
  } catch (err: any) {
    error.value = `Error ${err.response?.status}: ${err.response?.data?.message || "Failed to delete task."}`;
  }
};

const addTask = async () => {
  if (!newTaskName.value.trim()) return;
  try {
    error.value = null;
    await api.addTask({ Name: newTaskName.value, IsCompleted: false });
    newTaskName.value = "";
    fetchTasks();
  } catch (err: any) {
    error.value = `Error ${err.response?.status}: ${err.response?.data?.message || "Failed to add task."}`;
  }
};

onMounted(fetchTasks);
</script>

<style scoped>
.error {
  color: red;
}
</style>
