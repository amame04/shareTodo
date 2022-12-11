<?php
require_once "../app/DbClass.php";

function newtodo($userId, $todo, $date, $shareUserArray) {
  $db = new Db();
  $dbh = $db->getDbh();
  if (empty($date)) $date = null;

  try {
    $dbh->beginTransaction();
    $sql = "INSERT INTO todo(todo, date, user) VALUES(:todo, :date, :userId);";
    $stmt = $dbh->prepare($sql);
    $stmt->bindValue(':todo', $todo);
    $stmt->bindValue(':date', $date);
    $stmt->bindValue(':userId', $userId);
    $stmt->execute();

    $todoid = $dbh->lastInsertId();

    $shareUserArray[] = $userId;

    foreach ($shareUserArray as $shareUser) {
      if (empty($shareUser)) break;
      $sql = "INSERT INTO share(id, user) VALUES(:id, :user);";
      $stmt = $dbh->prepare($sql);
      $stmt->bindValue(':id', $todoid);
      $stmt->bindValue(':user', $shareUser);
      $stmt->execute();
    }

    $dbh->commit();
  } catch (Exception $e) {
    $dbh->rollBack();
    return $e;
  }

  return true;
}

function deletetodo($user, $todoid) {
  $db = new Db();
  $dbh = $db->getDbh();

  try {
    $sql = "UPDATE todo SET deleteFlag = true WHERE id = :id AND user = :user";
    $stmt = $dbh->prepare($sql);
    $stmt->bindValue(':id', $todoid);
    $stmt->bindValue(':user', $user);
    $stmt->execute();
  } catch (Exception $e) {
    return $e;
  }

  return true;
}

function donetodo($user, $todoid) {
  $db = new Db();
  $dbh = $db->getDbh();

  try {
    $sql = "UPDATE share SET done = true WHERE id = :id AND user = :user";
    $stmt = $dbh->prepare($sql);
    $stmt->bindValue(':id', $todoid);
    $stmt->bindValue(':user', $user);
    $stmt->execute();
  } catch (Exception $e) {
    return $e;
  }

  return true;
}

session_start();

// CSRF 対策のトークンチェック
if ($_POST['token'] === $_SESSION['token']) {
  if ($_GET['request'] === 'new') {
    newtodo($_SESSION['id'], $_POST['todo'], $_POST['date'], $_POST['shareUser']);
  }
  if ($_GET['request'] === 'delete') {
    deletetodo($_SESSION['id'], $_POST['todoid']);
  }
  if ($_GET['request'] === 'done') {
    donetodo($_SESSION['id'], $_POST['todoid']);
  }
}

header('Location: /');
exit();
?>