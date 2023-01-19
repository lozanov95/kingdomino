import { memo, MouseEventHandler } from "react"
import { Domino } from "./common"
import { BoardCell } from "./board"

export const DiceSection = memo(
    function DiceSection({ dices, selectedDice, handleDiceSelect }: { dices?: Domino[] | null, selectedDice?: Domino[] | null, handleDiceSelect: MouseEventHandler }) {

        return (
            <div className="dice-section">
                <div className="available">
                    {dices ? <h2>Available dice</h2> : ""}
                    {dices?.map(({ name, nobles }, idx) => {
                        return <BoardCell id={idx.toString()} key={idx} name={name} nobles={nobles} onClick={handleDiceSelect} />
                    })}
                </div>
                <div className="selected-dice">
                    {selectedDice ? <h2>Selected</h2> : ""}
                    {selectedDice?.map(({ name, nobles }, idx) => {
                        return <BoardCell id={idx.toString()} key={idx} name={name} nobles={nobles} onClick={handleDiceSelect} />
                    })}
                </div>
            </div>
        )
    }
)
