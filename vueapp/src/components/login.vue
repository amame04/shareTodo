<template>
  <div id="app">
    <h2 v-if="path == '/login'">ログイン</h2>
    <h2 v-else-if="path == '/logout'">ログアウト</h2>
    <h2 v-else-if="path == '/register'">新規会員登録</h2>

    <div class="mx-auto text-center">
      <div v-if="path != '/logout' && !result">
        <form @submit.prevent="exec">
          <div class="form-group">
            <p><label><input type="text" name="id" class="form-control" placeholder="ID" v-model="id" required autofocus></label></p>
          </div>
          <div class="form-group">
            <p><label><input type="password" name="password" class="form-control" placeholder="パスワード" v-model="password" required></label></p>
          </div>
          <div class="form-group">
            <p>
              <button class="btn btn-primary" type="submit">{{ submit }}</button>
            </p>
          </div>
        </form>
      </div>
      <p><router-link v-bind:to="link">{{ link_str }}</router-link></p>
      <p class="text-danger">{{ error_message }}</p>
      <p class="text-success">{{ message }}</p>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      path: '',
      submit: '',
      result: false,
      link: '',
      link_str: '',
      message: '',
      error_message: '',
      id: '',
      password: ''
    }
  },
  watch: {
    $route () {
      location.reload()
    }
  },
  mounted() {
    this.path = this.$route.path
    if (this.path == '/login') {
      this.submit = 'ログイン'
      this.link = '/register'
      this.link_str = '新規会員登録'
    } else if (this.path == '/register') {
      this.submit = '登録'
      this.link = '/login'
      this.link_str = 'ログイン'
    } else {
      this.link = '/login'
      this.link_str = 'ログイン'
      this.exec()
    }
  },
  methods: {
    exec: function () {
      if (this.path == '/logout') {
        var func = '/login'
      } else {
        var func = this.path
      }
      this.axios.get('http://localhost:8888'+func, {
        withCredentials: true,
        params: {
          id: this.id,
          password: this.password
        }
      })
      .then(response => {
        console.log(response.data)
        if (response.data.success) {
          switch (this.path) {
            case '/login':
              this.message = '認証に成功しました'
              this.$router.push('/')
              location.reload()
              break
            case '/register':
              this.message = '登録が完了しました'
              this.result = true;
              break
          }
          this.error_message = ''
            
        } else {
          switch (this.path) {
            case '/login':
              this.message = ''
              this.error_message = '認証に失敗しました'
              break
            case '/logout':
              this.message = 'ログアウトしました'
              this.error_message = ''
              break
            case '/register':
              this.message = ''
              this.error_message = '同じユーザー名が存在しています'
              break
          }
        }
      })
      .catch(err => {
        this.message = ''
        this.error_message = '接続に失敗しました'
      })
    }
  }
}
</script>
