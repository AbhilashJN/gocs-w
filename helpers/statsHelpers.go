package helpers

import (
	"fmt"
	"sort"

	events "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
)

type DamageMap map[string]map[string]int
type DeathsMap map[string]map[string]int

type damageSummaryItem struct {
	Name  string `json:"name"`
	Given int    `json:"given"`
	Taken int    `json:"taken"`
}

type DamageSummary struct {
	Category string              `json:"category"`
	Items    []damageSummaryItem `json:"items"`
}

type deathsSummaryItem struct {
	Name   string `json:"name"`
	Kills  int    `json:"kills"`
	Deaths int    `json:"deaths"`
}

type DeathsSummary struct {
	Category string              `json:"category"`
	Items    []deathsSummaryItem `json:"items"`
}

type accuracySummaryItem struct {
	Name          string `json:"name"`
	Fired         int    `json:"fired"`
	Hits          int    `json:"hits"`
	HitPercentage string `json:"hitPercentage"`
	Headshots     int    `json:"headshots"`
}

type AccuracySummary struct {
	Category string                `json:"category"`
	Items    []accuracySummaryItem `json:"items"`
}

func updateDamageMap(e events.PlayerHurt, damageMap DamageMap, category string, reqPlayer string) DamageMap {
	var keyGiven string
	var keyTaken string
	switch category {
	case "vs Player":
		keyGiven = e.Player.Name
		keyTaken = e.Attacker.Name

	case "Weapon":
		keyGiven = e.Weapon.String()
		keyTaken = e.Weapon.String()

	case "WeaponClass":
		keyGiven = getEquipmentClassName(e.Weapon.Class())
		keyTaken = getEquipmentClassName(e.Weapon.Class())
	}

	if e.Attacker.Name == reqPlayer {
		_, ok := damageMap[keyGiven]
		if !ok {
			damageMap[keyGiven] = map[string]int{"given": e.HealthDamageTaken, "taken": 0}
		} else {
			damageMap[keyGiven]["given"] += e.HealthDamageTaken
		}
	} else {
		_, ok := damageMap[keyTaken]
		if !ok {
			damageMap[keyTaken] = map[string]int{"given": 0, "taken": e.HealthDamageTaken}
		} else {
			damageMap[keyTaken]["taken"] += e.HealthDamageTaken
		}
	}

	return damageMap
}

func updateDeathsMap(e events.Kill, deathsMap DeathsMap, category string, reqPlayer string) DeathsMap {
	var keyKills string
	var keyDeaths string
	switch category {
	case "vs Player":
		keyKills = e.Victim.Name
		keyDeaths = e.Killer.Name

	case "Weapon":
		keyKills = e.Weapon.String()
		keyDeaths = e.Weapon.String()

	case "WeaponClass":
		keyKills = getEquipmentClassName(e.Weapon.Class())
		keyDeaths = getEquipmentClassName(e.Weapon.Class())
	}

	if e.Killer.Name == reqPlayer {
		_, ok := deathsMap[keyKills]
		if !ok {
			deathsMap[keyKills] = map[string]int{"kills": 1, "deaths": 0}
		} else {
			deathsMap[keyKills]["kills"] += 1
		}
	} else {
		_, ok := deathsMap[keyDeaths]
		if !ok {
			deathsMap[keyDeaths] = map[string]int{"kills": 0, "deaths": 1}
		} else {
			deathsMap[keyDeaths]["deaths"] += 1
		}
	}

	return deathsMap
}

func generateDamageMapSummary(damageMap DamageMap, category string) DamageSummary {
	damageSummaryItems := []damageSummaryItem{}
	for k, v := range damageMap {
		damageSummaryItems = append(damageSummaryItems, damageSummaryItem{Name: k, Given: v["given"], Taken: v["taken"]})
	}

	sort.Slice(damageSummaryItems, func(i, j int) bool {
		sum1 := damageSummaryItems[i].Taken + damageSummaryItems[i].Given
		sum2 := damageSummaryItems[j].Taken + damageSummaryItems[j].Given
		return sum1 > sum2
	})
	damageSummary := DamageSummary{Category: category, Items: damageSummaryItems}
	return damageSummary
}

func GenerateDamageSummary(damageMapByWeaponClass, damageMapByWeapon, damageMapByPlayer map[string]map[string]int) []DamageSummary {
	damageSummaryByPlayer := generateDamageMapSummary(damageMapByPlayer, "vs Player")
	damageSummaryByWeapon := generateDamageMapSummary(damageMapByWeapon, "Weapon")
	damageSummaryByWeaponClass := generateDamageMapSummary(damageMapByWeaponClass, "WeaponClass")

	allDamageSummaries := []DamageSummary{damageSummaryByWeapon, damageSummaryByPlayer, damageSummaryByWeaponClass}
	return allDamageSummaries

}

func generateDeathsSummaryFromMap(deathsMap map[string]map[string]int, category string) DeathsSummary {
	deathsSummaryItems := []deathsSummaryItem{}
	for k, v := range deathsMap {
		deathsSummaryItems = append(deathsSummaryItems, deathsSummaryItem{Name: k, Kills: v["kills"], Deaths: v["deaths"]})
	}

	sort.Slice(deathsSummaryItems, func(i, j int) bool {
		sum1 := deathsSummaryItems[i].Deaths + deathsSummaryItems[i].Kills
		sum2 := deathsSummaryItems[j].Deaths + deathsSummaryItems[j].Kills
		return sum1 > sum2
	})
	deathsSummary := DeathsSummary{Category: category, Items: deathsSummaryItems}
	return deathsSummary
}

func GenerateDeathsSummary(deathsMapByPlayer, deathsMapByWeapon, deathsMapByWeaponClass DeathsMap) []DeathsSummary {

	deathsSummaryByPlayer := generateDeathsSummaryFromMap(deathsMapByPlayer, "vs Player")
	deathsSummaryByWeapon := generateDeathsSummaryFromMap(deathsMapByWeapon, "Weapon")
	deathsSummaryByWeaponClass := generateDeathsSummaryFromMap(deathsMapByWeaponClass, "WeaponClass")
	allSummaries := []DeathsSummary{deathsSummaryByWeapon, deathsSummaryByPlayer, deathsSummaryByWeaponClass}
	return allSummaries
}

func calcHitPercentage(fired, hits int) string {
	firedF := float32(fired)
	hitsF := float32(hits)
	perc := hitsF / firedF
	return fmt.Sprintf("%.2f%%", perc)

}

func GenerateAccuracySummary(weaponFireMap map[string]int, weaponShotsHitMap map[string]map[string]int) []AccuracySummary {
	accuracySummaryItems := []accuracySummaryItem{}
	for k, v := range weaponFireMap {
		item := accuracySummaryItem{Name: k, Fired: v}
		_, ok := weaponShotsHitMap[k]
		if ok {
			item.Hits = weaponShotsHitMap[k]["total"]
			item.Headshots = weaponShotsHitMap[k]["hs"]
		}
		item.HitPercentage = calcHitPercentage(item.Fired, item.Hits)
		accuracySummaryItems = append(accuracySummaryItems, item)
	}
	sort.Slice(accuracySummaryItems, func(i, j int) bool { return accuracySummaryItems[i].Fired > accuracySummaryItems[j].Fired })
	accuracySummary := AccuracySummary{Category: "Weapon", Items: accuracySummaryItems}
	allSummaries := []AccuracySummary{accuracySummary}
	return allSummaries
}
