package cmd

// GetZenMap returns half character to full map
func GetZenMap() map[string]string {
	zenMap := make(map[string]string)

	zenMap["1"] = "１"
	zenMap["2"] = "２"
	zenMap["3"] = "３"
	zenMap["4"] = "４"
	zenMap["5"] = "５"
	zenMap["6"] = "６"
	zenMap["7"] = "７"
	zenMap["8"] = "８"
	zenMap["9"] = "９"
	zenMap["0"] = "０"
	zenMap[" "] = "　"

	return zenMap
}
