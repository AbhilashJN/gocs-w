package main

import (
	api "github.com/AbhilashJN/gocs-core/api"

	"github.com/golang/geo/r2"
	"github.com/wailsapp/wails"
)

type DemoFile struct {
	filepath string
	runtime  *wails.Runtime
}

func NewDemo() *DemoFile {
	d := &DemoFile{}
	return d
}

func (d *DemoFile) WailsInit(runtime *wails.Runtime) error {
	d.runtime = runtime
	d.runtime.Window.Fullscreen()
	return nil
}

func (d *DemoFile) SelFile() string {
	// path := d.runtime.Dialog.SelectFile("Select a demo file", "*.dem")
	// fmt.Println(path, d.filepath)
	// if len(path) > 0 {
	// 	d.filepath = path
	// } else {
	// 	fmt.Println("Cancelling")
	// 	os.Exit(0)
	// }

	d.filepath = "demo/mydemo.dem"

	return api.GetMapName(d.filepath)
}

func (d *DemoFile) GetPlayersList() []string {
	return api.ListPlayers(d.filepath)
}

func (d *DemoFile) GetDamageSummaryForPlayer(player string) map[string]interface{} {
	result := api.GetDamageSummaryForPlayer(d.filepath, player)
	return map[string]interface{}{
		"item_fields":       []string{"name", "given", "taken"},
		"data":              result,
		"comparison_fields": []string{"given", "taken"},
	}
}

func (d *DemoFile) GetDeathsPositionForPlayer(player string) map[string][]r2.Point {
	return api.GetHeatMapPositions(d.filepath, player)
}

func (d *DemoFile) GetDeathsSummaryForPlayer(player string) map[string]interface{} {
	result := api.GetDeathsSummaryForPlayer(d.filepath, player)
	return map[string]interface{}{
		"item_fields":       []string{"name", "kills", "deaths"},
		"data":              result,
		"comparison_fields": []string{"kills", "deaths"},
	}
}

func (d *DemoFile) GetAccuracySummaryForPlayer(player string) map[string]interface{} {
	result := api.GenerateAccuracySummaryForPlayer(d.filepath, player)
	return map[string]interface{}{
		"item_fields":       []string{"name", "fired", "hits", "hitPercentage", "headshots"},
		"data":              result,
		"comparison_fields": []string{},
	}
}

func (d *DemoFile) GetStatsForPlayerWrapper(player string, statType string) interface{} {
	switch statType {
	case "Damage":
		return d.GetDamageSummaryForPlayer(player)
	case "Deaths/Kills":
		return d.GetDeathsSummaryForPlayer(player)
	case "Accuracy":
		return d.GetAccuracySummaryForPlayer(player)
	default:
		return d.GetDamageSummaryForPlayer(player)
	}
}
