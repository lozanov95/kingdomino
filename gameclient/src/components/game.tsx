import { useEffect, useState } from "react"
import { getBoard, getDices } from "../api/api"

function Game() {
    return (
        <div>
            <h2>Kingdomino</h2>
            <Board />
            <DiceSection />
        </div>
    )
}

function Board() {
    const [board, setBoard] = useState([[0]])
    const newBoard = getBoard()

    if (newBoard.toString() !== board.toString()) {
        setBoard(newBoard)
    }

    return (
        <div className="board center" >
            {board.map((el, idx) => {
                return <Row key={idx} elements={el} />
            })}
        </div>
    )
}

function Row(props: { elements: number[] }) {
    return (
        <div className="row">
            {props.elements.map((el, idx) => {
                return <Cell key={idx} text={el} />
            })}
        </div>
    )
}

function Cell(props: { text: string | number }) {
    return (
        <div className="cell">{props.text}</div>
    )
}

function DiceSection() {
    const [dices, setDices] = useState([0])
    const newDice = getDices()

    if (newDice.toString() !== dices.toString()) {
        setDices(newDice)
    }

    return (
        <div className="dice-section">
            {dices.map((dice, idx) => {
                return <Dice key={idx} value={dice} />
            })}
        </div>
    )
}

function Dice(props: { value: number }) {
    return (
        <div className="cell">
            {props.value}
        </div>
    )
}

export default Game