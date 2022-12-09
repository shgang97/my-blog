<template>
  <div class="article-container">
    <div class="meta-container">
      <div class="date vertical-line">1个月前</div>
      <div class="tag-list ">
        <a href="/" class="tag vertical-line">后端</a>
        <a href="/" class="tag vertical-line">网络协议</a>
      </div>
      <div class="category vertical-line">
        <a href="/">Java</a>
      </div>
    </div>
    <div class="content-wrapper">
      <div class="title">
        <router-link :to="'/article/' + id" >{{ title }}</router-link>
      </div>
      <div class="abstract">{{ description }}</div>
      <ul class="action-list">
        <el-space size="14" spacer="|">
          <li class="item view"><i class="iconfont icon-yuedu"></i><span>{{ viewCount }}</span></li>
        </el-space>
        <el-space size="14" spacer="|">
          <li class="item comment"><i class="iconfont icon-pinglun"></i><span>{{ commentCount }}</span></li>
        </el-space>
        <el-space size="14" spacer="|">
          <li class="item like"><i class="iconfont icon-dianzan"></i><span>{{ likeCount }}</span></li>
        </el-space>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  name: "ArticleCard",
  props: {
    id: {
      type: String,
      required: true
    },
    title: {
      type: String,
      required: true
    },
    description: {
      type: String,
      // required: true,
      default: "description"
    },
    viewCount: {
      type: Number,
      default: 0
    },
    commentCount: {
      type: Number,
      default: 0
    },
    likeCount: {
      type: Number,
      default: 0
    },
  }
}
</script>

<style lang="less" scoped>
.meta-container {
  display: flex;
  color: #86909c;

}

a {
  color: inherit;
  text-decoration: none;
}

.vertical-line {
  border-left: 1px #86909c;
  padding: 0 5px;
}

.title {
  font-weight: 700;
  font-size: 16px;
  line-height: 24px;
  color: #1d2129;
  width: 100%;
  display: -webkit-box;
  overflow: hidden;
  text-overflow: ellipsis;
  -webkit-box-orient: vertical;
  -webkit-line-clamp: 1;

}

ul {
  display: flex;
  align-items: center;
}

li {
  list-style: none;
}

.article-container {
  padding: 5px;
  width: 100%;
}

.action-list {
  padding: 0;
}
</style>