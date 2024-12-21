<template>
  <div>
    <h1 class="text-2xl font-bold mb-4">Users</h1>
    <ul class="list-none p-0">
      <UserComponent v-for="user in users" :key="user.ID" :user="user" />
    </ul>
    <div v-if="error" class="text-red-500 mt-4 text-center">{{ error }}</div>
  </div>
</template>

<script setup lang="ts">
import { ref, onMounted } from "vue";
import api  from "../services/api";
import { User,Task} from "../services/api";
import UserComponent from "../components/User.vue";

const users = ref<User[]>([]);
const error = ref<string | null>(null);

const fetchUsers = async () => {
  try {
    users.value = await api.getUsers();
    error.value = null;
  } catch (err) {
    error.value = `Failed to fetch users. Error: ${err.message}`;
  }
};

onMounted(() => {
  fetchUsers();
});
</script>
