<template>
  <div :class="['min-h-screen']" id="app">
    <nav class="bg-transparent shadow-md p-4 flex justify-between items-center border-b border-gray-300">
      <div class="flex space-x-6 text-lg">
        <router-link to="/" :class="linkClass('/')">Home</router-link>
        <router-link to="/tasks" :class="linkClass('/tasks')">Tasks</router-link>
        <router-link to="/users" :class="linkClass('/users')">Users</router-link>
        <router-link to="/login" :class="linkClass('/login')">Login</router-link>
        <router-link to="/register" :class="linkClass('/register')">Register</router-link>
      </div>
      <button @click="logout" class="bg-red-500 text-white px-4 py-2 rounded hover:bg-red-700">Logout</button>
    </nav>
    <router-view class="p-4"></router-view>
  </div>
</template>

<script setup lang="ts">
import { computed } from 'vue';
import { useRouter, useRoute } from 'vue-router';
import api from './services/api';

const router = useRouter();
const route = useRoute();

const isDarkMode = computed(() => {
  // Replace this with your actual dark mode condition
  return window.matchMedia && window.matchMedia('(prefers-color-scheme: dark)').matches;
});

function logout() {
  api.logout();
  router.push('/login');
}

function linkClass(path: string) {
  return route.path === path ? 'text-green-500 hover:text-green-700' : 'text-blue-500 hover:text-blue-700';
}
</script>
