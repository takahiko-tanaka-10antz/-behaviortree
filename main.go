package main

import (
	"fmt"

	. "github.com/takahiko-tanaka-10antz/behaviortree/behaviortree"
	"github.com/takahiko-tanaka-10antz/behaviortree/models/logic"
	"github.com/takahiko-tanaka-10antz/behaviortree/scenario"
)

func main() {
	input := 0.5 // SkillAの発動確率(50%)

	// 各Logicのインスタンスを作成
	user := logic.NewUserLogic(150, input, 1.0)
	enemy := logic.NewEnemyLogic(120)

	err := RootNode(
		NewBranch(
			SequencerNode,

			// 出発する
			NewBranch(
				ActionNode,
				scenario.Start(),
			),

			// 敵に寄る
			NewBranch(
				ConditionNode,
				// 自分のHPが100以上
				user.IsHP100OrMore(),
				user.GetCloser(),
			),

			// 友達Aと友達Bを同時に呼ぶ
			NewBranch(
				ParallelNode,
				user.CallAFriends("A"),
				user.CallAFriends("B"),
			),

			// SKILL A,Bをランダムに2回発動
			RepeaterNode(
				2,
				NewBranch(
					SelectorNode,
					NewBranch(
						ConditionNode,
						user.GetSkills("1"),
						user.ActivateSkills(&enemy.Enemy, "1"),
					),
					NewBranch(
						ConditionNode,
						user.GetSkills("2"),
						user.ActivateSkills(&enemy.Enemy, "2"),
					),
				),
			),

			NewBranch(
				SelectorNode,

				// 敵が死んだら、End1
				NewBranch(
					ConditionNode,
					enemy.Dead(),
					scenario.END1(),
				),

				// 敵が生きたら、End2
				NewBranch(
					ConditionNode,
					enemy.Alive(),
					scenario.END2(),
				),
			),
		),
	)

	// 敵の残りHPを出力
	fmt.Printf("\n敵HP = %d", enemy.Enemy.HP)

	if err != nil {
		panic(err)
	}
}
