import axios from "axios";

// Define a Task type for better type safety
export interface Task {
  ID: number;
  Name: string;
  IsCompleted: boolean;
}
export interface User {
  ID: number;
  Username: string;
  Password: string;
}

// Axios instance
const apiClient = axios.create({
  baseURL: import.meta.env.VITE_API_BASE_URL,
  headers: {
    "Content-Type": "application/json",
  },
});

function saveAccessToken(authResult) {
  localStorage.setItem("access_token", authResult.token);
  setAuthorizationBearer();
}

function setAuthorizationBearer() {
  const accessToken = localStorage.getItem("access_token");
  if (accessToken) {
    apiClient.defaults.headers.common["Authorization"] = `Bearer ${accessToken}`;
  }
}

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
    setAuthorizationBearer();

    const response = await apiClient.get<Task[]>("/tasks");
    return response.data;
  },

  async addTask(task: Partial<Task>): Promise<Task> {

    const response = await apiClient.post<Task>("/tasks", task);
    return response.data;
  },

  async updateTask(id: number, updatedTask: Partial<Task>): Promise<Task> {
    setAuthorizationBearer();

    const response = await apiClient.put<Task>(`/tasks/${id}`, updatedTask);
    return response.data;
  },

  async deleteTask(id: number): Promise<void> {
    setAuthorizationBearer();

    await apiClient.delete(`/tasks/${id}`);
  },
  async login(username: string, password: string) {
    const res = await apiClient.post("/login", { username, password });
    saveAccessToken(res.data);
  },
  async register(username: string, password: string) {
    const res = await apiClient.post("/register", { username, password });
  },
  async getUsers(): Promise<User[]> {
    setAuthorizationBearer();

    const response = await apiClient.get<User[]>("/users");
    return response.data;
  },
  logout:()=>{
    localStorage.setItem("access_token", "");
  },
};
