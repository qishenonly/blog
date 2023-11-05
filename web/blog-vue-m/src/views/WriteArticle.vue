<template>
  <div class="container">
    <div class="write_container ">
      <div class="title_container">
        <span class="span_title">文章标题:</span>
        <el-input
            class="input_title"
            v-model="title"
            placeholder="请输入文章标题"
            clearable
        ></el-input>
      </div>
      <div class="article_introduction_container">
        <span class="span_introduction">文章简介:</span>
        <el-input
            class="input_introduction"
            v-model="introduction"
            type="textarea"
            :rows="3"
            placeholder="请输入文章简介"
            clearable
        ></el-input>
      </div>
      <div class="category_container">
        <span class="span_category">文章分类:</span>
        <el-select
            v-model="category"
            placeholder="请选择文章分类"
            clearable
        >
          <el-option
              v-for="item in categoryList"
              :key="item.id"
              :label="item.name"
              :value="item.id"
          ></el-option>
        </el-select>
      </div>
      <div class="article_cover_container">
        <span class="span_cover">文章封面:</span>
        <el-upload
            class="avatar-uploader"
            :on-preview="handlePreview"
            :on-remove="handleRemove"
            :http-request="onUploadCoverHandler"
            :file-list="coverList"
            list-type="picture">
          <el-button >选择图片上传<i class="el-icon-upload el-icon--right"></i></el-button>
          <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
        </el-upload>
      </div>
      <div class="edit_container">
        <span class="span_content">文章内容:</span>
        <!--  新增时输入 -->
        <quill-editor
            v-model="content"
            ref="myQuillEditor"
            :options="editorOption"
            @blur="onEditorBlur($event)"
            @focus="onEditorFocus($event)"
            @change="onEditorChange($event)"
        >
        </quill-editor>

        <el-upload
            class="avatar-uploader"
            :show-file-list="false"
            :http-request="onUploadHandler"
        >
        </el-upload>
        <!-- 从数据库读取展示 -->
        <div v-html="str" class="ql-editor">
          {{ str }}
        </div></div>
    </div>
  </div>
</template>

<script>
import { quillEditor } from 'vue-quill-editor' // 调用编辑器
import 'quill/dist/quill.core.css'
import 'quill/dist/quill.snow.css'
import 'quill/dist/quill.bubble.css'
import {fetchUploadArticleImage} from '@/api'
import { Quill } from 'vue-quill-editor'
import { ImageDrop } from 'quill-image-drop-module'
import ImageResize from 'quill-image-resize-module'
Quill.register('modules/imageDrop', ImageDrop)
Quill.register('modules/imageResize', ImageResize)
// 自定义字体大小
const Size = Quill.import('attributors/style/size')
Size.whitelist = ['10px', '12px', '16px', '18px', '20px', '30px', '32px']
Quill.register(Size, true)


// 自定义字体类型
var fonts = ['SimSun', 'SimHei', 'Microsoft-YaHei', 'KaiTi', 'FangSong', 'Arial', 'sans-serif']
var Font = Quill.import('formats/font')
Font.whitelist = fonts
Quill.register(Font, true)

export default {
  name: 'Editor',
  components: {
    quillEditor
  },
  data () {
    return {
      content: ``,
      coverList: [],
      title: '',
      introduction: '',
      category: '',
      categoryList: [],
      editorOption: {
        placeholder: '请在这里输入',
        modules: {
          toolbar: {
            container: [
              ['bold', 'italic', 'underline', 'strike'], // 加粗，斜体，下划线，删除线
              ['blockquote', 'code-block'], // 引用，代码块
              [{ 'header': 1 }, { 'header': 2 }], // 标题，键值对的形式；1、2表示字体大小
              [{ 'list': 'ordered' }, { 'list': 'bullet' }], // 列表
              [{ 'script': 'sub' }, { 'script': 'super' }], // 上下标
              [{ 'indent': '-1' }, { 'indent': '+1' }], // 缩进
              [{ 'direction': 'rtl' }], // 文本方向
              [{ 'size': ['10px', '12px', '16px', '18px', '20px', '30px', '32px'] }], // 字体大小
              [{ 'header': [1, 2, 3, 4, 5, 6, false] }], // 几级标题
              [{ 'color': [] }, { 'background': [] }], // 字体颜色，字体背景颜色
              // [{ 'font': [false, 'heiti', 'songti', 'kaiti', 'lishu', 'arial', 'monospace'] }], // 字体
              [{ 'font': ['SimSun', 'SimHei', 'Microsoft-YaHei', 'KaiTi', 'FangSong', 'Arial', 'sans-serif'] }], // 字体
              [{ 'align': [] }], // 对齐方式
              ['clean'], // 清除字体样式
              ['image', 'video'], // 上传图片、上传视频
              // 调整图片大小
            ],
            handlers: {
              'image': function (value) {
                console.log('value', value)
                if (value) { // value === true
                  console.log('上传图片', value, this.quill)
                  document.querySelector('.avatar-uploader input').click()
                } else {
                  console.log('取消上传图片', value)
                  this.quill.format('image', false)
                }
              }
            }
          },
          imageDrop: true, // 拖动加载图片组件。
          imageResize: { //调整大小组件。
            displayStyles: {
              backgroundColor: 'black',
              border: 'none',
              color: 'white'
            },
            modules: [ 'Resize', 'DisplaySize', 'Toolbar' ]
          },
        },
        theme: "snow", //主题 snow/bubble
        syntax: true, //语法检测
      }
    }
  },
  computed: {
    // 当前富文本实例
    editor () {
      return this.$refs.myQuillEditor.quill
    }
  },
  mounted () {
    const content = ''// 请求后台返回的内容字符串
    this.str = this.escapeStringHTML(content)
  },
  methods: {
    onEditorReady (editor) { // 准备编辑器
    },
    onEditorBlur () {}, // 失去焦点事件
    onEditorFocus () {}, // 获得焦点事件
    onEditorChange () {
      this.$emit('change', this.escapeStringHTML(this.content))
    }, // 内容改变事件
    // 转码
    escapeStringHTML (str) {
      str = str.replace(/&lt;/g, '<')
      str = str.replace(/&gt;/g, '>')
      return str
    },
    async onUploadHandler(e) {
      fetchUploadArticleImage(e.file).then(res => {
        console.log("上传图片成功", res)
        let imageUrl = res.data.data.path
        // 获取光标所在位置
        let quill = this.$refs.myQuillEditor.$refs.editor.__quill
        // quill = this.$refs.myQuillEditor.quill
        let length = quill.getSelection().index
        // 插入图片
        console.log("imgUrl", imageUrl)
        quill.insertEmbed(length, 'image', imageUrl)
        // 调整光标到最后
        quill.setSelection(length + 1)
        // this.content += imageUrl
        console.log("content", this.content)
      }).catch(err => {
        console.log("上传图片失败", err)

      })

    },
    handleRemove(file, fileList) {
      console.log("--------------",file, fileList);
    },
    handlePreview(file) {
      console.log("--------------",file);
    },
    async onUploadCoverHandler(e) {
      console.log("上传封面", e, e.file, this.coverList.length)
      if (this.coverList.length > 0) {
        this.$message.error('只能上传一张封面')
        this.coverList.shift()
      }
      let coverData = {
        name : e.file.name,
        url : 'xxxx'
      }
      this.coverList.push(coverData)

    },

  }
}
</script>

<style>
/*.write_container {*/
/*  !* 设置容器的宽度和高度 *!*/
/*  font-family: Arial, sans-serif; !* 字体，可以根据需要修改 *!*/
/*  color: #333; !* 文字颜色，可以根据需要修改 *!*/

/*  !* 设置容器的文本对齐方式 *!*/
/*  text-align: left; !* 文本对齐方式，可以根据需要修改 *!*/

/*  !* 设置容器的最大宽度 *!*/
/*  max-width: 800px; !* 最大宽度，可以根据需要调整 *!*/

/*  !* 禁用用户选择内容 *!*/
/*  user-select: none;*/

/*  !* 禁用容器内文字内容被拖动 *!*/
/*  -webkit-user-drag: none;*/
/*}*/

/* 鼠标悬停时的样式 */
.write_container:hover {
  cursor: pointer; /* 鼠标悬停时的光标样式 */
}

.container {
  width: 100%;
  margin: auto;
  margin-top: 5%;
  max-width: 800px;
  font-family: Arial, sans-serif;
  color: #333;
  text-align: left;
  user-select: none;
  -webkit-user-drag: none;
  display: flex;
}

.edit_container, .title_container, .category_container, .article_cover_container, .article_introduction_container {
  display: flex;
  margin-top: 2%;
}

.span_content {
  width: 18%;
  margin-top: 0.5%;
  margin-right: 1%;
  overflow: auto;
}

.span_title {
  width: 12%;
  margin-top: 0.5%;
  margin-right: 1%;
  overflow: auto;
}

.span_category {
  width: 10.5%;
  margin-top: 0.5%;
  margin-right: 1%;
  overflow: auto;
}

.span_introduction {
  width: 12%;
  margin-top: 0.5%;
  margin-right: 1%;
  overflow: auto;
}

.span_cover {
  width: 10.5%;
  margin-top: 0.5%;
  margin-right: 1%;
  overflow: auto;
}


/* 给文本内容加高度，滚动 */
.quill-editor /deep/ .ql-container {
   min-height: 800px;
}

.ql-container {
  min-height: 400px;
}

.ql-snow .ql-tooltip [data-mode="link"]::before {
  content: "请输入链接地址:";
  left: 0;
}

.ql-snow .ql-tooltip.ql-editing {
  left: 0 !important;
}

.ql-snow .ql-tooltip {
  left: 0 !important;
}

.ql-snow .ql-editor {
  max-height: 500px;
}

.ql-snow .ql-tooltip.ql-editing a.ql-action::after {
  border-right: 0px;
  content: "保存";
  padding-right: 0px;
}

.ql-snow .ql-tooltip[data-mode="video"]::before {
  content: "请输入视频地址:";
}

.ql-snow .ql-picker.ql-size .ql-picker-label::before, .ql-snow .ql-picker.ql-size .ql-picker-item::before {
  content: "14px" !important;
  font-size: 14px;
}

.ql-snow .ql-picker.ql-size .ql-picker-label[data-value="10px"]::before, .ql-snow .ql-picker.ql-size .ql-picker-item[data-value="10px"]::before {
  content: "10px" !important;
  font-size: 10px;
}

.ql-snow .ql-picker.ql-size .ql-picker-label[data-value="12px"]::before, .ql-snow .ql-picker.ql-size .ql-picker-item[data-value="12px"]::before {
  content: "12px" !important;
  font-size: 12px;
}

.ql-snow .ql-picker.ql-size .ql-picker-label[data-value="16px"]::before, .ql-snow .ql-picker.ql-size .ql-picker-item[data-value="16px"]::before {
  content: "16px" !important;
  font-size: 16px;
}

.ql-snow .ql-picker.ql-size .ql-picker-label[data-value="18px"]::before, .ql-snow .ql-picker.ql-size .ql-picker-item[data-value="18px"]::before {
  content: "18px" !important;
  font-size: 18px;
}

.ql-snow .ql-picker.ql-size .ql-picker-label[data-value="20px"]::before, .ql-snow .ql-picker.ql-size .ql-picker-item[data-value="20px"]::before {
  content: "20px" !important;
  font-size: 20px;
}

.ql-snow .ql-picker.ql-size .ql-picker-label[data-value="30px"]::before, .ql-snow .ql-picker.ql-size .ql-picker-item[data-value="30px"]::before {
  content: "30px" !important;
  font-size: 30px;
}

.ql-snow .ql-picker.ql-size .ql-picker-label[data-value="32px"]::before, .ql-snow .ql-picker.ql-size .ql-picker-item[data-value="32px"]::before {
  content: "32px" !important;
  font-size: 32px;
}

.ql-snow .ql-picker.ql-header .ql-picker-label::before, .ql-snow .ql-picker.ql-header .ql-picker-item::before {
  content: "文本" !important;
}

.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="1"]::before, .ql-snow .ql-picker.ql-header .ql-picker-item[data-value="1"]::before {
  content: "标题1" !important;
}

.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="2"]::before, .ql-snow .ql-picker.ql-header .ql-picker-item[data-value="2"]::before {
  content: "标题2" !important;
}

.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="3"]::before, .ql-snow .ql-picker.ql-header .ql-picker-item[data-value="3"]::before {
  content: "标题3" !important;
}

.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="4"]::before, .ql-snow .ql-picker.ql-header .ql-picker-item[data-value="4"]::before {
  content: "标题4" !important;
}

.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="5"]::before, .ql-snow .ql-picker.ql-header .ql-picker-item[data-value="5"]::before {
  content: "标题5" !important;
}

.ql-snow .ql-picker.ql-header .ql-picker-label[data-value="6"]::before, .ql-snow .ql-picker.ql-header .ql-picker-item[data-value="6"]::before {
  content: "标题6" !important;
}

.ql-snow .ql-picker.ql-font .ql-picker-label::before, .ql-snow .ql-picker.ql-font .ql-picker-item::before {
  content: "标准字体" !important;
}

.ql-snow .ql-picker.ql-font .ql-picker-label[data-value="serif"]::before, .ql-snow .ql-picker.ql-font .ql-picker-item[data-value="serif"]::before {
  content: "衬线字体" !important;
}

.ql-snow .ql-picker.ql-font .ql-picker-label[data-value="monospace"]::before, .ql-snow .ql-picker.ql-font .ql-picker-item[data-value="monospace"]::before {
  content: "等宽字体" !important;
}

.ql-snow .ql-picker.ql-font .ql-picker-label[data-value="SimSun"]::before, .ql-snow .ql-picker.ql-font .ql-picker-item[data-value="SimSun"]::before {
  content: "宋体" !important;
  font-family: "SimSun";
}

.ql-snow .ql-picker.ql-font .ql-picker-label[data-value="SimHei"]::before, .ql-snow .ql-picker.ql-font .ql-picker-item[data-value="SimHei"]::before {
  content: "黑体" !important;
  font-family: "SimHei";
}

.ql-snow .ql-picker.ql-font .ql-picker-label[data-value="Microsoft-YaHei"]::before, .ql-snow .ql-picker.ql-font .ql-picker-item[data-value="Microsoft-YaHei"]::before {
  content: "微软雅黑" !important;
  font-family: "Microsoft YaHei";
}

.ql-snow .ql-picker.ql-font .ql-picker-label[data-value="KaiTi"]::before, .ql-snow .ql-picker.ql-font .ql-picker-item[data-value="KaiTi"]::before {
  content: "楷体" !important;
  font-family: "KaiTi";
}

.ql-snow .ql-picker.ql-font .ql-picker-label[data-value="FangSong"]::before, .ql-snow .ql-picker.ql-font .ql-picker-item[data-value="FangSong"]::before {
  content: "仿宋" !important;
  font-family: "FangSong",serif;
}

.ql-snow .ql-picker.ql-font .ql-picker-label[data-value="Arial"]::before, .ql-snow .ql-picker.ql-font .ql-picker-item[data-value="Arial"]::before {
  content: "Arial" !important;
  font-family: "Arial";
}

.ql-snow .ql-picker.ql-font .ql-picker-label[data-value="Times-New-Roman"]::before, .ql-snow .ql-picker.ql-font .ql-picker-item[data-value="Times-New-Roman"]::before {
  content: "Times New Roman" !important;
  font-family: "Times New Roman";
}

.ql-snow .ql-picker.ql-font .ql-picker-label[data-value="sans-serif"]::before, .ql-snow .ql-picker.ql-font .ql-picker-item[data-value="sans-serif"]::before {
  content: "sans-serif" !important;
  font-family: "sans-serif";
}

.ql-font-SimSun {
  font-family: "SimSun",serif;
}

.ql-font-SimHei {
  font-family: "SimHei";
}

.ql-font-Microsoft-YaHei {
  font-family: "Microsoft YaHei";
}

.ql-font-KaiTi {
  font-family: "KaiTi";
}

.ql-font-FangSong {
  font-family: "FangSong";
}

.ql-font-Arial {
  font-family: "Arial";
}

.ql-font-Times-New-Roman {
  font-family: "Times New Roman";
}

.ql-font-sans-serif {
  font-family: "sans-serif";
}

.ql-snow.ql-toolbar .ql-formats .ql-revoke {
  /*background-image: url("../../../../assets/images/icons8-rotate-left-18.png");*/
  width: 20px;
  height: 20px;
  filter: grayscale(100%);
  opacity: 1;
}

.ql-snow.ql-toolbar .ql-formats .ql-revoke:hover {
  /*background-image: url("../../../../assets/images/icons8-rotate-left-18.png");*/
  width: 20px;
  height: 20px;
  filter: none;
  opacity: 0.9;
}



/*加上height和滚动属性就可以，滚动条样式是系统默认样式，可能不同*/
.ql-toolbar.ql-snow .ql-picker.ql-expanded .ql-picker-options {
  border-color: #ccc;
  height: 125px;
  overflow: auto;
}

</style>