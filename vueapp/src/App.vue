<template>
  <div id="app">
    <header>
      <div class="text-center">
        <h1>Share ToDo</h1>
      </div>
      <div id="helloMsg" v-if="username != null">
        <p class="my-0">こんにちは {{ username }} さん！</p>
        <p class="py-0 my-0"><router-link to="/logout">ログアウト</router-link></p>
      </div>
    </header>
    <router-view/>
  </div>
</template>

<script>
export default {
  name: 'App',
  data () {
    return {
      username: ''
    }
  },
  watch: {
    $route () {
      this.reload()
    }
  },
  created () {
    this.reload()
    this.$root.reload = this.reload
  },
  methods: {
    reload: async function () {
      await this.axios.get(this.$root.ApiServer + '/user', {
        withCredentials: true
      })
        .then(response => {
          this.username = response.data.user
          this.$root.token = response.data.token
        })
        .catch(err => {
          console.error(err)
        })
    }
  }
}
</script>

<style>
#app {
  font-family: 'Avenir', Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
}
</style>
