var CONTENT_KEY = "CACHE_CONTENT"; // 编辑器内容缓存key
var TITLE_KEY = "CACHE_TITLE"; // 标题缓存key
var AUTO_SAVE_TIME = 5000; // 自动保存时间
var cos = null;
var MdEditor = null;
var headInput = null;
var ArticleItem = {};

function setAjaxToken(xhr) {
  xhr.setRequestHeader("Authorization", localStorage.getItem("AUTH_TOKEN"));
}
function initEditor() {
  // 取默认标题
  headInput.val(ArticleItem.Title);
  // 初始化编辑器
  MdEditor = editormd("editormd", {
    width: "99.5%",
    height: window.innerHeight - 78,
    syncScrolling: "single",
    editorTheme: "default",
    path: CNDURL + "/lib/",
    placeholder: "",
    appendMarkdown: ArticleItem.Markdown,
    codeFold: true,
    saveHTMLToTextarea: true,
    // tocm: true,
    imageUpload: true,
    taskList: true,
    // emoji: true,
    imageFormats: ["jpg", "jpeg", "gif", "png", "bmp", "webp"],
    // imageUploadURL: "/api/v1/uploadfile",
    imageUploadCalback: function (files, cb) {
      uploadImage(files[0], cb);
    },
  });
}
function uploadImage(file, cb) {
  const config = {
    useCdnDomain: true,
    //自行去改七牛云的空间区域的配置
    region: qiniu.region.z0
  };
  const putExtra = {
  };
  // 异步获取临时密钥
  $.ajax({
    url: "/api/qiniu/token",
    type: "GET",
    contentType: "application/json",
    success: function (res) {
      if (res.code !== 200) return alert(res.error);
      const token = res.data;
      const path = "blog/upload/" + Date.now() + "_";
      const observable = qiniu.upload(file, path + file.name, token, putExtra, config)
      const observer = {
        next(res){
          // ...
        },
        error(err){
          // ...
        },
        complete(res){
          console.log(res)
          cb("http://rk3udllz3.hd-bkt.clouddn.com/" + res.key)
        }
      }
      const subscription = observable.subscribe(observer) // 上传开始

    },
    beforeSend: setAjaxToken,
  });

}


function getArticleItem(id) {
  $.ajax({
    url: "/api/post/" + id,
    type: "GET",
    contentType: "application/json",
    success: function (res) {
      if (res.code != 200) {
        initEditor();
        return alert(res.error);
      }
      ArticleItem = res.data.Article || {};
      initActive();
      initEditor();
    },
    beforeSend: setAjaxToken,
  });
}
function initActive() {
  $(".category li[value=" + ArticleItem.CategoryId + "]")
    .addClass("active")
    .siblings()
    .removeClass("active");
  $(".type-box li[value=" + ArticleItem.Type + "]")
    .addClass("active")
    .siblings()
    .removeClass("active");
  $(".slug-input").val(ArticleItem.slug);
}
function initCache() {
  headInput = $(".header-input");
  var query = new URLSearchParams(location.search);
  var _id = query.get("id");
  if (_id) return getArticleItem(_id);
  // 取本地缓存
  ArticleItem.Title = window.localStorage.getItem(TITLE_KEY);
  ArticleItem.Markdown = window.localStorage.getItem(CONTENT_KEY) || "";
  // initEditor
  initEditor();
  // 自动保存
  setInterval(() => saveHandler, AUTO_SAVE_TIME);
}

function saveHandler() {
  window.localStorage.setItem(TITLE_KEY, headInput.val());
  window.localStorage.setItem(CONTENT_KEY, MdEditor.getMarkdown());
}
function clearHandler() {
  window.localStorage.removeItem(TITLE_KEY);
  window.localStorage.removeItem(CONTENT_KEY);
}

// 发布
function publishHandler() {
  if (!ArticleItem.CategoryId) return $(".publish-tip").text("请选择分类");
  ArticleItem.Slug = $(".slug-input").val();
  if (ArticleItem.Type == 1 && !ArticleItem.Slug)
    return $(".publish-tip").text("请输入自定义链接");
  ArticleItem.Title = headInput.val();
  if (!ArticleItem.Title) return $(".publish-tip").text("请输入标题");
  ArticleItem.Markdown = MdEditor.getMarkdown();
  if (!ArticleItem.Markdown) return $(".publish-tip").text("请输入正文");
  ArticleItem.Content = MdEditor.getPreviewedHTML();

  $.ajax({
    url: "/api/post",
    type: ArticleItem.Pid ? "PUT" : "POST",
    contentType: "application/json",
    data: JSON.stringify(ArticleItem),
    success: function (res) {
      if (res.code !== 200) {
        return alert(res.error);
      }
      if (ArticleItem.Pid) {
        $(".publish-tip").text("更新成功");
      } else {
        $(".publish-tip").text("发布成功");
      }
      ArticleItem = res.data || {};
      if (!ArticleItem.Pid) {
        clearHandler();
      }
      location.href = "/p/" + ArticleItem.Pid;
    },
    beforeSend: setAjaxToken,
  });
}

$(function () {
  // 初始化缓存
  initCache();
  // 返回首页
  var back = $(".home-btn");
  back.click(function () {
    saveHandler();
    location.href = ArticleItem.pid ? "/p/" + ArticleItem.pid + ".html" : "/";
  });
  if (location.search) back.text("查看");
  // 保存
  $(".save-btn").click(saveHandler);
  var drop = $(".publish-drop");
  // 显示
  $(".publish-show").click(function () {
    drop.show();
  });
  // 隐藏
  $(".publish-close").click(function () {
    drop.hide();
  });
  $(".cancel-btn").click(function () {
    drop.hide();
  });
  // 发布逻辑
  $(".publish-btn").click(publishHandler);
  // 选择分类
  $(".category").on("click", "li", function (event) {
    var target = $(event.target);
    target.addClass("active").siblings().removeClass("active");
    ArticleItem.CategoryId = target.attr("value");
    $(".publish-tip").text("");
  });
  // 选择类型
  $(".type-box").on("click", "li", function (event) {
    var target = $(event.target);
    target.addClass("active").siblings().removeClass("active");
    ArticleItem.Type = Number(target.attr("value") || 0);
  });
});
