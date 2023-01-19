import { memo, MouseEventHandler } from "react"
import { Domino } from "./common"
import { BoardCell } from "./board"

export const DiceSection = memo(
    function DiceSection({ dices, handleDiceSelect }: { dices: Domino[] | undefined, handleDiceSelect: MouseEventHandler }) {

        return (
            <div className="dice-section">
                {dices?.map(({ name, nobles }, idx) => {
                    return <BoardCell id={idx.toString()} key={idx} name={name} nobles={nobles} onClick={handleDiceSelect} />
                })}
            </div>
        )
    }
)
