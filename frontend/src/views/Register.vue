<template>
  <div class="auth-page">
    <div class="auth-box">
      <h2>用户注册</h2>
      <form @submit.prevent="handleRegister">
        <div class="form-group">
          <label>邮箱</label>
          <input v-model="form.email" type="email" required placeholder="请输入邮箱" />
          <div v-if="errors.email" class="error">{{ errors.email }}</div>
        </div>

        <div class="form-group">
          <label>密码</label>
          <input
            v-model="form.password"
            type="password"
            required
            placeholder="至少8位密码"
            minlength="8"
          />
          <div v-if="errors.password" class="error">{{ errors.password }}</div>
        </div>

        <div class="form-group">
          <label>确认密码</label>
          <input
            v-model="form.confirmPassword"
            type="password"
            required
            placeholder="请再次输入密码"
          />
          <div v-if="errors.confirmPassword" class="error">
            {{ errors.confirmPassword }}
          </div>
        </div>

        <button type="submit" :disabled="loading">
          {{ loading ? "注册中..." : "立即注册" }}
        </button>

        <div class="link">已有账号？<router-link to="/login">立即登录</router-link></div>

        <div v-if="errorMessage" class="error-message">{{ errorMessage }}</div>
      </form>
    </div>
  </div>
</template>

<script>
import { reactive, ref } from "vue";
import { useRouter } from "vue-router";
import axios from "../api/axios";

export default {
  setup() {
    const router = useRouter();
    const form = reactive({
      email: "",
      password: "",
      confirmPassword: "",
    });

    const errors = reactive({
      email: "",
      password: "",
      confirmPassword: "",
    });

    const loading = ref(false);
    const errorMessage = ref("");

    const validateForm = () => {
      let isValid = true;

      errors.email = "";
      errors.password = "";
      errors.confirmPassword = "";

      // 邮箱验证
      if (!/^\w+([.-]?\w+)*@\w+([.-]?\w+)*(\.\w{2,3})+$/.test(form.email)) {
        errors.email = "请输入有效的邮箱地址";
        isValid = false;
      }

      // 密码验证
      if (form.password.length < 8) {
        errors.password = "密码至少需要8个字符";
        isValid = false;
      }

      // 确认密码验证
      if (form.password !== form.confirmPassword) {
        errors.confirmPassword = "两次输入的密码不一致";
        isValid = false;
      }

      return isValid;
    };

    const handleRegister = async () => {
      if (!validateForm()) return;

      try {
        loading.value = true;
        errorMessage.value = "";

        const response = await axios.post("/auth/register", {
          email: form.email,
          password: form.password,
        });

        localStorage.setItem("jwt", response.data.token);
        router.push("/");
      } catch (error) {
        errorMessage.value = error.response?.data?.error || "注册失败，请稍后重试";
      } finally {
        loading.value = false;
      }
    };

    return {
      form,
      errors,
      loading,
      errorMessage,
      handleRegister,
    };
  },
};
</script>

<style scoped>
/* 复用登录页的样式 */
</style>
