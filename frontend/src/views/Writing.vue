<template>
  <div class="edit-container">
    <div class="header">
      <input placeholder="输入文章标题..." spellcheck="false" maxlength="80" class="title-input title-input">
    </div>
    <div class="v-md-editor-container">
      <v-md-editor v-model="text" height="400px"></v-md-editor>
    </div>
  </div>

</template>

<script>
export default {
  name: 'Writing',
  data() {
    return {
      text: ''
    };
  }
};
</script>

<style lang="less" scoped>

</style>