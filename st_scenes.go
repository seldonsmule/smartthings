package smartthings

type StScenes struct {
	Links struct {
		Next     interface{} `json:"next"`
		Previous interface{} `json:"previous"`
	} `json:"_links"`
	Items []struct {
		APIVersion       string      `json:"apiVersion"`
		CreatedBy        string      `json:"createdBy"`
		CreatedDate      float64     `json:"createdDate"`
		Editable         bool        `json:"editable"`
		LastExecutedDate float64     `json:"lastExecutedDate"`
		LastUpdatedDate  float64     `json:"lastUpdatedDate"`
		LocationID       string      `json:"locationId"`
		SceneColor       interface{} `json:"sceneColor"`
		SceneIcon        *string     `json:"sceneIcon"`
		SceneID          string      `json:"sceneId"`
		SceneName        string      `json:"sceneName"`
	} `json:"items"`
}
