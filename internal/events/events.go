package events

// ============================================================================
// EVENT NAMES - Constantes para eventos de sistema/view
// ============================================================================
// Nota: Eventos de Battle estão em internal/game/events.go para evitar import cycle

const (
	// --- Scene Events ---

	// EventSceneChangeRequested é emitido quando uma mudança de cena é solicitada
	// Payload: SceneChangeRequestedData
	EventSceneChangeRequested = "scene.change.requested"
)

// ============================================================================
// EVENT PAYLOADS - Structs tipados para dados de eventos de sistema
// ============================================================================

// SceneChangeRequestedData contém os dados de uma solicitação de mudança de cena
type SceneChangeRequestedData struct {
	FromScene string
	ToScene   string
	Data      map[string]interface{} // Dados opcionais para passar entre cenas
}
