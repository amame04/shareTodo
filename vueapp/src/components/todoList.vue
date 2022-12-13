<template>
  <div id="todoForm" class="mx-auto py-3">

    <form class="row mx-2" action="request.php?request=new" method="post">
      <!-- CSRFトークン -->
      <input type="hidden" name="token" value="<?php echo $_SESSION['token']; ?>">
      <div class="col">
        <input class="form-control mx-2" type="text" name="todo" placeholder="やること" autofocus required>
      </div>
      <div class="col-sm-2">
        <flat-pickr v-model="date" :config="config" placeholder="いつまで" name="date" class="form-control mx-2"></flat-pickr>
      </div>
      <div class="col-sm-2">
        <div class="form-control">
          <!-- ユーザー名一覧をプルダウンチェックリストで表示-->
          <select multiple="multiple" id="select" class="mx-1" name="shareUser[]" placeholder="共有">
            <!-- ユーザー一覧を取得 -->
            <option v-for="item in userList" v-bind:value="item.id">{{ item.id }}</option>
          </select>
        </div>
      </div>

      <button class="btn btn-primary col-auto mx-3" type="submit">追加</button>
    </form>

    <div class="mx-4 text-center">
      <!-- 該当するTODOを表示 -->
      <!--
      $deleteBtn = '<button class="btn btn-primary col-auto mx-2 deleteBtn" formaction="request.php?request=delete" type="submit">削除</button>';

      $doneBtn = '';
      if ($value['done'] != true) {
        $doneBtn = '<button class="btn btn-primary col-auto mx-2" formaction="request.php?request=done" type="submit">完了</button>';
      }

        $shareUsers .= '<p class="card-text text-success d-inline mx-2"><i class="bi bi-check-square"></i>';
      } else {
        $shareUsers .= '<p class="card-text text-danger d-inline mx-2"><i class="bi bi-square"></i>';
      }
      -->

      <div class="card mx-auto my-2" v-for="item in todoList">
        <div class="card-body">
          <h5 class="card-title">{{ item.todoContent }}</h5>
          <p class="card-text text-muted">{{ item.todoDate }}</p>
          <p class="card-text text-danger d-inline mx-2">
            <i class="bi bi-square"></i>
            $shareUsers
          </p>
          <p class="card-text text-success d-inline mx-2">
            <i class="bi bi-check-square"></i>
            $shareUsers
          </p>
          <form method="post">
            <input type="hidden" name="token" value="$csrfToken">
            <input type="hidden" name="todoid" v-bind:value="item.todoId">
            <button class="btn btn-primary col-auto mx-2" formaction="" type="submit" v-if="!item.deleteFlag">完了</button>

            <button class="btn btn-primary col-auto mx-2 deleteBtn" formaction="request.php?request=delete" type="submit" v-if="!item.deleteFlag">削除</button>
          </form>
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
      date: null,
      config: {
        locale: Japanese,
        dateFormat: 'Y-m-d H:i',
        enableTime: true
      },
      userList: '',
      user: null,
      todoList: null
    }
  },
  components: {
    flatPickr
  },
  created () {
    this.axios.get('http://localhost:8888/', {
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

    this.axios.get('http://localhost:8888/userList', {
      })
      .then(response => {
        this.userList = response.data.userList
      })
      .catch(err => {
        console.error(err)
      })
  },
}

$(function () {
  $('#select').multipleSelect()
})
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
</style>
