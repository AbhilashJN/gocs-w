package parser

import (
	"math"
	"os"

	"gocs-w/helpers"

	"github.com/golang/geo/r2"
	dem "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs"
	events "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/events"
	metadata "github.com/markus-wa/demoinfocs-golang/v2/pkg/demoinfocs/metadata"
)

func GetMapName(demoPath string) string {
	f, err := os.Open(demoPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//init parser
	p := dem.NewParser(f)
	defer p.Close()

	h, herr := p.ParseHeader()
	if herr != nil {
		panic(herr)
	}
	return h.MapName
}

func ListPlayers(demoPath string) []string {
	f, err := os.Open(demoPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//init parser
	p := dem.NewParser(f)
	defer p.Close()

	// // skip to the end of the knife round
	// knifeRoundEnded := false
	// handlerId := p.RegisterEventHandler(func(e events.RoundEnd) {
	// 	if !knifeRoundEnded {
	// 		knifeRoundEnded = true
	// 	}
	// })
	// for !knifeRoundEnded {
	// 	p.ParseNextFrame()
	// }
	// p.UnregisterEventHandler(handlerId)

	//skip to warmup
	isWarmup := false
	for !isWarmup {
		p.ParseNextFrame()
		isWarmup = p.GameState().IsWarmupPeriod()
	}

	//skip to match
	matchStarted := false
	for !matchStarted {
		p.ParseNextFrame()
		matchStarted = !(p.GameState().IsWarmupPeriod())
	}

	allPlayers := p.GameState().Participants().Playing()
	pnames := []string{}

	for _, p := range allPlayers {
		pnames = append(pnames, p.Name)
	}

	return pnames
}

func GetDamageSummaryForPlayer(demoPath string, player string) []helpers.DamageSummary {
	//open demo file
	f, err := os.Open(demoPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//init parser
	p := dem.NewParser(f)
	defer p.Close()

	reqPlayer := player

	// skip to the end of the knife round
	// knifeRoundEnded := false
	// handlerId := p.RegisterEventHandler(func(e events.RoundEnd) {
	// 	if !knifeRoundEnded {
	// 		knifeRoundEnded = true
	// 	}
	// })
	// for !knifeRoundEnded {
	// 	p.ParseNextFrame()
	// }
	// p.UnregisterEventHandler(handlerId)

	//skip to warmup
	isWarmup := false
	for !isWarmup {
		p.ParseNextFrame()
		isWarmup = p.GameState().IsWarmupPeriod()
	}

	//skip to match
	matchStarted := false
	for !matchStarted {
		p.ParseNextFrame()
		matchStarted = !(p.GameState().IsWarmupPeriod())
	}

	//record damage events
	damageMapByWeaponClass := make(helpers.DamageMap)
	damageMapByWeapon := make(helpers.DamageMap)
	damageMapByPlayer := make(helpers.DamageMap)

	p.RegisterEventHandler(func(e events.PlayerHurt) {
		helpers.DamageEventsByPlayer(e, reqPlayer, damageMapByWeaponClass, damageMapByWeapon, damageMapByPlayer)
	})

	// Parse demo to end
	parseErr := p.ParseToEnd()
	if parseErr != nil {
		panic(parseErr)
	}

	//generate damage summary
	damageSummary := helpers.GenerateDamageSummary(damageMapByWeaponClass, damageMapByWeapon, damageMapByPlayer)
	return damageSummary
}

func GenerateDeathsSummaryForPlayer(demoPath string, player string) []helpers.DeathsSummary {
	//open demo file
	f, err := os.Open(demoPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//init parser
	p := dem.NewParser(f)
	defer p.Close()

	reqPlayer := player

	// skip to the end of the knife round
	// knifeRoundEnded := false
	// handlerId := p.RegisterEventHandler(func(e events.RoundEnd) {
	// 	if !knifeRoundEnded {
	// 		knifeRoundEnded = true
	// 	}
	// })
	// for !knifeRoundEnded {
	// 	p.ParseNextFrame()
	// }
	// p.UnregisterEventHandler(handlerId)

	//skip to warmup
	isWarmup := false
	for !isWarmup {
		p.ParseNextFrame()
		isWarmup = p.GameState().IsWarmupPeriod()
	}

	//skip to match
	matchStarted := false
	for !matchStarted {
		p.ParseNextFrame()
		matchStarted = !(p.GameState().IsWarmupPeriod())
	}

	// record deaths
	deathsMapByPlayer := make(map[string]map[string]int)
	deathsMapByWeapon := make(map[string]map[string]int)
	deathsMapByWeaponClass := make(map[string]map[string]int)
	p.RegisterEventHandler(func(e events.Kill) {
		helpers.DeathEvents(e, reqPlayer, deathsMapByPlayer, deathsMapByWeapon, deathsMapByWeaponClass)
	})

	// Parse demo to end
	parseErr := p.ParseToEnd()
	if parseErr != nil {
		panic(parseErr)
	}

	deathsSummary := helpers.GenerateDeathsSummary(deathsMapByPlayer, deathsMapByWeapon, deathsMapByWeaponClass)
	return deathsSummary
}

func GenerateHeatMapPositions(demoPath string, player string) map[string][]r2.Point {
	//open demo file
	f, err := os.Open(demoPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//init parser
	p := dem.NewParser(f)
	defer p.Close()

	//parse demo headers
	h, herr := p.ParseHeader()
	if herr != nil {
		panic(herr)
	}
	mapName := h.MapName
	mapmeta := metadata.MapNameToMap[mapName]
	reqPlayer := player

	// skip to the end of the knife round
	// knifeRoundEnded := false
	// handlerId := p.RegisterEventHandler(func(e events.RoundEnd) {
	// 	if !knifeRoundEnded {
	// 		knifeRoundEnded = true
	// 	}
	// })
	// for !knifeRoundEnded {
	// 	p.ParseNextFrame()
	// }
	// p.UnregisterEventHandler(handlerId)

	//skip to warmup
	isWarmup := false
	for !isWarmup {
		p.ParseNextFrame()
		isWarmup = p.GameState().IsWarmupPeriod()
	}

	//skip to match
	matchStarted := false
	for !matchStarted {
		p.ParseNextFrame()
		matchStarted = !(p.GameState().IsWarmupPeriod())
	}

	// record deaths
	deathPositions := make(map[string][]r2.Point)
	bombPlantPositions := []r2.Point{}

	deathPositions["kill"] = []r2.Point{}
	deathPositions["death"] = []r2.Point{}
	p.RegisterEventHandler(func(e events.Kill) {
		helpers.DeathTaken(e, reqPlayer, mapmeta, deathPositions)
	})
	p.RegisterEventHandler(func(e events.BombPlanted) {
		bombPlantPositions = helpers.BombPlantPosition(e, reqPlayer, mapmeta, bombPlantPositions)
	})

	// Parse demo to end
	parseErr := p.ParseToEnd()
	if parseErr != nil {
		panic(parseErr)
	}

	normalizedDeathPts := []r2.Point{}
	normalizedKillPts := []r2.Point{}
	normalizedBombPlantPts := []r2.Point{}

	for _, p := range deathPositions["death"] {
		normalizedDeathPts = append(normalizedDeathPts, r2.Point{X: math.Floor(p.X / 2), Y: math.Floor(p.Y / 2)})
	}
	for _, p := range deathPositions["kill"] {
		normalizedKillPts = append(normalizedKillPts, r2.Point{X: math.Floor(p.X / 2), Y: math.Floor(p.Y / 2)})
	}

	for _, p := range bombPlantPositions {
		normalizedBombPlantPts = append(normalizedBombPlantPts, r2.Point{X: math.Floor(p.X / 2), Y: math.Floor(p.Y / 2)})
	}

	finalPositions := make(map[string][]r2.Point)

	finalPositions["kill"] = normalizedKillPts
	finalPositions["death"] = normalizedDeathPts
	finalPositions["bomb_plant"] = normalizedBombPlantPts

	return finalPositions
}

func GenerateAccuracySummaryForPlayer(demoPath string, player string) []helpers.AccuracySummary {
	//open demo file
	f, err := os.Open(demoPath)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	//init parser
	p := dem.NewParser(f)
	defer p.Close()

	reqPlayer := player

	// skip to the end of the knife round
	// knifeRoundEnded := false
	// handlerId := p.RegisterEventHandler(func(e events.RoundEnd) {
	// 	if !knifeRoundEnded {
	// 		knifeRoundEnded = true
	// 	}
	// })
	// for !knifeRoundEnded {
	// 	p.ParseNextFrame()
	// }
	// p.UnregisterEventHandler(handlerId)

	//skip to warmup
	isWarmup := false
	for !isWarmup {
		p.ParseNextFrame()
		isWarmup = p.GameState().IsWarmupPeriod()
	}

	//skip to match
	matchStarted := false
	for !matchStarted {
		p.ParseNextFrame()
		matchStarted = !(p.GameState().IsWarmupPeriod())
	}

	// record deaths
	weaponFireMap := make(map[string]int)
	weaponShotsHitMap := make(map[string]map[string]int)

	p.RegisterEventHandler(func(e events.WeaponFire) {
		helpers.WeaponFiredByPlayer(e, reqPlayer, weaponFireMap)
	})
	p.RegisterEventHandler(func(e events.PlayerHurt) {
		helpers.WeaponShotsHitByPlayer(e, reqPlayer, weaponShotsHitMap)
	})

	// Parse demo to end
	parseErr := p.ParseToEnd()
	if parseErr != nil {
		panic(parseErr)
	}

	accuracySummary := helpers.GenerateAccuracySummary(weaponFireMap, weaponShotsHitMap)
	return accuracySummary
}
