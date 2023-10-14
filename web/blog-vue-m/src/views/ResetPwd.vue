<template>
  <div class="reset-password-container">
    <div v-if="showResetView === false" class="inactive-link">
      <!-- 未激活链接和刷新页面的按钮 -->
      <p class="inactive-text">未激活链接！ 请在邮件中激活链接！</p>
      <button @click="refreshPage" class="refresh-button">刷新页面</button>
    </div>
    <div v-else>
      <div class="reset-password-form">
        <div class="form-header">
          <h1>重置密码</h1>
          <p>请输入您的新密码</p>
        </div>
        <form @submit="resetPassword">
          <div class="form-group">
            <label for="email">邮箱：</label>
            <input type="text" id="email" v-model="email" required>
          </div>
          <div class="form-group">
            <label for="password">新密码：</label>
            <input type="password" id="password" v-model="password" required>
          </div>
          <div class="form-group">
            <label for="confirmPassword">确认密码：</label>
            <input type="password" id="confirmPassword" v-model="confirmPassword" required>
          </div>
          <div class="captcha-container">
            <label for="code" style="margin-bottom: 2%;">验证码：</label>
            <div class="cc" style="display: flex">
              <input type="text" v-model="captchaInput" required>
              <canvas ref="captchaCanvas" width="120" @click="generateCaptcha()" height="40" style="margin-left: 5%;"></canvas>
            </div>
          </div>
          <button type="submit" style="margin-top: 10px;">重置密码</button>
        </form>
        <p class="login-link">已经有帐户？<router-link to="/auth/login">登录</router-link></p>
      </div>
    </div>
  </div>
</template>

<script>
import {fetchResetPwd, fetchResetPwdCode} from '../api'
export default {
  props: {
    showResetView: Boolean, // 从路由参数中传递的数据
    toEmail: String,
  },
  data() {
    return {
      username: '',
      password: '',
      email: '',
      confirmPassword: '',
      captcha: '',
      captchaInput: '',
      registerCode: '',
    };
  },
  mounted() {
    console.log(this.showResetView)
    this.getResetCode();
    this.generateCaptcha();
  },
  methods: {
    resetPassword(event) {
      event.preventDefault();
      // 在这里处理注册逻辑，包括用户名、密码、确认密码和验证码的验证
      // 如果注册成功，可以跳转到登录页面或执行其他操作

      // 判断密码是否过于简单，要求八位以上，且包含数字和字母
      const reg = /^(?=.*[a-zA-Z])(?=.*\d)[^]{8,}$/;
      if (!reg.test(this.password)) {
        this.$message.warning('密码过于简单！要求八位以上，且包含数字和字母！');
        return;
      }

      let data = {
        username : this.username,
        password : this.password,
        email : this.email,
        code : this.captchaInput
      }

      fetchResetPwd(data)
          .then(response => response.json())
          .then(data => {
            console.log(data);
          })
          .catch(error => {
            console.error('Error:', error);
          });
    },

    async getResetCode() {
      await fetchResetPwdCode().then(res => {
        this.registerCode = res.data.data || []
      }).catch(err => {
        console.log(err)
      })
    },

    refreshPage() {
      // 刷新页面的逻辑
      // location.reload();
      fetchResetPwd(this.$route.params.email).then(res => {
        console.log(res)
        this.showResetView = res.data.data.custom_code === 41000;
      }).catch(err => {
        console.log(err)
      })
    },

    async generateCaptcha() {
      await this.getRegisterCode();

      const canvas = this.$refs.captchaCanvas;
      const ctx = canvas.getContext('2d');


      // 随机生成验证码文字
      const captchaText = this.registerCode;


      // 设置字体样式
      ctx.font = '40px Arial';

      // 绘制背景色
      ctx.fillStyle = '#f0f0f0';
      ctx.fillRect(0, 0, canvas.width, canvas.height);

      // 绘制验证码文字
      ctx.fillStyle = this.getRandomColor();
      ctx.fillText(captchaText, 15, 32);

      // 添加干扰线
      this.drawRandomLines(ctx);

      // 添加干扰点
      this.drawRandomDots(ctx);
    },
    generateRandomText() {
      // 生成随机的验证码文字
      const characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789';
      let captchaText = '';

      for (let i = 0; i < 4; i++) {
        const randomIndex = Math.floor(Math.random() * characters.length);
        captchaText += characters.charAt(randomIndex);
      }

      return captchaText;
    },
    getRandomColor() {
      // 生成随机的颜色
      const letters = '0123456789ABCDEF';
      let color = '#';

      for (let i = 0; i < 6; i++) {
        color += letters[Math.floor(Math.random() * 16)];
      }

      return color;
    },
    drawRandomLines(ctx) {
      // 绘制随机干扰线
      for (let i = 0; i < 5; i++) {
        ctx.beginPath();
        ctx.moveTo(Math.random() * 200, Math.random() * 80);
        ctx.lineTo(Math.random() * 200, Math.random() * 80);
        ctx.strokeStyle = this.getRandomColor();
        ctx.stroke();
      }
    },
    drawRandomDots(ctx) {
      // 绘制随机干扰点
      for (let i = 0; i < 50; i++) {
        ctx.beginPath();
        ctx.arc(Math.random() * 200, Math.random() * 80, 2, 0, Math.PI * 2);
        ctx.fillStyle = this.getRandomColor();
        ctx.fill();
      }
    },
    refreshCaptcha() {
      // 刷新验证码
      this.generateCaptcha();
      this.captchaInput = '';
    },

  },
};
</script>

<style scoped>
.reset-password-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.reset-password-form {
  background-color: #ffffff;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  width: 400px;
}

.form-header {
  text-align: center;
  margin-bottom: 20px;
}

.form-header h1 {
  font-size: 28px;
}

.form-header p {
  font-size: 14px;
  color: #666;
}

.form-group {
  margin-bottom: 20px;
}

.form-group label {
  display: block;
  font-size: 16px;
  margin-bottom: 8px;
}

input[type="text"],
input[type="password"] {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 4px;
  font-size: 16px;
}

.captcha-input {
  display: flex;
  align-items: center;
}

.captcha-image {
  cursor: pointer;
  margin-left: 10px;
}

button[type="submit"] {
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 4px;
  padding: 10px 20px;
  font-size: 16px;
  cursor: pointer;
}

.login-link {
  text-align: center;
  margin-top: 20px;
}

.login-link a {
  text-decoration: none;
  color: #007bff;
  font-weight: bold;
}

.login-link a:hover {
  text-decoration: underline;
}

.inactive-link {
  text-align: center;
  background-color: #f5f5f5;
  padding: 20px;
  border-radius: 5px;
  box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
  margin: 20px auto;
  max-width: 300px;
}

.inactive-text {
  font-size: 24px;
  color: #007bff;
}

.refresh-button {
  background-color: #007bff;
  color: #fff;
  font-weight: bold;
  padding: 10px 20px;
  border: none;
  border-radius: 5px;
  margin-top: 20px;
  cursor: pointer;
  transition: background-color 0.3s;
}

.refresh-button:hover {
  background-color: #0056b3;
}
</style>
