<template>
  <div class="flex items-center justify-center min-h-screen">
    <div class="p-8 rounded border-4 border-indigo-500 w-full max-w-md">
      <h1 class="text-2xl font-bold mb-6">Login</h1>
      <form @submit.prevent="handleLogin" class="space-y-4">
        <div>
          <label for="username" class="block text-sm font-medium text-gray-700"
            >Username:</label
          >
          <input
            type="text"
            id="username"
            v-model="username"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm bg-transparent"
          />
        </div>
        <div>
          <label for="password" class="block text-sm font-medium text-gray-700"
            >Password:</label
          >
          <input
            type="password"
            id="password"
            v-model="password"
            class="mt-1 block w-full px-3 py-2 border border-gray-300 rounded-md shadow-sm focus:outline-none focus:ring-indigo-500 focus:border-indigo-500 sm:text-sm bg-transparent"
          />
        </div>
        <button
          type="submit"
          class="w-full py-2 px-4 bg-indigo-600 text-white font-semibold rounded-md shadow hover:bg-indigo-700 focus:outline-none focus:ring-2 focus:ring-offset-2 focus:ring-indigo-500"
        >
          Login
        </button>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref } from "vue";
import { useRouter } from "vue-router";
import api from "@/services/api";

const username = ref("");
const password = ref("");
const router = useRouter();

const handleLogin = async () => {
  try {
    await api.login(username.value, password.value);
    router.push("/tasks");
  } catch (error) {
    alert("Login failed. Please check your credentials and try again.");
  }
};
</script>

<style scoped></style>
