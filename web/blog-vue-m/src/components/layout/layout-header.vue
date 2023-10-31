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
          <span class="el-dropdown-link" >
            主页<i class="el-icon-arrow-down el-icon--right"></i>
          </span>
          <el-dropdown-menu slot="dropdown" >
            <el-dropdown-item><router-link to="/user">主页</router-link></el-dropdown-item>
            <el-dropdown-item><router-link to="/user/write_article">写文章</router-link></el-dropdown-item>
            <el-dropdown-item command="update_pwd">修改密码</el-dropdown-item>
            <el-dropdown-item command="update_motto">修改格言</el-dropdown-item>
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
          <el-input v-model="form.email" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码" :label-width="formLabelWidth">
          <el-input v-model="form.password" placeholder="请输入密码" show-password autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="确认密码" :label-width="formLabelWidth">
          <el-input v-model="form.ensure_password" placeholder="请再次输入密码" show-password autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" @click="fetchToUpdatePwd">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog title="修改格言" :visible.sync="dialogFormVisible2">
      <el-form :model="form">
        <el-form-item label="注册邮箱" :label-width="formLabelWidth">
          <el-input v-model="form.email" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="格言" :label-width="formLabelWidth">
          <el-input v-model="form.motto" placeholder="请输入新的格言" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible2 = false">取 消</el-button>
        <el-button type="primary" @click="fetchToUpdatePwd">确 定</el-button>
      </div>
    </el-dialog>

    <el-dialog title="修改社交帐号" :visible.sync="dialogFormVisible3">
      <el-form :model="form">
        <el-form-item label="注册邮箱" :label-width="formLabelWidth">
          <el-input v-model="form.email" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="QQ" :label-width="formLabelWidth">
          <el-input v-model="form.social_account_qq" placeholder="请输入新的QQ帐号" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="GitHub" :label-width="formLabelWidth">
          <el-input v-model="form.social_account_github" placeholder="请输入新的GitHub帐号" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="Gitee" :label-width="formLabelWidth">
          <el-input v-model="form.social_account_gitee" placeholder="请输入新的Gitee帐号" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="CSDN" :label-width="formLabelWidth">
          <el-input v-model="form.social_account_csdn" placeholder="请输入新的CSDN帐号" autocomplete="off"></el-input>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="dialogFormVisible3 = false">取 消</el-button>
        <el-button type="primary" @click="fetchToUpdatePwd">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import HeaderSearch from '@/components/header-search'
import {fetchCategory, fetchIsLogin, fetchLogOut} from '../../api'

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
      form: {
        email: localStorage.getItem('email') ? localStorage.getItem('email') : '请输入邮箱！',
        type: [],
        password: '',
        ensure_password: '',
        motto: '',
        social_account_qq: '',
        social_account_github: '',
        social_account_gitee: '',
        social_account_csdn: '',
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
      let data = {
        token: localStorage.getItem('token')
      }
      fetchIsLogin(data).then(res => {
        this.is_login = res.data.data.is_login
        console.log("-----", this.is_login, res.data.data.is_login)
      }).catch(err => {
        console.log(err)
      })
    },
    fetchTo(command) {
      if (command === 'logout') {
        let data = {
          token: localStorage.getItem('token')
        }
        fetchLogOut(data).then(res => {
          console.log(res)
          localStorage.removeItem('token')
          localStorage.removeItem('is_login')
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
      }
    },
    fetchToUpdatePwd() {
      let data = {
        email: this.form.email,
        password: this.form.password,
        token: localStorage.getItem('token')
      }
      // fetchUpdatePwd(data).then(res => {
      //   console.log(res)
      //   this.dialogFormVisible = false
      // }).catch(err => {
      //   console.log(err)
      // })
      this.dialogFormVisible = false
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
