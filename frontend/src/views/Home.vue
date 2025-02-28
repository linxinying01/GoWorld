<template>
  <div class="home">
    <h1>欢迎回来, {{ userEmail }}</h1>
    <button @click="handleLogout">退出登录</button>
  </div>
</template>

<script>
import { ref } from "vue";
import { useRouter } from "vue-router";

export default {
  setup() {
    const router = useRouter();
    const userEmail = ref(localStorage.getItem("email") || "");

    const handleLogout = () => {
      localStorage.removeItem("jwt");
      localStorage.removeItem("email");
      router.push("/login");
    };

    return {
      userEmail,
      handleLogout,
    };
  },
};
</script>

<style scoped>
.home {
  text-align: center;
  padding: 2rem;
}

button {
  padding: 0.5rem 1rem;
  background: #ff4d4f;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
}
</style>
