import axios from "axios";

// Define a Task type for better type safety
export interface Task {
  ID: number;
  Name: string;
  IsCompleted: boolean;
}

// Axios instance
const apiClient = axios.create({
  baseURL: "http://localhost:8080", // Replace with your backend URL
  headers: {
    "Content-Type": "application/json",
  },
});

export default {
  async getTasks(): Promise<Task[]> {
    const response = await apiClient.get<Task[]>("/tasks");
    return response.data;
  },

  async addTask(task: Partial<Task>): Promise<Task> {
    const response = await apiClient.post<Task>("/tasks", task);
    return response.data;
  },

  async updateTask(id: number, updatedTask: Partial<Task>): Promise<Task> {
    const response = await apiClient.put<Task>(`/tasks/${id}`, updatedTask);
    return response.data;
  },

  async deleteTask(id: number): Promise<void> {
    await apiClient.delete(`/tasks/${id}`);
  },
};
