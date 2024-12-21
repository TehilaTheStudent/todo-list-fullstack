import { createRouter, createWebHistory } from 'vue-router';


const routes = [
  { path: '/', name: 'home', component: () => import('../views/HomeView.vue') },
  { path: '/tasks', name: 'tasks', component: () => import('../views/TasksView.vue') },
  { path: '/users', name: 'users', component: () => import('../views/UsersView.vue') },
  { path: '/login', name: 'login', component: () => import('../views/LoginView.vue') },
  { path: '/register', name: 'register', component: () => import('../views/RegisterView.vue') },
];


const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
