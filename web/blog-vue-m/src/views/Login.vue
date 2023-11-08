<template>
  <div class="login-container">
    <div class="login-form">
      <div class="form-header">
        <h1>欢迎回来！</h1>
        <p>登录到您的帐户</p>
      </div>
      <form @submit="login">
        <div class="form-group">
          <label for="email">邮箱：</label>
          <input type="text" id="email" v-model="email" required>
        </div>
        <div class="form-group">
          <label for="password">密码：</label>
          <div class="password-input">
            <input type="password" id="password" v-model="password" required :type="showPassword ? 'text' : 'password'">
            <button @click="toggleShowPassword">
              <i class="fas" :class="showPassword ? 'fa-eye-slash' : 'fa-eye'"></i>
            </button>
          </div>
        </div>
        <div class="form-group">
          <div class="captcha-container">
            <label for="code" style="margin-bottom: 2%;">验证码：</label>
            <div class="cc" style="display: flex">
              <input type="text" v-model="captchaInput" required>
              <canvas ref="captchaCanvas" width="120" @click="generateCaptcha()"
                      height="40" style="margin-left: 5%;"></canvas>
            </div>
          </div>
        </div>
        <button type="submit">登录</button>
      </form>
      <p class="register-link">还没有帐户？<router-link to="/auth/register">注册</router-link></p>
<!--      <p class="register-link">忘记密码？<router-link to="/reset_pwd">重置</router-link></p>-->
      <p class="register-link">忘记密码？
        <button class="reset_but" @click="toResetPwdPage">重置</button>
      </p>
      <div class="social-login">
        <p>或使用以下方式登录：</p>
        <div class="icons">
          <img src="../../public/github.png" alt="" class="img_github"
               @click="open('http://localhost:8080/api/auth/github')">
          <img src="../../public/gitee.png" alt="" class="img_gitee"
               >
        </div>
      </div>
    </div>
    <div class="login-image">
      <!-- 放置登录页面的图片 -->
      <!--      <img src="/path/to/login-image.jpg" alt="登录图片">-->
    </div>
  </div>
</template>


<script>
import {fetchLogin, fetchLoginCode, fetchRegister, fetchResetPwd} from '../api'
import router from "@/router";
export default {
  data() {
    return {
      email: '',
      password: '',
      showPassword: false,
      captcha: '',
      captchaInput: '',
      loginCode: '',
    };
  },
  mounted() {
    this.generateCaptcha();
    this.getLoginCode();
  },
  methods: {
    login(event) {
      event.preventDefault();
      // 在这里处理登录逻辑，可以向后端发送登录请求
      // 使用 this.username、this.password 和 this.captcha 获取用户输入
      // 如果登录成功，可以跳转到其他页面或执行其他操作
      if (this.email === '') {
        this.$message({
          message: '请输入邮箱',
          type: 'warning'
        });
        return
      } else if (this.password === '') {
        this.$message({
          message: '请输入密码',
          type: 'warning'
        });
        return
      } else if (this.captchaInput === '') {
        this.$message({
          message: '请输入验证码',
          type: 'warning'
        });
        return
      }

      // 判断邮箱格式是否正确
      const regEmail = /^([a-zA-Z]|[0-9])(\w|\-)+@[a-zA-Z0-9]+\.([a-zA-Z]{2,4})$/;
      if (!regEmail.test(this.email)) {
        this.$message.warning('邮箱格式不正确！');
        return;
      }

      let data = {
        email: this.email,
        password: this.password,
        code: this.captchaInput
      }

      fetchLogin(data)
          .then(res => {
            if (res.data.code === 400) {
              this.$message({
                message: res.data.data,
                type: 'error'
              });
            } else {
              localStorage.setItem('token', res.data.data.token);
              this.$message({
                message: '登录成功',
                type: 'success'
              });
              this.$router.push({
                name: 'home',
              });
              window.location.reload();
            }
          })
          .catch(error => {
            console.error('Error:', error);
          });
    },

    async getLoginCode() {
      await fetchLoginCode().then(res => {
        this.loginCode = res.data.data || []
      }).catch(err => {
        console.log(err)
      })
    },

    toggleShowPassword() {
      this.showPassword = !this.showPassword;
    },

    open(url) {
      const a = document.createElement('a');
      a.href = url;
      a.target = '_blank';
      a.click()
    },

    toResetPwdPage() {
      if (this.email === '') {
        this.$message({
          message: '请输入邮箱',
          type: 'warning'
        });
        return
      }

      fetchResetPwd(this.email).then(res => {
        console.log(res)
        if (res.data.code === 400) {
          this.$router.push({
            name: 'reset_pwd',
            params: {
              // 将数据作为参数传递
              showResetView: true,
              email: this.email,
            }
          });
        } else {
          if (res.data.data.activated === false) {
            this.$router.push({
              name: 'reset_pwd',
              params: {
                // 将数据作为参数传递
                showResetView: false,
                email: this.email,
              }
            });
          } else {
            this.$router.push({
              name: 'reset_pwd',
              params: {
                // 将数据作为参数传递
                showResetView: true,
                email: this.email,
              }
            });
          }
        }
      }).catch(err => {
        console.log(err)
      })
    },

    async generateCaptcha() {
      const canvas = this.$refs.captchaCanvas;
      const ctx = canvas.getContext('2d');

      await this.getLoginCode()
      // 随机生成验证码文字
      const captchaText = this.loginCode;

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
.login-container {
  display: flex;
  align-items: center;
  justify-content: center;
  min-height: 100vh;
  background-color: #f9f9f9;
}

.login-form {
  background-color: #ffffff;
  padding: 40px;
  box-shadow: 0 0 20px rgba(0, 0, 0, 0.1);
  border-radius: 5px;
  width: 400px;
  text-align: center;
}

.form-header {
  margin-bottom: 20px;
}

.form-header h1 {
  font-size: 32px;
  margin-bottom: 10px;
}

.form-header p {
  color: #888;
}

.form-group {
  margin-bottom: 20px;
}

label {
  font-size: 18px;
  font-weight: bold;
  display: block;
  margin-bottom: 10px;
}

input[type="text"],
input[type="password"] {
  width: 100%;
  padding: 12px;
  border: 1px solid #ccc;
  border-radius: 5px;
  outline: none;
  font-size: 16px;
}

button[type="submit"] {
  width: 100%;
  padding: 12px;
  background-color: #007bff;
  color: #fff;
  border: none;
  border-radius: 5px;
  font-size: 18px;
  cursor: pointer;
  transition: background-color 0.3s;
}

button[type="submit"]:hover {
  background-color: #0056b3;
}

.register-link {
  text-align: center;
  margin-top: 20px;
  font-size: 16px;
  color: #007bff;
}

.social-login {
  text-align: center;
  margin-top: 20px;
}

.social-login p {
  font-size: 18px;
  margin-bottom: 10px;
}

.social-icons button {
  background-color: transparent;
  border: none;
  font-size: 24px;
  margin: 0 10px;
  cursor: pointer;
  color: #007bff;
  transition: color 0.3s;
}

.social-icons button:hover {
  color: #0056b3;
}

.password-input {
  position: relative;
}

.password-input input[type="password"] {
  width: 100%;
  padding: 12px;
  border: 1px solid #ccc;
  border-radius: 5px;
  outline: none;
  font-size: 16px;
  padding-right: 40px;
}

.password-input i {
  position: absolute;
  top: 50%;
  right: 10px;
  transform: translateY(-50%);
  cursor: pointer;
}

.captcha-input {
  display: flex;
  align-items: center;
}

.captcha-input input[type="text"] {
  flex: 1;
  padding: 12px;
  border: 1px solid #ccc;
  border-radius: 5px;
  outline: none;
  font-size: 16px;
}

.captcha-input .captcha-image {
  margin-left: 10px;
  cursor: pointer;
  width: 120px; /* 调整验证码图片的大小 */
  height: 40px; /* 调整验证码图片的大小 */
}

.img_github {
  width: 40px;
  height: 40px;
  margin-right: 10px;
  border-radius: 50%;
}

.img_gitee {
  width: 40px;
  height: 40px;
  margin-right: 10px;
}

.reset_but {
/* 在你的CSS文件中定义一个样式类，例如.custom-button */
  border: none;         /* 去掉边框 */
  background: none;     /* 去掉背景颜色 */
  padding: 0;           /* 去掉内边距 */
  cursor: pointer;      /* 添加鼠标指针样式，让按钮看起来可以点击 */
  /* 可以继续添加其他样式，如颜色、字体大小等 */
  color: #007bff;
  font-size: 16px;
  font-weight: bold;
}
</style>
