import { createRouter, createWebHistory } from "vue-router";
import type { RouteRecordRaw } from "vue-router";
import HomePage from "@/views/HomePage.vue";
import NotFoundPage from "@/views/NotFoundPage.vue";
import StatsPage from "@/views/StatsPage.vue";

const routes: RouteRecordRaw[] = [
  { path: "/not_found", component: NotFoundPage },
  { path: "/stats/:hash", component: StatsPage },
  { path: "/", component: HomePage },
  { path: "/:pathMatch(.*)*", component: NotFoundPage },
];

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes,
});

export default router;
