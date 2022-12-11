<template>
  <div id="app">
    <h2>新規会員登録</h2>
    <div class="mx-auto text-center">
      <form action="register.php" method="post">
      <div class="form-group">
        <p><label><input type="text" name="id" class="form-control" placeholder="ID" required autofocus></label></p>
      </div>
      <div class="form-group">
        <p><label><input type="password" name="password" class="form-control" placeholder="パスワード" required></label></p>
      </div>
      <div class="form-group">
        <p><button class="btn btn-primary" type="submit">登録</button></p>
      </div>
    </form>
    <p><a href="login.php">ログイン</a></p>
    <p> {{ message }}</p>
  </div>
  </div>
</template>

<script></script>

<style></style>