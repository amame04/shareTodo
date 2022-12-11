<!DOCTYPE html>
<html lang="ja">
  <head>
    <meta charset="utf-8">

    <title>Share ToDo</title>

    <link href="https://cdn.jsdelivr.net/npm/bootstrap@5.0.2/dist/css/bootstrap.min.css" rel="stylesheet" integrity="sha384-EVSTQN3/azprG1Anm3QDgpJLIm9Nao0Yz1ztcQTwFspd3yD65VohhpuuCOmLASjC" crossorigin="anonymous">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/flatpickr/dist/flatpickr.min.css">
    <link rel="stylesheet" href="https://cdn.jsdelivr.net/npm/bootstrap-icons@1.8.0/font/bootstrap-icons.css">
    <link rel="stylesheet" href="https://unpkg.com/multiple-select@1.5.2/dist/multiple-select.min.css">
    <link rel="stylesheet" href="css/default.css">

    <script src="https://cdn.jsdelivr.net/npm/flatpickr"></script>
    <script src="https://cdn.jsdelivr.net/npm/flatpickr/dist/l10n/ja.js"></script>
    <script src="https://code.jquery.com/jquery-3.6.0.min.js" integrity="sha256-/xUj+3OJU5yExlq6GSYGSHk7tPXikynS7ogEvDej/m4=" crossorigin="anonymous"></script>
  </head>

  <header>
    <div class="text-center">
      <h1>Share ToDo</h1>
    </div>
    <?php 
    if(isset($_SESSION['id'])) {
      $username = htmlspecialchars($_SESSION['id'], \ENT_QUOTES, 'UTF-8');
      $title = "Share ToDo";
      $helloMsg = "<div id=\"helloMsg\" class=\"\"><p class=\"my-0\">こんにちは $username さん！</p><p class=\"py-0 my-0\"><a href=\"logout.php\">ログアウト</a></p></div>";
      echo $helloMsg;
    }
    ?>
  </header>
  <body>