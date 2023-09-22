package database

import "fmt"

func GetCacheByUID(uid string) *Cache {
	if profile, ok := profiles[uid]; ok {
		return profile.Cache
	}

	fmt.Println("Profile with UID ", uid, " does not have cache")
	return nil
}

func GetQuestCacheByUID(uid string) *QuestCache {
	if profile, ok := profiles[uid]; ok {
		return profile.Cache.Quests
	}
	return nil
}

func GetTraderCacheByUID(uid string) *TraderCache {
	if profile, ok := profiles[uid]; ok {
		return profile.Cache.Traders
	}
	return nil
}

func (profile *Profile) SetCache() *Cache {
	var cache *Cache
	if profile.Cache == nil {
		cache = &Cache{
			Skills: &SkillsCache{
				Common: make(map[string]int8),
			},
			Hideout: &HideoutCache{
				Areas: make(map[int8]int8),
			},
			Quests: &QuestCache{
				Index: make(map[string]int8),
			},
			Traders: &TraderCache{
				Index:         make(map[string]*AssortIndex),
				Assorts:       make(map[string]*Assort),
				LoyaltyLevels: make(map[string]int8),
			},
		}
	} else {
		cache = profile.Cache
	}

	if profile.Character != nil {
		for index, quest := range profile.Character.Quests {
			cache.Quests.Index[quest.QID] = int8(index)
		}

		for index, commonSkill := range profile.Character.Skills.Common {
			cache.Skills.Common[commonSkill.ID] = int8(index)
		}

		for index, area := range profile.Character.Hideout.Areas {
			cache.Hideout.Areas[int8(area.Type)] = int8(index)
		}

		cache.Inventory = SetInventoryContainer(&profile.Character.Inventory)
	}
	return cache
}

type Cache struct {
	Inventory *InventoryContainer
	Skills    *SkillsCache
	Hideout   *HideoutCache
	Quests    *QuestCache
	Traders   *TraderCache
}

type SkillsCache struct {
	Common map[string]int8
}

type HideoutCache struct {
	Areas map[int8]int8
}

type TraderCache struct {
	Index         map[string]*AssortIndex
	Assorts       map[string]*Assort
	LoyaltyLevels map[string]int8
}

type QuestCache struct {
	Index map[string]int8
}
