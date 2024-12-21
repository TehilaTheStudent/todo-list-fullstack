<template>
  <li class="user-item p-4 border-b border-gray-300">
    <span class="font-semibold">{{ user.Username }}</span> 
    <span class="text-gray-500">(ID: {{ user.ID }}, Password: {{ user.Password }})</span>
  </li>
</template>

<script setup lang="ts">
import { User } from "../services/api";

const props = defineProps<{ user: User }>();
</script>
