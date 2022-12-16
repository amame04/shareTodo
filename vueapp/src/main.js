// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import App from './App'
import router from './router'
import './assets/css/default.css'
import axios from 'axios'
import VueAxios from 'vue-axios'

Vue.use(VueAxios, axios)

Vue.config.productionTip = false

/* eslint-disable no-new */
new Vue({
  el: '#app',
  router,
  components: { App },
  template: '<App/>',
  data () {
    return {
      ApiServer: ''
    }
  },
  methods: {
    getCSRFToken: function () {
      this.axios.get(this.$root.ApiServer + '/user', {
        withCredentials: true
      })
        .then(response => {
          this.token = response.data.token
          return this.token
        })
        .catch(err => {
          console.error(err)
        })
    }
  }
})
