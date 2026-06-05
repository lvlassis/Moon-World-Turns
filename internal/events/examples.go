package events

// Este arquivo contém exemplos de uso do Event Bus
// Você pode deletar este arquivo após entender como usar

/*
// ============================================================================
// EXEMPLO 1: Emitindo um evento de ação
// ============================================================================

// No action_btn.go
import (
    "lvlassis/moon-world-turns/internal/game"
    "lvlassis/moon-world-turns/internal/events"
)

func NewActionButton(action *game.Action, character *game.Character, bus *events.Bus) *gui.Button {
    action_btn := gui.NewButton(action.Name)
    action_btn.Subscribe(gui.OnClick, func(name string, ev interface{}) {
        // Emite evento type-safe (constantes e payloads estão em game.*)
        events.Emit(bus, game.EventActionRequested, game.ActionRequestedData{
            Actor:  character,
            Action: action,
            Target: nil, // Será definido pelo listener
        })
    })
    return action_btn
}

// ============================================================================
// EXEMPLO 2: Escutando eventos
// ============================================================================

// No battle_scene.go (dentro de NewBattleScene ou setupGUI)
func (bs *BattleScene) setupEventListeners() {
    // Listener type-safe - o payload já vem tipado!
    events.On(bs.bus, game.EventActionRequested, func(data game.ActionRequestedData) {
        // data já é do tipo ActionRequestedData, sem precisar de cast!
        target := bs.battle.GetCharacter2() // Lógica de targeting
        bs.battle.DoAction(data.Actor, data.Action, target)

        // Pode emitir outro evento após execução
        events.Emit(bs.bus, game.EventActionExecuted, game.ActionExecutedData{
            Actor:  data.Actor,
            Action: data.Action,
            Target: target,
            Damage: 50, // Valor real do dano
        })
    })

    // Outro exemplo: animações baseadas em eventos
    events.On(bs.bus, game.EventCharacterDamaged, func(data game.CharacterDamagedData) {
        // Tocar animação de dano no personagem
        fmt.Printf("%s tomou %d de dano! HP restante: %d\n",
            data.Character.Name, data.Damage, data.RemainingHP)
    })
}

// ============================================================================
// EXEMPLO 3: Múltiplos listeners para o mesmo evento
// ============================================================================

// No battle.go - emitir evento quando personagem toma dano
func (b *Battle) DoAction(actor *game.Character, action *game.Action, target *game.Character) {
    // ... lógica de dano ...
    damage := 50
    target.Life -= damage

    // Emite evento para quem quiser escutar
    events.Emit(b.bus, game.EventCharacterDamaged, game.CharacterDamagedData{
        Character:   target,
        Damage:      damage,
        RemainingHP: target.Life,
    })
}

// Listener 1: Atualizar HUD
events.On(bus, game.EventCharacterDamaged, func(data game.CharacterDamagedData) {
    healthBar.Update(data.RemainingHP)
})

// Listener 2: Tocar efeito sonoro
events.On(bus, game.EventCharacterDamaged, func(data game.CharacterDamagedData) {
    audio.Play("hit.mp3")
})

// Listener 3: Verificar morte
events.On(bus, game.EventCharacterDamaged, func(data game.CharacterDamagedData) {
    if data.RemainingHP <= 0 {
        events.Emit(bus, game.EventBattleEnded, game.BattleEndedData{
            Winner: otherCharacter,
            Loser:  data.Character,
        })
    }
})

*/
