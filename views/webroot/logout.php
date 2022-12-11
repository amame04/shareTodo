<?php
session_start();
// sessionを破棄
$_SESSION = array();
session_destroy();
?>

<?php include "../template/header.php"; ?>

  <div class="mx-auto text-center">
    <p>ログアウトしました。</p>
    <a href="login.php">ログインへ</a>
  </div>

<?php include "../template/footer.php"; ?>