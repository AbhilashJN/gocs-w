package helpers

import (
	"fmt"

	"github.com/golang/geo/r2"
	"github.com/golang/geo/r3"
	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
	metadata "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/metadata"
)

func getEquipmentClassName(ec common.EquipmentClass) string {
	switch ec {
	case common.EqClassUnknown:
		return "Unknown"
	case common.EqClassPistols:
		return "Pistols"
	case common.EqClassSMG:
		return "SMG"
	case common.EqClassHeavy:
		return "Heavy"
	case common.EqClassRifle:
		return "Rifle"
	case common.EqClassEquipment:
		return "Equipment"
	case common.EqClassGrenade:
		return "Grenade"
	}
	return "Unknown"
}

func getPlayerPositions(p dem.Parser) map[string]r3.Vector {
	allPlayers := p.GameState().Participants().Playing()
	playerpos := make(map[string]r3.Vector)
	for _, player := range allPlayers {
		if player.IsAlive() {
			playerpos[player.Name] = player.Position()
			fmt.Println(player.Name)
		}
	}
	return playerpos
}

func BombPlantPosition(e events.BombPlanted, reqPlayer string, mapmeta metadata.Map, positions []r2.Point) []r2.Point {
	if e.BombEvent.Player.Name != reqPlayer {
		return positions
	}
	pos := e.BombEvent.Player.LastAlivePosition
	x, y := mapmeta.TranslateScale(pos.X, pos.Y)
	positions = append(positions, r2.Point{X: x, Y: y})
	return positions
}

func DeathTaken(e events.Kill, reqPlayer string, mapmeta metadata.Map, deathPositions map[string][]r2.Point) {
	if e.Killer == nil {
		return
	}

	if e.Killer.Name != reqPlayer && e.Victim.Name != reqPlayer {
		return
	}
	var eventType string
	if e.Killer.Name == reqPlayer {
		eventType = "kill"
	} else {
		eventType = "death"
	}
	deathPos := e.Victim.LastAlivePosition
	x, y := mapmeta.TranslateScale(deathPos.X, deathPos.Y)
	deathPositions[eventType] = append(deathPositions[eventType], r2.Point{X: x, Y: y})
}

func DeathEvents(e events.Kill, reqPlayer string, deathsMapByPlayer, deathsMapByWeapon, deathsMapByWeaponClass map[string]map[string]int) {
	if e.Killer == nil {
		return
	}

	if e.Killer.Name != reqPlayer && e.Victim.Name != reqPlayer {
		return
	}

	updateDeathsMap(e, deathsMapByPlayer, "vs Player", reqPlayer)
	updateDeathsMap(e, deathsMapByWeapon, "Weapon", reqPlayer)
	updateDeathsMap(e, deathsMapByWeaponClass, "WeaponClass", reqPlayer)
}

func DamageEventsByPlayer(e events.PlayerHurt, reqPlayer string, damageMapByWeaponClass, damageMapByWeapon, damageMapByPlayer DamageMap) {
	if e.Attacker == nil {
		return
	}

	if e.Attacker.Name != reqPlayer && e.Player.Name != reqPlayer {
		return
	}

	updateDamageMap(e, damageMapByPlayer, "vs Player", reqPlayer)
	updateDamageMap(e, damageMapByWeapon, "Weapon", reqPlayer)
	updateDamageMap(e, damageMapByWeaponClass, "WeaponClass", reqPlayer)
}

func WeaponFiredByPlayer(e events.WeaponFire, reqPlayer string, weaponFireMap map[string]int) {
	if e.Shooter != nil && e.Shooter.Name == reqPlayer {
		weaponClass := e.Weapon.Class()
		if weaponClass != common.EqClassUnknown && weaponClass != common.EqClassEquipment && weaponClass != common.EqClassGrenade {
			weaponName := e.Weapon.String()
			_, ok := weaponFireMap[weaponName]
			if !ok {
				weaponFireMap[weaponName] = 1
			} else {
				weaponFireMap[weaponName] += 1
			}
		}
	}
}

func WeaponShotsHitByPlayer(e events.PlayerHurt, reqPlayer string, weaponShotsHitMap map[string]map[string]int) {
	if e.Attacker != nil && e.Attacker.Name == reqPlayer {
		weaponClass := e.Weapon.Class()
		if weaponClass != common.EqClassUnknown && weaponClass != common.EqClassEquipment && weaponClass != common.EqClassGrenade {
			weaponName := e.Weapon.String()
			isHeadshot := e.HitGroup == events.HitGroupHead
			_, ok := weaponShotsHitMap[weaponName]
			if !ok {
				weaponShotsHitMap[weaponName] = map[string]int{"total": 1, "hs": 0}
			} else {
				weaponShotsHitMap[weaponName]["total"] += 1
			}

			if isHeadshot {
				weaponShotsHitMap[weaponName]["hs"] += 1
			}
		}
	}
}
