# behaviortree
アルゴリズム研修課題

## 概要
### 課題仕様
[新卒研修資料_BT_2020_06](https://docs.google.com/presentation/d/1_O9o0hmT-GXCBYGkrTRHGskJfimxt3B43eEkMmYR-U4/edit#slide=id.g88b206bce4_1_60)

### 使用した物
* go（1.14）

## 実行方法
コードの取得

`go get -u github.com/takahiko-tanaka-10antz/behaviortree`

実行ファイルは main.go になります。

 `$go run main.go`

main.goの12行目でSkill1の発動確率を入力しています。（初期値50%）

https://github.com/takahiko-tanaka-10antz/behaviortree/blob/0047b9b6ad5c92eaaa9a86a6a7236fe6c820f62e/main.go#L12

## 出力例
END1（敵が死ぬ）
```
$go run main.go
出発
敵に寄る
フレンドAを呼ぶ
フレンドBを呼ぶ
スキル2発動
スキル2発動
End1

敵HP = 0
```

END2（敵が生きる）
```
$go run main.go
出発
敵に寄る
フレンドAを呼ぶ
フレンドBを呼ぶ
スキル2発動
スキル1発動
End2

敵HP = 10
```

```
$go run main.go
出発
敵に寄る
フレンドAを呼ぶ
フレンドBを呼ぶ
スキル1発動
スキル2発動
End2

敵HP = 10
```

```
$go run main.go
出発
敵に寄る
フレンドAを呼ぶ
フレンドBを呼ぶ
スキル1発動
スキル1発動
End2

敵HP = 20
```

## 参考
[go-behaviortree](https://github.com/joeycumines/go-behaviortree)

[Behavior tree (artificial intelligence, robotics and control)](https://en.wikipedia.org/wiki/Behavior_tree_(artificial_intelligence,_robotics_and_control))

#10antz
