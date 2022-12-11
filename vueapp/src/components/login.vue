<template>
  <div id="app">
    <h2>ログイン</h2>
    <div class="mx-auto text-center">
      <form action="login.php" method="post">
        <div class="form-group">
          <p><label><input type="text" name="id" class="form-control" placeholder="ID" required autofocus></label></p>
        </div>
        <div class="form-group">
          <p><label><input type="password" name="password" class="form-control" placeholder="パスワード" required></label></p>
        </div>
        <div class="form-group">
          <p><button class="btn btn-primary" type="submit">ログイン</button></p>
          <p><a href="register.php">新規会員登録</a></p>
        </div>
      </form>
      <p>{{ message }}</p>
    </div>
  </div>
</template>

<script></script>

<style></style>