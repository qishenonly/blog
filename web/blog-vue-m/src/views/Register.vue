<template>
  <div class="register-container">
    <div class="register-form">
      <div class="form-header">
        <h1>注册</h1>
        <p>创建您的新帐户</p>
      </div>
      <form @submit="register">
        <div class="form-group">
          <label for="username">用户名：</label>
          <input type="text" id="username" v-model="username" required>
        </div>
        <div class="form-group">
          <label for="email">邮箱：</label>
          <input type="text" id="email" v-model="email" required>
        </div>
        <div class="form-group">
          <label for="password">密码：</label>
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
            <canvas ref="captchaCanvas" width="120" @click="generateCaptcha()"
                    height="40" style="margin-left: 5%;"></canvas>
          </div>
        </div>
        <button type="submit" style="margin-top: 10px;" @click="register">注册</button>
      </form>
      <p class="login-link">已经有帐户？<router-link to="/auth/login">登录</router-link></p>
    </div>
  </div>
</template>

<script>
import {fetchRegister, fetchRegisterCode} from '../api'
export default {
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
    this.getRegisterCode();
    this.generateCaptcha();
  },
  methods: {
    register(event) {
      event.preventDefault();
      // 在这里处理注册逻辑，包括用户名、密码、确认密码和验证码的验证
      // 如果注册成功，可以跳转到登录页面或执行其他操作
      if (this.password !== this.confirmPassword) {
        this.$message.warning('两次输入的密码不一致');
        return;
      }

      // 判断密码是否过于简单，要求八位以上，且包含数字和字母
      const reg = /^(?=.*[a-zA-Z])(?=.*\d)[^]{8,}$/;
      if (!reg.test(this.password)) {
        this.$message.warning('密码过于简单！要求八位以上，且包含数字和字母！');
        return;
      }

      // 判断邮箱格式是否正确
      const regEmail = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
      if (!regEmail.test(this.email)) {
        this.$message.warning('邮箱格式不正确！');
        return;
      }


      let data = {
        username : this.username,
        password : this.password,
        email : this.email,
        code : this.captchaInput
      }

      fetchRegister(data)
          .then(res => {
            if (res.data.code === 200) {
              this.$message.success('注册成功');
              this.$router.push('/auth/login');
            } else {
              this.$message.error(res.data.data);
            }
          })
          .catch(error => {
            console.error('Error:', error);
          });
    },

    async getRegisterCode() {
      await fetchRegisterCode().then(res => {
        this.registerCode = res.data.data || []
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
.register-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
}

.register-form {
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


</style>
