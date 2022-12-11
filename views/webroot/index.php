<?php
require_once "../app/DbClass.php";

session_start();
if (empty($_SESSION['id'])) {
  header('Location: ./login.php');
  exit;
}

$csrfToken = bin2hex(random_bytes(32));
$_SESSION['token'] = $csrfToken;

$db = new Db();
$dbh = $db->getDbh();
?>

<?php include "../template/header.php"; ?>

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
              <?php
                try {
                  $sql = 'SELECT id FROM users';
                  $stmt = $dbh->prepare($sql);
                  $stmt->bindValue(':id', $_SESSION['id']);
                  $stmt->execute();
                  $result = $stmt->fetchAll();
                } catch (Exception $e) {
                  echo $e;
                }

                if ($result !== false && sizeof($result) !== 0) {
                  foreach ($result as $user) {
                    if($user['id'] === $_SESSION['id']) continue;
                    $userId = htmlspecialchars($user['id'], ENT_QUOTES);
                    echo <<< EOM
                      <option value="$userId">
                     $userId 
                      </option>
                    EOM;
                  }
                }
              ?>
          </select>
        </div>
      </div>

      <button class="btn btn-primary col-auto mx-3" type="submit">追加</button>
    </form>

    <div class="mx-4 text-center">
      <?php 
        // 該当するTODOを表示 
        try {
          $sql = 'SELECT todo.id, todo, date, todo.user, done 
                  FROM todo INNER JOIN share
                    ON todo.id = share.id
                  WHERE share.user = :id AND deleteFlag != true 
                  ORDER BY date IS NULL ASC, date;';
          $stmt = $dbh->prepare($sql);
          $stmt->bindValue(':id', $_SESSION['id']);
          $stmt->execute();
          $result = $stmt->fetchAll();
        } catch (Exception $e) {
          echo $e;
        }

        if ($result !== false && sizeof($result) !== 0) {
          foreach($result as $value) {
            $todo = htmlspecialchars($value['todo'], ENT_QUOTES);
            $todoId = htmlspecialchars($value['id'], ENT_QUOTES);
            $csrfToken = $_SESSION['token'];

            if (isset($value['date'])) {
              $date = '期限:'. htmlspecialchars($value['date'], ENT_QUOTES);
            } else {
              $date = '無期限';
            }

            $deleteBtn = '';
            if ($value['user'] === $_SESSION['id']) {
              $deleteBtn = '<button class="btn btn-primary col-auto mx-2 deleteBtn" formaction="request.php?request=delete" type="submit">削除</button>';
            }

            $doneBtn = '';
            if ($value['done'] != true) {
              $doneBtn = '<button class="btn btn-primary col-auto mx-2" formaction="request.php?request=done" type="submit">完了</button>';
            }

            try {
              $sql = 'SELECT user, done 
                      FROM share 
                      WHERE id = :id
                      ORDER BY done';
              $stmt = $dbh->prepare($sql);
              $stmt->bindValue(':id', $todoId);
              $stmt->execute();
              $shareUserList = $stmt->fetchAll();
            } catch (Exception $e) {
              echo $e;
            }

            $shareUsers = '';
            if ($shareUserList !== false) {
              $shareUsers .= '<div class="my-2">';
              foreach ($shareUserList as $shareUser) {
                if ($shareUser['done'] == true) {
                  $shareUsers .= '<p class="card-text text-success d-inline mx-2"><i class="bi bi-check-square"></i>';
                } else {
                  $shareUsers .= '<p class="card-text text-danger d-inline mx-2"><i class="bi bi-square"></i>';
                }
                $shareUserName = htmlspecialchars($shareUser['user'], ENT_QUOTES);
                if ($shareUserName === $_SESSION['id']) $shareUserName = 'あなた';
                $shareUsers .= ' ' . $shareUserName . '</p>';
              }
              $shareUsers .= '</div>';
            }

            echo <<< EOM
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
            EOM;
          }
        } else {
          // todoがない場合
          echo '<h4>TODOを追加しましょう!</h4>';
        }
      ?>
    </div>

  </div>

<?php include "../template/footer.php"; ?>