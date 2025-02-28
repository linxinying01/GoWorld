<template>
  <div class="auth-page">
    <div class="auth-box">
      <h2>用户登录</h2>
      <form @submit.prevent="handleLogin">
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
            placeholder="请输入密码"
            minlength="8"
          />
          <div v-if="errors.password" class="error">{{ errors.password }}</div>
        </div>

        <button type="submit" :disabled="loading">
          {{ loading ? "登录中..." : "立即登录" }}
        </button>

        <div class="link">
          没有账号？<router-link to="/register">立即注册</router-link>
        </div>

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
    });

    const errors = reactive({
      email: "",
      password: "",
    });

    const loading = ref(false);
    const errorMessage = ref("");

    const validateForm = () => {
      let isValid = true;

      // 清空错误信息
      errors.email = "";
      errors.password = "";

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

      return isValid;
    };

    const handleLogin = async () => {
      if (!validateForm()) return;

      try {
        loading.value = true;
        errorMessage.value = "";

        const response = await axios.post("/auth/login", {
          email: form.email,
          password: form.password,
        });

        // 打印完整响应对象
      console.log('完整响应:', response);
      
      // 打印响应数据
      console.log('响应数据:', response.data);

        localStorage.setItem("jwt", response.data.token);
        router.push("/");
      } catch (error) {
        errorMessage.value = error.response?.data?.error || "登录失败，请稍后重试";
      } finally {
        loading.value = false;
      }
    };

    return {
      form,
      errors,
      loading,
      errorMessage,
      handleLogin,
    };
  },
};
</script>

<style scoped>
.auth-page {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background: #f0f2f5;
}

.auth-box {
  background: white;
  padding: 2rem;
  border-radius: 8px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
  width: 400px;
}

h2 {
  text-align: center;
  margin-bottom: 1.5rem;
}

.form-group {
  margin-bottom: 1rem;
}

label {
  display: block;
  margin-bottom: 0.5rem;
}

input {
  width: 100%;
  padding: 0.5rem;
  border: 1px solid #ddd;
  border-radius: 4px;
}

button {
  width: 100%;
  padding: 0.75rem;
  background: #1890ff;
  color: white;
  border: none;
  border-radius: 4px;
  cursor: pointer;
  margin-top: 1rem;
}

button:disabled {
  background: #8cc8ff;
  cursor: not-allowed;
}

.link {
  margin-top: 1rem;
  text-align: center;
}

.error {
  color: #ff4d4f;
  font-size: 0.875rem;
  margin-top: 0.25rem;
}

.error-message {
  color: #ff4d4f;
  text-align: center;
  margin-top: 1rem;
}
</style>
