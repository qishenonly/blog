<template>
  <div id="layout-header" :class="{'fixed':fixed,'hidden':hidden}" @click.stop="mobileShow=false">
    <div class="site-logo">
      <router-link to="/">
        <img src="@/assets/site-logo.svg" alt="">
        <p class="site-name">Gblog</p>
      </router-link>
    </div>
    <div class="menus-btn" @click.stop="mobileShow=!mobileShow">
      Menus
    </div>
    <div class="site-menus" :class="{'mobileShow':mobileShow}" @click.stop="mobileShow=!mobileShow">
      <div class="menu-item header-search">
        <header-search/>
      </div>
      <div class="menu-item">
        <router-link to="/">首页</router-link>
      </div>
      <div class="menu-item hasChild">
        <a href="#">文章</a>
        <div class="childMenu" v-if="category.length">
          <div class="sub-menu" v-for="item in category" :key="item.title">
            <router-link :to="`/category/${item.title}`">{{ item.title }}</router-link>
          </div>
        </div>
      </div>
      <div class="menu-item">
        <router-link to="/friend">友链</router-link>
      </div>
      <div class="menu-item">
        <router-link to="/about">关于</router-link>
      </div>
      <div v-if="this.is_login" class="menu-item">
        <!--        <router-link to="/user">主页</router-link>-->
        <el-dropdown @command="fetchTo">
          <span class="el-dropdown-link">
            主页<i class="el-icon-arrow-down el-icon--right"></i>
          </span>
          <el-dropdown-menu slot="dropdown">
            <el-dropdown-item>
              <router-link to="/user">主页</router-link>
            </el-dropdown-item>
            <el-dropdown-item command="pathToWriteArticle">写文章</el-dropdown-item>
            <el-dropdown-item command="update_pwd">修改密码</el-dropdown-item>
<!--            <el-dropdown-item command="update_email">修改邮箱</el-dropdown-item>-->
            <el-dropdown-item command="update_motto">修改个性签名</el-dropdown-item>
            <el-dropdown-item command="update_social_account">修改社交帐号</el-dropdown-item>
            <el-dropdown-item divided command="logout">退出登录</el-dropdown-item>
          </el-dropdown-menu>
        </el-dropdown>
      </div>
      <div v-else class="menu-item">
        <router-link to="/auth/login">登录</router-link>
      </div>
    </div>
    <el-dialog title="修改密码" :visible.sync="dialogFormVisible">
      <el-form :model="form">
        <el-form-item label="注册邮箱" :label-width="formLabelWidth">
          <el-input placeholder="请输入内容" v-model="form.email" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="密码" :label-width="formLabelWidth">
          <el-input v-model="form.password" placeholder="请输入密码" show-password autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="确认密码" :label-width="formLabelWidth">
          <el-input v-model="form.ensure_password" placeholder="请再次输入密码" show-password
                    autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="fetchToUpdatePwd">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog title="修改个性签名" :visible.sync="dialogFormVisible2">
      <el-form :model="form">
        <el-form-item label="注册邮箱" :label-width="formLabelWidth">
          <el-input placeholder="请输入内容" v-model="form.email" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="个性签名" :label-width="formLabelWidth">
          <el-input v-model="form.motto" placeholder="请输入新的个性签名" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible2 = false">取 消</el-button>
        <el-button type="primary" @click="fetchToUpdateMotto">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog title="修改邮箱" :visible.sync="dialogFormVisible4">
      <el-form :model="form">
        <el-form-item label="注册邮箱" :label-width="formLabelWidth">
          <el-input placeholder="请输入内容" v-model="form.email" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="新的邮箱" :label-width="formLabelWidth">
          <el-input v-model="form.new_email" placeholder="请输入新的邮箱" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible4 = false">取 消</el-button>
        <el-button type="primary" @click="fetchToUpdateEmail">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog title="修改社交帐号" :visible.sync="dialogFormVisible3">
      <el-form :model="form">
        <el-form-item label="注册邮箱" :label-width="formLabelWidth">
          <el-input placeholder="请输入内容" v-model="form.email" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item label="选择帐号" :label-width="formLabelWidth">
          <el-select v-model="form.social_account_type" placeholder="请选择需要修改的帐号">
            <el-option label="QQ" value="QQ"></el-option>
            <el-option label="GitHub" value="GitHub"></el-option>
            <el-option label="Gitee" value="Gitee"></el-option>
            <el-option label="CSDN" value="CSDN"></el-option>
          </el-select>
        </el-form-item>
        <el-form-item v-if="form.social_account_type==='QQ'"
                      label="旧账户" :label-width="formLabelWidth">
          <el-input placeholder="请输入内容" v-model="form.qq" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item v-if="form.social_account_type==='QQ'" v-model="form.social_account_qq"
                      label="新的QQ帐号" :label-width="formLabelWidth">
          <el-input v-model="form.social_account_qq" placeholder="请输入新的QQ帐号"></el-input>
        </el-form-item>
        <el-form-item v-if="form.social_account_type==='GitHub'"
                      label="旧账户" :label-width="formLabelWidth">
          <el-input placeholder="请输入内容" v-model="form.github" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item v-if="form.social_account_type==='GitHub'" v-model="form.social_account_github"
                      label="新的GitHub账户" :label-width="formLabelWidth">
          <el-input v-model="form.social_account_github" placeholder="请输入新的GitHub账户地址"></el-input>
        </el-form-item>
        <el-form-item v-if="form.social_account_type==='Gitee'"
                      label="旧账户" :label-width="formLabelWidth">
          <el-input placeholder="请输入内容" v-model="form.gitee" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item v-if="form.social_account_type==='Gitee'" v-model="form.social_account_gitee"
                      label="新的Gitee账户" :label-width="formLabelWidth">
          <el-input v-model="form.social_account_gitee" placeholder="请输入新的Gitee账户地址"></el-input>
        </el-form-item>
        <el-form-item v-if="form.social_account_type==='CSDN'"
                      label="旧账户" :label-width="formLabelWidth">
          <el-input placeholder="请输入内容" v-model="form.csdn" :disabled="true"></el-input>
        </el-form-item>
        <el-form-item v-if="form.social_account_type==='CSDN'" v-model="form.social_account_csdn"
                      label="新的CSDN账户" :label-width="formLabelWidth">
          <el-input v-model="form.social_account_csdn" placeholder="请输入新的CSDN账户地址"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible3 = false">取 消</el-button>
        <el-button type="primary" @click="fetchToUpdateSocialAccount">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import HeaderSearch from '@/components/header-search'
import {
  fetchCategory,
  fetchIsLogin,
  fetchLogOut,
  fetchToUpdateMotto,
  fetchToUpdateSocialAccountQQ,
  fetchToUpdateSocialAccountGithub,
  fetchToUpdateSocialAccountGitee,
  fetchToUpdateSocialAccountCSDN,
  fetchToUpdateEmail,
  fetchUpdatePwd
} from '../../api'

export default {
  name: "layout-header",
  components: {HeaderSearch},
  data() {
    return {
      lastScrollTop: 0,
      fixed: false,
      hidden: false,
      category: [],
      mobileShow: false,
      is_login: false,
      dialogTableVisible: false,
      dialogFormVisible: false,
      dialogFormVisible2: false,
      dialogFormVisible3: false,
      dialogFormVisible4: false,
      form: {
        email: JSON.parse(localStorage.getItem('userinfo')).email ? JSON.parse(localStorage.getItem('userinfo')).email : '请输入邮箱！',
        qq: JSON.parse(localStorage.getItem('userinfo')).qq ? JSON.parse(localStorage.getItem('userinfo')).qq : '数据为空！',
        github: JSON.parse(localStorage.getItem('userinfo')).github ? JSON.parse(localStorage.getItem('userinfo')).github : '数据为空！',
        gitee: JSON.parse(localStorage.getItem('userinfo')).gitee ? JSON.parse(localStorage.getItem('userinfo')).gitee : '数据为空！',
        csdn: JSON.parse(localStorage.getItem('userinfo')).csdn ? JSON.parse(localStorage.getItem('userinfo')).csdn : '数据为空！',
        type: [],
        password: '',
        ensure_password: '',
        motto: '',
        social_account_qq: '',
        social_account_github: '',
        social_account_gitee: '',
        social_account_csdn: '',
        social_account_type: '',
        social_account: ''
      },
      formLabelWidth: '120px'
    }
  },
  mounted() {
    window.addEventListener('scroll', this.watchScroll)
    this.fetchIsLogin()
    localStorage.getItem('is_login') ? this.is_login = true : this.is_login = false

    this.fetchCategory()
  },
  created() {
    this.fetchCategory()
  },
  beforeDestroy() {
    window.removeEventListener("scroll", this.watchScroll)
  },
  methods: {
    watchScroll() {
      let scrollTop = window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop
      if (scrollTop === 0) {
        this.fixed = false;
      } else if (scrollTop >= this.lastScrollTop) {
        this.fixed = false;
        this.hidden = true;
      } else {
        this.fixed = true
        this.hidden = false
      }
      this.lastScrollTop = scrollTop
    },
    fetchCategory() {
      fetchCategory().then(res => {
        this.category = res.data
      }).catch(err => {
        console.log(err)
      })
    },
    fetchIsLogin() {
        let token = localStorage.getItem('token')
      fetchIsLogin(token).then(res => {
        this.is_login = res.data.data.is_login
        console.log("-----", this.is_login, res.data.data.is_login)
      }).catch(err => {
        console.log(err)
      })
    },
    fetchTo(command) {
      if (JSON.parse(localStorage.getItem('userinfo')).email === '请输入邮箱！') {
        window.location.reload()
      }
      if (command === 'logout') {
        let token = localStorage.getItem('token')
        fetchLogOut(token).then(res => {
          console.log(res)
          localStorage.removeItem('token')
          localStorage.removeItem('is_login')
          localStorage.removeItem('userinfo')
          this.is_login = false
          this.$router.push({path: '/'})
        }).catch(err => {
          console.log(err)
        })
      } else if (command === 'update_pwd') {
        this.dialogFormVisible = true
      } else if (command === 'update_motto') {
        this.dialogFormVisible2 = true
      } else if (command === 'update_social_account') {
        this.dialogFormVisible3 = true
      } else if (command === 'update_email') {
        this.dialogFormVisible4 = true
      } else if (command === 'pathToWriteArticle') {
        this.$router.push({
          name: 'write_article',
          params: {
            email: this.form.email,
          }
        });
      }
    },
    fetchToUpdatePwd() {
      // 判断密码是否过于简单，要求八位以上，且包含数字和字母
      let reg = /^(?![0-9]+$)(?![a-zA-Z]+$)[0-9A-Za-z]{8,}$/
      if (!reg.test(this.form.password)) {
        this.$message({
          message: '密码过于简单，要求八位以上且包含数字和字母！',
          type: 'warning'
        })
        return
      }

      if (this.form.password !== this.form.ensure_password) {
        this.$message({
          message: '两次输入的密码不一致，请重新输入！',
          type: 'warning'
        })
        return
      }

      let data = {
        new_password: this.form.password,
      }
      fetchUpdatePwd(data, localStorage.getItem('token')).then(res => {
        console.log(res)
        this.dialogFormVisible = false
        if (res.data.code === 200) {
          this.$message({
            message: '密码修改成功！',
            type: 'success'
          })
        } else {
          this.$message({
            message: '密码修改失败！',
            type: 'error'
          })
        }
      }).catch(err => {
        console.log(err)
      })
      this.dialogFormVisible = false
    },
    fetchToUpdateMotto() {
      if (this.form.motto === "") {
        this.$message({
          message: '个性签名不能为空！',
          type: 'warning'
        })
        return
      }

      let data = {
        motto: this.form.motto,
      }
      fetchToUpdateMotto(data, localStorage.getItem('token')).then(res => {
        console.log(res)
        this.dialogFormVisible2 = false
        if (res.data.code === 200) {
          this.$message({
            message: '个性签名修改成功！',
            type: 'success'
          })
        } else {
          this.$message({
            message: '个性签名修改失败！',
            type: 'error'
          })
        }
      }).catch(err => {
        console.log(err)
      })
      this.dialogFormVisible2 = false
    },
    fetchToUpdateEmail() {
      if (this.form.email === "") {
        this.$message({
          message: '邮箱不能为空！',
          type: 'warning'
        })
        return
      }
      let reg = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(.[a-zA-Z0-9_-])+/
      if (!reg.test(this.form.email)) {
        this.$message({
          message: '邮箱格式不正确！',
          type: 'warning'
        })
        return
      }

      let data = {
        email: this.form.email,
      }
      fetchToUpdateEmail(data, localStorage.getItem('token')).then(res => {
        console.log(res)
        this.dialogFormVisible4 = false
        if (res.data.code === 200) {
          this.$message({
            message: '邮箱修改成功！',
            type: 'success'
          })
          localStorage.setItem('email', this.form.email)
        } else {
          this.$message({
            message: '邮箱修改失败！',
            type: 'error'
          })
        }
      }).catch(err => {
        console.log(err)
      })
      this.dialogFormVisible4 = false
    },
    fetchToUpdateSocialAccount() {
      if (this.form.social_account_type === "") {
        this.$message({
          message: '请选择社交账号类型！',
          type: 'warning'
        })
        return
      }
      if (this.form.social_account_type === "QQ") {
        let reg = /^[1-9][0-9]{4,}$/
        if (!reg.test(this.form.social_account_qq)) {
          this.$message({
            message: 'QQ号格式不正确！',
            type: 'warning'
          })
          return
        }
        let data = {
          qq: this.form.social_account_qq,
        }
        this.methodToSelectSocialAccountToServer(data, 'QQ', fetchToUpdateSocialAccountQQ)
      } else if (this.form.social_account_type === "GitHub") {
        // 判断是不是github用户页面的url
        let reg = /^https:\/\/github.com\/[a-zA-Z0-9_-]+$/;
        if (!reg.test(this.form.social_account_github)) {
          this.$message({
            message: 'GitHub用户url格式不正确！',
            type: 'warning'
          })
          return
        }
        let data = {
          github: this.form.social_account_github,
        }
        this.methodToSelectSocialAccountToServer(data, 'GitHub', fetchToUpdateSocialAccountGithub)
      } else if (this.form.social_account_type === "Gitee") {
        // 判断是不是gitee用户页面的url
        let reg = /^https:\/\/gitee.com\/[a-zA-Z0-9_-]+$/;
        if (!reg.test(this.form.social_account_gitee)) {
          this.$message({
            message: 'Gitee用户url格式不正确！',
            type: 'warning'
          })
          return
        }
        let data = {
          gitee: this.form.social_account_gitee,
        }
        this.methodToSelectSocialAccountToServer(data, 'Gitee', fetchToUpdateSocialAccountGitee)
      } else if (this.form.social_account_type === "CSDN") {
        // 判断是不是csdn用户页面的url
        let reg = /^https:\/\/blog.csdn.net\/[a-zA-Z0-9_-]+$/;
        if (!reg.test(this.form.social_account_csdn)) {
          this.$message({
            message: 'CSDN用户url格式不正确！',
            type: 'warning'
          })
          return
        }
        let data = {
          csdn: this.form.social_account_csdn,
        }
        this.methodToSelectSocialAccountToServer(data, 'CSDN', fetchToUpdateSocialAccountCSDN)
      }
    },
    methodToSelectSocialAccountToServer(data, account_type, func) {
      func(data, localStorage.getItem('token')).then(res => {
        console.log(res)
        this.dialogFormVisible3 = false
        if (res.data.code === 200) {
          this.$message({
            message: account_type + '社交账号修改成功！',
            type: 'success'
          })
        } else {
          this.$message({
            message: account_type + '社交账号修改失败！',
            type: 'error'
          })
        }
      }).catch(err => {
        console.log(err)
      })
    }
  }
}
</script>

<style scoped lang="less">
#layout-header {
  position: fixed;
  top: 0;
  z-index: 9;
  width: 100%;
  height: 80px;
  padding: 0 80px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  transition: .3s all ease;
  -webkit-transition: .3s all ease;
  -moz-transition: .3s all linear;
  -o-transition: .3s all ease;
  -ms-transition: .3s all ease;

  &.hidden {
    top: -100px;
  }

  &.fixed {
    background-color: #FFFFFF;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
  }
}

.site-logo {
  text-align: center;

  img {
    width: 60px;
    height: 60px;
  }

  p.site-name {
    font-size: 15px;
    font-weight: bold;
    position: relative;
    top: -10px;
  }
}

.menus-btn {
  display: none;
  visibility: hidden;
}

.site-menus {
  display: flex;
  align-items: center;

  .menu-item {
    min-width: 60px;
    height: 50px;
    line-height: 50px;
    text-align: center;
    position: relative;

    a {
      padding: 12px 10px;
      color: #545454;
      font-weight: 500;
      font-size: 16px;

      &:hover {
        color: #ff6d6d;
      }
    }

    &:not(:last-child) {
      margin-right: 15px;
    }

    &.hasChild:hover .childMenu {
      opacity: 1;
      visibility: visible;
      transform: translateY(-5px);
    }
  }

  .childMenu {
    width: 130px;
    background-color: #FDFDFD;
    border-radius: 3px;
    border: 1px solid #ddd;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);
    position: absolute;
    top: 55px;
    z-index: 2;
    opacity: 0;
    visibility: hidden;
    transition: .7s all ease;
    -webkit-transition: .6s all ease;
    -moz-transition: .6s all linear;
    -o-transition: .6s all ease;
    -ms-transition: .6s all ease;

    &:before, &:after {
      content: '';
      position: absolute;
      width: 0;
      height: 0;
      border-left: 6px solid transparent;
      border-right: 6px solid transparent;
      border-bottom: 8px solid white;
      top: -8px;
      left: 20px;
    }

    &:before {
      top: -9px;
      border-bottom: 8px solid #ddd;
    }

    .sub-menu {
      height: 40px;
      line-height: 40px;
      position: relative;

      &:not(:last-child):after {
        /*position: absolute;*/
        content: '';
        width: 50%;
        height: 1px;
        color: #ff6d6d;
        bottom: 0;
        left: 25%;
        z-index: 99;
      }
    }
  }
}

@media (max-width: 960px) {
  #layout-header {
    padding: 0 20px;
  }
}

@media (max-width: 600px) {
  #layout-header {
    padding: 0 10px;
  }

  .menus-btn {
    display: block;
    visibility: visible;
  }

  .site-menus {
    position: absolute;
    display: none;
    visibility: hidden;
    background-color: #F9F9F9;
    width: 100%;
    left: 0;
    top: 80px;
    z-index: -9;
    box-shadow: 0 2px 12px 0 rgba(0, 0, 0, 0.1);

    .menu-item {
      position: relative;
      height: unset;

      &:not(:last-child) {
        margin-right: 0;
      }
    }

    .childMenu {
      position: relative;
      width: 100%;
      top: 0;
      background-color: #F3F3F3;
      opacity: 1;
      visibility: visible;
      border: none;
      box-shadow: none;

      &:before, &:after {
        content: '';
        position: relative;
        width: 0;
        height: 0;
        border-left: 0;
        border-right: 0;
        border-bottom: 0;
      }
    }
  }

  .site-menus.mobileShow {
    display: inline-block;
    visibility: visible;
    z-index: 99;
  }
}
</style>
