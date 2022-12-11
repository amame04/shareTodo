<?php
require_once "../app/RegisterClass.php";

function register() {
  $reg = new Register();
  return $reg->register($_POST['id'], $_POST['password']);
}

$msg = false;
if (isset($_POST['id'])) {
  $msg = register();
}

if ($msg === true) {
  $html = <<< EOM
  <p>登録が完了しました。</p>
  <a href="./login.php">ログイン</a>
  EOM;
} else {
  //登録が失敗した時 OR まだ登録ボタンを押していないとき
  $html = <<< EOM
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
  EOM;
}
?>

<?php include "../template/header.php"; ?>

  <h2>新規会員登録</h2>

  <div class="mx-auto text-center">
    <?php echo $html; ?>
    <p>
      <?php 
        // $msg に文字列が入っている場合の時のみ表示する
        if ($msg !== false && $msg !== true) echo $msg; 
      ?>
    </p>
  </div>

<?php include "../template/footer.php"; ?>