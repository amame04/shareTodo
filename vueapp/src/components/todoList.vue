<template>
  <div class="mx-auto py-3">

    <form class="row mx-2" action="request.php?request=new" method="post">
      <!-- CSRFトークン -->
      <input type="hidden" name="token" value="<?php echo $_SESSION['token']; ?>">
      <div class="col">
        <input class="form-control mx-2" type="text" name="todo" placeholder="やること" autofocus required>
      </div>
      <div class="col-sm-2">
        <input id="flatpickr" class="form-control mx-2" type="text" name="date" placeholder="いつまで" required>
      </div>
      <div class="col-sm-2">
        <div class="form-control">
          <!-- ユーザー名一覧をプルダウンチェックリストで表示-->
          <select multiple="multiple" id="select" class="mx-1" name="shareUser[]" placeholder="共有">
            <!-- ユーザー一覧を取得 -->
            <option value="$userId">$userId</option>
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

      <div class="card mx-auto my-2">
        <div class="card-body">
          <h5 class="card-title">$todo</h5>
          <p class="card-text text-muted">$date</p>
          $shareUsers
          <form method="post">
            <input type="hidden" name="token" value="$csrfToken">
            <input type="hidden" name="todoid" value="$todoId">
            $doneBtn
            $deleteBtn
          </form>
        </div>
      </div>
      <h4>TODOを追加しましょう!</h4>
    </div>

  </div>
</template>

<script>
</script>

<style>
</style>