package game

// ============================================================================
// EVENT NAMES - Constantes para todos os eventos de batalha
// ============================================================================

const (
	// --- Battle Events ---

	// EventActionRequested é emitido quando um jogador clica em um botão de ação
	// Payload: ActionRequestedData
	EventActionRequested = "battle.action.requested"

	// EventActionExecuted é emitido após uma ação ser executada com sucesso
	// Payload: ActionExecutedData
	EventActionExecuted = "battle.action.executed"

	// EventTurnChanged é emitido quando o turno muda de personagem
	// Payload: TurnChangedData
	EventTurnChanged = "battle.turn.changed"

	// EventBattleEnded é emitido quando a batalha termina
	// Payload: BattleEndedData
	EventBattleEnded = "battle.ended"

	// EventCharacterDamaged é emitido quando um personagem toma dano
	// Payload: CharacterDamagedData
	EventCharacterDamaged = "battle.character.damaged"
)

// ============================================================================
// EVENT PAYLOADS - Structs tipados para dados de eventos de batalha
// ============================================================================

// ActionRequestedData contém os dados de uma ação solicitada pelo jogador
type ActionRequestedData struct {
	Actor  *Character
	Action *Action
	Target *Character // Pode ser nil se o targeting for decidido depois
}

// ActionExecutedData contém os dados de uma ação que foi executada
type ActionExecutedData struct {
	Actor  *Character
	Action *Action
	Target *Character
	Damage int // Ou outros resultados da ação
}

// TurnChangedData contém os dados de uma mudança de turno
type TurnChangedData struct {
	PreviousCharacter *Character
	CurrentCharacter  *Character
	TurnNumber        int
}

// BattleEndedData contém os dados do fim da batalha
type BattleEndedData struct {
	Winner *Character
	Loser  *Character
}

// CharacterDamagedData contém os dados de dano a um personagem
type CharacterDamagedData struct {
	Character   *Character
	Damage      int
	RemainingHP int
}
