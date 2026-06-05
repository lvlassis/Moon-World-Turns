package scenes

type SceneManager struct {
	scenes       map[string]Scene
	currentScene Scene
}

func NewSceneManager() *SceneManager {
	return &SceneManager{
		scenes: make(map[string]Scene),
	}
}

func (sm *SceneManager) AddScene(id string, scn Scene) {
	sm.scenes[id] = scn
}

func (sm *SceneManager) SetCurrentScene(id string) {
	scn, exists := sm.scenes[id]
	if !exists {
		panic("Scene not found: " + id)
	}
	sm.currentScene = scn
}

func (sm *SceneManager) GetCurrentScene() Scene {
	return sm.currentScene
}

