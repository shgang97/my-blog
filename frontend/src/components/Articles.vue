<template>
  <div>
    <div class="article-list" v-for="i in 4">
      <ArticleCard></ArticleCard>
    </div>
    <div>
      <Pagination></Pagination>
    </div>
  </div>
</template>

<script>
import ArticleCard from './ArticleCard.vue'
import Pagination from './Pagination.vue'
export default {
  name: "Articles",
  components: {
    ArticleCard,
    Pagination
  }
}
</script>

<style lang="less" scoped>

</style>