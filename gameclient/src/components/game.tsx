import { useEffect, useState } from "react"
import { getBoard, getDices } from "../api/api"
import badgeEmpty from "../assets/empty.svg"
import badgeCastle from "../assets/castle.svg"
import badgeChecked from "../assets/checked.svg"
import badgeFilled from "../assets/filled.svg"
import badgeDot from "../assets/dot.svg"
import badgeDoubleDot from "../assets/doubledot.svg"
import badgeLine from "../assets/line.svg"
import badgeDoubleLine from "../assets/doubleline.svg"
import badgeQuestion from "../assets/question.svg"

enum Badge {
    EMPTY = 0,
    CASTLE,
    DOT,
    LINE,
    DOUBLEDOT,
    DOUBLELINE,
    FILLED,
    CHECKED,
    QUESTIONMARK,
}

function Game() {
    return (
        <div>
            <h2>Kingdomino</h2>
            <Board />
            <DiceSection />
        </div>
    )
}

function GetBadgeIcon(id: number) {
    switch (id) {
        case Badge.EMPTY:
            return badgeEmpty
        case Badge.CASTLE:
            return badgeCastle
        case Badge.DOT:
            return badgeDot
        case Badge.LINE:
            return badgeLine
        case Badge.DOUBLEDOT:
            return badgeDoubleDot
        case Badge.DOUBLELINE:
            return badgeDoubleLine
        case Badge.FILLED:
            return badgeFilled
        case Badge.CHECKED:
            return badgeChecked
        case Badge.QUESTIONMARK:
            return badgeQuestion
        default:
            break;
    }
}

function Board() {
    const [board, setBoard] = useState([[0]])
    const newBoard = getBoard()

    if (newBoard.toString() !== board.toString()) {
        setBoard(newBoard)
    }

    return (
        <div className="board center" >
            {newBoard.map((el, idx) => {
                return <Row key={idx} elements={el} />
            })}
        </div>
    )
}

function Row(props: { elements: number[] }) {
    return (
        <div className="row">
            {props.elements.map((el, idx) => {
                return <Cell key={idx} badge={GetBadgeIcon(el) || ""} />
            })}
        </div>
    )
}

function Cell(props: { badge: string }) {
    return (
        <div className="cell"><img src={props.badge}></img></div>
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
            <img src={GetBadgeIcon(props.value)} />
        </div>
    )
}

export default Game