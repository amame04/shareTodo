<?php
require_once "../app/LoginClass.php";

function login() {
  session_start();
  $id = $_POST['id'];
  $password = $_POST['password'];

  $login = new Login();

  // Login
  if ($login->login($id, $password)) {
    //DBのユーザー情報をセッションに保存
    $_SESSION['id'] = $login->getUser()->getid();
    header('Location: /');
    exit();
  } else {
    return 'IDもしくはパスワードが間違っています。';
  }
}

$msg = '';
if (isset($_POST['id'])) {
  $msg = login();
}
?>

<?php include "../template/header.php"; ?>

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
    <p><?php echo $msg; ?></p>
  </div>


<?php include "../template/footer.php"; ?>