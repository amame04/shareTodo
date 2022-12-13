import Vue from 'vue'
import Router from 'vue-router'
import todoList from '@/components/todoList'
import login from '@/components/login'

Vue.use(Router)

export default new Router({
  routes: [
    {
      path: '/',
      component: todoList
    },
    {
      path: '/login',
      component: login
    },
    {
      path: '/register',
      component: login
    },
    {
      path: '/logout',
      component: login
    }
  ]
})
