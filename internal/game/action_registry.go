package game


type ActionRegistry struct {
	actions map[string]*Action
}

var actionRegistry *ActionRegistry

func GetActionRegistry() *ActionRegistry {
	if actionRegistry == nil {
		actionRegistry = &ActionRegistry{
			actions: make(map[string]*Action),
		}
		actionRegistry.LoadActions()
	}
	return actionRegistry
}

func (ar *ActionRegistry) Register(action *Action) {
	ar.actions[action.Id] = action
}

func (ar *ActionRegistry) Get(id string) (*Action, bool) {
	action, exists := ar.actions[id]
	return action, exists
}

