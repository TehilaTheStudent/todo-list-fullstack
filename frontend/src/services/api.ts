import axios from "axios";

// Define a Task type for better type safety
export interface Task {
  ID: number;
  Name: string;
  IsCompleted: boolean;
}

// Axios instance
const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

// Add a response interceptor to handle and log errors
apiClient.interceptors.response.use(
  (response) => response, // Pass through successful responses
  (error) => {
    console.error(
      "API Error:",
      error.response ? error.response.data : error.message
    );
    return Promise.reject(error); // Ensure error is still propagated
  }
);

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
