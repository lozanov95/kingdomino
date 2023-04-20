package game

import (
	"encoding/json"
	"errors"
	"io"
	"log"
	"sync"
	"time"
)

var (
	ErrPlayerDisconnected = errors.New("the player has been disconnected")
)

type DiePos struct {
	Cell int `json:"cell"`
	Row  int `json:"row"`
}

type Player struct {
	Id                 int64  `json:"id"`
	Name               string `json:"name"`
	Conn               Connectionable
	Board              *Board
	BonusCard          *BonusMap
	Connected          bool
	GameState          chan GameState
	ClientMsg          chan string
	Dices              []DiceResult
	SelectedCoatOfArms BadgeName
}

type Connectionable interface {
	Read([]byte) (int, error)
	Write([]byte) (int, error)
	SetReadDeadline(time.Time) error
	Close() error
}

type Scoreboard struct {
	PlayerName string       `json:"name"`
	Scores     []BadgeScore `json:"scores"`
	TotalScore int          `json:"totalScore"`
}

type BadgeScore struct {
	Badge BadgeName `json:"badge"`
	Score int       `json:"score"`
}

// Creates a new player instance and returns a pointer to it.
func NewPlayer(conn Connectionable) *Player {

	player := &Player{
		Id:        time.Now().UnixNano(),
		Name:      "",
		Conn:      conn,
		Board:     NewBoard(),
		BonusCard: NewBonusMap(),
		Connected: true,
		GameState: make(chan GameState),
		ClientMsg: make(chan string),
	}

	return player
}

// Increases the bonus of a specific card
func (p *Player) IncreaseBonus(b Dice) {
	if b.Nobles != 0 {
		return
	}

	tmp := (*p.BonusCard)[b.Name]
	tmp.Increment()
	(*p.BonusCard)[b.Name] = tmp
}

func (p *Player) GameStateLoop() {
	for send := range p.GameState {
		msg, err := json.Marshal(send)
		if err != nil {
			if err == io.EOF {
				p.Connected = false
				return
			}
			log.Println(err)
		}
		p.Conn.Write(msg)
	}
}

func (p *Player) GetInput() ClientPayload {
	buf := make([]byte, 1024)
	for {
		err := p.Conn.SetReadDeadline(time.Now().Add(TIMEOUT))
		if err != nil {
			p.Disconnect()
			panic(err)
		}
		n, err := p.Conn.Read(buf[0:])
		if err != nil {
			p.Disconnect()
			panic(ErrPlayerDisconnected)
		}

		payload := ClientPayload{
			DiePos:      DiePos{-1, -1},
			SelectedDie: -1,
		}

		err = json.Unmarshal(buf[:n], &payload)
		if err != nil {
			log.Println("failed to parse output")
			continue
		}

		return payload
	}
}

func (p *Player) SendMessage(message string) {
	p.GameState <- GameState{Message: message}
}

func (p *Player) SendDice(d *[]DiceResult, m string) {
	p.GameState <- GameState{Dices: d, Message: m}
}

func (p *Player) SendGameState(d *[]DiceResult, m string) {
	p.GameState <- GameState{
		ID:        p.Id,
		Message:   m,
		Board:     p.Board,
		BonusCard: p.BonusCard,
		Dices:     d,
	}
}

func (p *Player) SendPlayerPowerPrompt(d *[]DiceResult, pp PlayerPower) {
	p.GameState <- GameState{PlayerPower: pp, Dices: d}
}

func (p *Player) SendScoreboard(p1Score, p2Score *Scoreboard, m string) {
	p.GameState <- GameState{Scoreboards: []Scoreboard{*p1Score, *p2Score}, Message: m}
}

// Places domino on the field, following the placement rules
func (p *Player) PlaceDomino(d *[]DiceResult) {
	p.SendGameState(d, "Select the dice that you want to place")
	choice := p.getSelectedDominoChoice(d)
	prevPos := p.placeOnBoard(d, choice, BoardPlacementInput{
		Board:                 p.Board,
		IgnoreConnectionRules: p.handleIgnoreConnectionRulesPower(),
	})

	p.SendGameState(d, "Select the dice that you want to place")
	choice = p.getSelectedDominoChoice(d)
	p.placeOnBoard(d, choice, BoardPlacementInput{PrevPosition: prevPos, Board: p.Board})
	p.SendGameState(d, "Waiting for all players to complete their turns.")
}

// Allows the user to place 2 separate dominos
func (p *Player) PlaceSeparatedDomino(d *[]DiceResult) {
	p.SendGameState(d, "Select the dice that you want to place")
	choice := p.getSelectedDominoChoice(d)
	p.placeOnBoard(d, choice, BoardPlacementInput{Board: p.Board, SeparateDice: true})

	p.SendGameState(d, "Select the dice that you want to place")
	choice = p.getSelectedDominoChoice(d)
	p.placeOnBoard(d, choice, BoardPlacementInput{Board: p.Board, SeparateDice: true})
	p.SendGameState(d, "Waiting for all players to complete their turns.")
}

func (p *Player) getSelectedDominoChoice(dr *[]DiceResult) int {
	for {
		var choice int
		payload := p.GetInput()

		choice = payload.SelectedDie

		if !isDicePlaceChoiceValid(dr, choice, p.Id) {
			p.SendMessage("Invalid choice!")
			continue
		}

		return choice
	}
}

func (p *Player) placeOnBoard(d *[]DiceResult, choice int, b BoardPlacementInput) DiePos {
	pos, err := p.getBoardPlacementInput(b)
	if err != nil {
		return DiePos{}
	}

	p.Board[pos.Row][pos.Cell] = *(*d)[choice].Dice
	(*d)[choice].IsPlaced = true

	return pos
}

func (p *Player) getBoardPlacementInput(bpi BoardPlacementInput) (DiePos, error) {
	p.SendMessage("Select the place on the board that you want to place it on")
	for p.Connected {
		payload := p.GetInput()

		if bpi.IgnoreConnectionRules && bpi.IsValid(&payload.DiePos) {
			return payload.DiePos, nil
		}
		if !p.Board.IsThereOccupiedNeighbourCell(payload.DiePos.Row, payload.DiePos.Cell) || !bpi.IsValid(&payload.DiePos) {
			p.SendMessage("Invalid position. Please select a new position")
			continue
		}

		return payload.DiePos, nil
	}
	return DiePos{}, io.EOF
}

// Returns TRUE if there is a free spot
func (p *Player) IsThereAFreeSpot() bool {
	for i := 0; i < len(p.Board); i++ {
		for j := 0; j < len(p.Board[i]); j++ {
			if p.Board[i][j].Name == EMPTY {
				return true
			}
		}
	}

	return false
}

func (p *Player) IsValidPlacementPossible() bool {
	for i := 0; i < len(p.Board); i++ {
		for j := 0; j < len(p.Board[i]); j++ {
			if p.Board[i][j].Name == EMPTY && p.Board.IsThereFreeNeighbourCell(i, j) {
				return true
			}
		}
	}

	return false
}

// Checks if there is a way to place a domino, either with the separate bonus (if it is completed) or not
func (p *Player) IsAnyPlacementPossible() bool {
	if p.IsBonusCompleted(PWRSeparateDominos) && p.IsThereAFreeSpot() {
		return true
	}

	return p.IsValidPlacementPossible()
}

func (p *Player) CalculateScore() Scoreboard {
	badges := []BadgeName{DOT, LINE, DOUBLEDOT, DOUBLELINE, CHECKED, FILLED}
	pChan := make(chan BadgeScore, 6)
	var wg sync.WaitGroup
	for _, badge := range badges {
		wg.Add(1)
		go func(badge BadgeName) {
			defer wg.Done()
			pts, dms := p.Board.CalculateBadgePoints(badge)
			if p.SelectedCoatOfArms == badge {
				pts += dms * 3
			}
			pChan <- BadgeScore{Badge: badge, Score: pts}
		}(badge)
	}

	sb := Scoreboard{
		PlayerName: p.Name,
		Scores:     make([]BadgeScore, 6),
	}

	wg.Wait()
	close(pChan)

	for p := range pChan {
		sb.Scores[int(p.Badge)-2] = p
		sb.TotalScore += p.Score
	}

	return sb
}

func (p *Player) UsePower(pt PowerType) {
	p.BonusCard.MarkUsed(pt)
}

func (p *Player) Disconnect() {
	p.Connected = false
}

// Returns true if the player have collected the required amount badges and haven't still used the bonus
func (p *Player) IsBonusUsable(pt PowerType) bool {
	return p.BonusCard.IsBonusUsable(pt)
}

// Does the player have collected the required amount of badges to unlock the specific bonus
func (p *Player) IsBonusCompleted(pt PowerType) bool {
	return p.BonusCard.IsBonusCompleted(pt)
}

func (p *Player) GetName() string {
	return p.Name
}

func (p *Player) AddBonus(b Dice) {
	p.BonusCard.AddBonus(b)
}

func (p *Player) SetBonusIneligible(b Dice) {
	bonus := (*p.BonusCard)[b.Name]
	bonus.Eligible = false
	(*p.BonusCard)[b.Name] = bonus
}

func (p *Player) IsBonusEligible(b Dice) bool {
	return (*p.BonusCard)[b.Name].Eligible
}

func (p *Player) handleUseAddNoblePower() {
	if !p.IsBonusUsable(PWRAddNoble) {
		return
	}

	p.SendGameState(nil, "Select a badge on your board that you will add a noble to.")
	payload := func() *ClientPayload {
		for {
			payload := p.GetInput()

			if p.Board.isCellOccupied(payload.DiePos.Row, payload.DiePos.Cell) &&
				p.Board[payload.DiePos.Row][payload.DiePos.Cell].Name != CASTLE {
				return &payload
			}
			p.SendGameState(nil, "Invalid choice! Select a badge on your board that you will add a noble to.")
		}
	}()
	p.UsePower(PWRAddNoble)
	b := p.Board[payload.DiePos.Row][payload.DiePos.Cell]
	b.Nobles++
	p.Board[payload.DiePos.Row][payload.DiePos.Cell] = b
}

func (p *Player) GetPlayerPowerChoice() bool {
	for {
		payload := p.GetInput()

		if payload.PlayerPower.Confirmed {
			return payload.PlayerPower.Use
		}
	}
}

// If the IgnoreConnectionRules power is available, prompts the player and asks if it should be played.
func (p *Player) handleIgnoreConnectionRulesPower() bool {
	if !p.IsBonusUsable(PWRNoConnectionRules) {
		return false
	}

	p.SendPlayerPowerPrompt(nil, PlayerPower{
		Type:        PWRNoConnectionRules,
		Description: "You can play your domino without following the Connection Rules",
	})

	if p.GetPlayerPowerChoice() {
		p.UsePower(PWRNoConnectionRules)
		return true
	}

	return false
}

func (p *Player) handleAddDomainPointsPower() {
	if !p.IsBonusUsable(PWRDomainPoints) {
		return
	}

	p.SendGameState(nil, "Choose a coat of arms. Each different	DOMAIN with this coat of arms will earn you	3 prestige points at the end of the game.")
	payload := func() *ClientPayload {
		for {
			payload := p.GetInput()

			if p.Board.isCellOccupied(payload.DiePos.Row, payload.DiePos.Cell) &&
				p.Board[payload.DiePos.Row][payload.DiePos.Cell].Name != CASTLE {
				return &payload
			}
			p.SendGameState(nil, "Invalid choice! Select a coat of arms that you want to get prestige points from.")
		}
	}()

	p.UsePower(PWRDomainPoints)
	p.SelectedCoatOfArms = p.Board[payload.DiePos.Row][payload.DiePos.Cell].Name
}

func (p *Player) SelectDie(dr *DiceResult) {
	dr.IsPicked = true
	dr.PlayerId = p.Id
	p.AddBonus(*dr.Dice)
}

func (p *Player) SelectQuestionmarkDie(dr *DiceResult) {
	dr.IsPicked = true
	dr.PlayerId = p.Id
}
