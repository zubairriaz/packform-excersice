import { createWebHistory, createRouter } from "vue-router";
import Home from "../components/Home.vue";
import Order from "../components/Order/index";

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
  },
  {
    path: "/orders",
    name: "Order",
    component: Order,
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;