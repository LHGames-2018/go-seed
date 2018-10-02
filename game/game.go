////////////////////////////////////////////
//         DO NOT TOUCH THIS FILE         //
////////////////////////////////////////////

package main

type GameInfo struct {
	Player Player
	Map    Map
}

type JSONGameInfo struct {
	Player JSONPlayer `json:"Player"`
	Map    string     `json:"CustomSerializedMap"`
	XMin   int        `json:"xMin"`
	YMin   int        `json:"yMin"`
}

func (json JSONGameInfo) GameInfo() GameInfo {
	var gameinfo GameInfo
	gameinfo.Player = json.Player.Player()
	gameinfo.Map.UnmarshalJSON([]byte(json.Map))
	gameinfo.Map.SetRelativeTo(json.XMin, json.YMin)
	return gameinfo
}
