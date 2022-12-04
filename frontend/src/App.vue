<template>
  <div>
    <router-view></router-view>
  </div>
</template>

<script>
export default {
  name: 'App',
  components: {
  },
  methods: {
    async get() {
      const data = await this.$http.get('/user/get')
      console.log(data)
    }
  }

}
</script>
