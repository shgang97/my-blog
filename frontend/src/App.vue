<template>
  <div>
    <button type="button" @click="get">GET</button>
  </div>
</template>

<script>
export default {
  name: 'App',
  methods: {
    async get() {
      const data = await this.$http.get('/user/get')
      console.log(data)
    }
  }

}
</script>
