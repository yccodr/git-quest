package services

import (
	"fmt"
	"git-quest-be/internal/quests"
)

func GetQuests() []quests.QuestMeta {
	metas := make([]quests.QuestMeta, 0, len(quests.Quests))
	for _, q := range quests.Quests {
		metas = append(metas, q.GetMeta())
	}

	return metas
}

func GetQuest(id int) (quests.Quest, error) {
	q, ok := quests.Quests[id]
	if !ok {
		return nil, fmt.Errorf("quest not found")
	}

	return q, nil
}
