<template>
  <div id="todoForm" class="mx-auto py-3">

    <form id="newTodo" class="row mx-4" @submit.prevent="registerTodo">
      <!-- CSRFトークン -->
      <input type="hidden" name="token" v-bind:value="token">
      <div class="col">
        <input v-model="todo" class="form-control" type="text" placeholder="やること" autofocus required>
      </div>
      <div class="col-sm-2">
        <flat-pickr id="datetime-flatpickr" v-bind:modelValue="datetime" :config="config" placeholder="いつまで" class="form-control"></flat-pickr>
      </div>
      <div class="col-sm-2">
        <div class="form-control">
          <!-- ユーザー名一覧をプルダウンチェックリストで表示-->
          <div class="multiselect">
            <div class="selectBox mx-1">
              <select disabled>
                <option>{{ selectPlaceholder }}</option>
              </select>
              <div class="overSelect"></div>
            </div>
            <div id="checkboxes" class="mx-2">
              <label class="text-start" v-for="item in userList" v-bind:key="item.id" v-if="item.id != user">
                <input type="checkbox" class="mx-1" v-bind:value="item.id" v-model="checkedUser" v-on:change="updatePlaceholder">{{ item.id }}
              </label>
            </div>
          </div>
        </div>
      </div>
      <button class="btn btn-primary col-auto" type="submit">追加</button>
    </form>

    <div id="todoList" class="mx-4 text-center">

      <div class="card my-2" v-for="item in todoList" v-bind:key="item.todoId">
        <div class="card-body">
          <h5 class="card-title">{{ item.todoContent }}</h5>
          <p class="card-text text-muted">{{ item.todoDate }}</p>
          <p class="card-text d-inline mx-2" v-for="user in item.shareUsers" v-bind:key="user.user">
            <span class="text-success" v-if="user.doneFlag">
              <i class="bi bi-check-square"></i>
              {{ user.user }}
            </span>
            <span class="text-danger" v-else>
              <i class="bi bi-square"></i>
              {{ user.user }}
            </span>
          </p>
          <div>
            <button class="btn btn-primary col-auto mx-2" v-on:click="doneTodo(item.todoId)" v-if="!item.doneFlag">完了</button>
            <button class="btn btn-primary col-auto mx-2 deleteBtn" v-on:click="deleteTodo(item.todoId)" v-if="item.createdUser == user">削除</button>
          </div>
        </div>
      </div>
      <h4 v-if="todoList == null">TODOを追加しましょう!</h4>
    </div>

  </div>
</template>

<script>
import flatPickr from 'vue-flatpickr-component'
import 'flatpickr/dist/flatpickr.css'
import {Japanese} from 'flatpickr/dist/l10n/ja.js'

export default {
  data () {
    return {
      config: {
        locale: Japanese,
        dateFormat: 'Y-m-d H:i',
        enableTime: true
      },
      userList: '',
      user: null,
      todoList: null,
      expand: false,
      selectPlaceholder: '共有',
      todo: '',
      datetime: '',
      checkedUser: [],
      token: ''
    }
  },
  components: {
    flatPickr
  },
  created () {
    this.reloadTodoList()

    this.axios.get('http://localhost:8888/userList', {
    })
      .then(response => {
        this.userList = response.data.userList
      })
      .catch(err => {
        console.error(err)
      })
  },
  methods: {
    reloadTodoList: function (e) {
      this.axios.get('http://localhost:8888/todoList', {
        withCredentials: true
      })
        .then(response => {
          this.user = response.data.user
          this.todoList = response.data.todoList
          if (response.data.user == null) {
            this.$router.push('/login')
            location.reload()
          }
        })
        .catch(err => {
          console.error(err)
        })
    },
    showCheckBoxes: function (e) {
      var checkboxes = document.getElementById('checkboxes')
      if (!this.expanded && e.target.closest('.multiselect')) {
        checkboxes.style.display = 'block'
        this.expanded = true
      } else {
        checkboxes.style.display = 'none'
        this.expanded = false
      }
    },
    updatePlaceholder: function () {
      if (this.checkedUser.length === 0) {
        this.selectPlaceholder = '共有'
      } else {
        this.selectPlaceholder = ''
        for (var user of this.checkedUser) {
          this.selectPlaceholder += user + ', '
        }
      }
    },
    registerTodo: function () {
      this.axios.get('http://localhost:8888/registerTodo', {
        withCredentials: true,
        params: {
          todo: this.todo,
          date: document.getElementById('datetime-flatpickr').value,
          shareUsers: this.checkedUser
        }
      })
        .then(response => {
          if (response.data.success) {
            this.todo = ''
            this.datetime = ''
            document.getElementById('datetime-flatpickr').value = ''
            this.checkedUser = []
            this.reloadTodoList()
            this.updatePlaceholder()
          } else {
            alert('failed')
          }
        })
        .catch(err => {
          console.error(err)
        })
    },
    doneTodo: function (todoId) {
      this.axios.get('http://localhost:8888/doneTodo', {
        withCredentials: true,
        params: {
          todoId: todoId
        }
      })
        .then(response => {
          if (response.data.success) {
            this.reloadTodoList()
          } else {
            alert('failed')
          }
        })
        .catch(err => {
          console.error(err)
        })
    },
    deleteTodo: function (todoId) {
      this.axios.get('http://localhost:8888/deleteTodo', {
        withCredentials: true,
        params: {
          todoId: todoId
        }
      })
        .then(response => {
          if (response.data.success) {
            this.reloadTodoList()
          } else {
            alert('failed')
          }
        })
        .catch(err => {
          console.error(err)
        })
    }
  },
  mounted: function () {
    document.addEventListener('click', (e) => this.showCheckBoxes(e), false)
  },
  destroyed: function () {
    document.removeEventListener('click', (e) => this.showCheckBoxes(e), false)
  }
}
</script>

<style>
@import 'flatpickr/dist/flatpickr.min.css';

/* 日曜日：赤 */
.flatpickr-weekdays .flatpickr-weekday:nth-child(1),
.flatpickr-days .flatpickr-day:not(.flatpickr-disabled):not(.prevMonthDay):not(.nextMonthDay):nth-child(7n+1) {
    color: #ff0000;
}

/* 土曜日：青 */
.flatpickr-weekdays .flatpickr-weekday:nth-child(7),
.flatpickr-days .flatpickr-day:not(.flatpickr-disabled):not(.prevMonthDay):not(.nextMonthDay):nth-child(7n) {
    color: #172dee;
}

#todoForm {
  width: 100%;
}

.row>.col,
.row>.col-sm-2 {
  padding-left: 0;
}

#newTodo {
  position: absolute;
  z-index: 1;
  left: 0;
  right: 0;
}

#newTodo > button {
  height: 100%;
}

#todoList {
  margin-top: 3em;
}

.multiselect {
  /*width: 200px;*/
  position: relative;
}

.selectBox {
  position: relative;
}

.selectBox select {
  width: 100%;
}

.overSelect {
  position: absolute;
  left: 0;
  right: 0;
  top: 0;
  bottom: 0;
}

.selectBox > select{
  border: none;
}

#checkboxes {
  display: none;
}

#checkboxes label {
  display: block;
}

#checkboxes label:hover {
  background-color: #1e90ff;
}
</style>
