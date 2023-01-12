import { useEffect, useState } from "react"
import { getBoard, getDices } from "../api/api"
import { getBadgeIcon } from "./common"


function Game() {
    return (
        <div className="game">
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
                return <BoardCell key={idx} nobles={1} badge={getBadgeIcon(el)} />
            })}
        </div>
    )
}

function BoardCell(props: { badge: string, nobles: number }) {
    return (
        <div className="boardCell">
            <Nobles amount={1} />
            <Cell imgSrc={props.badge} />
        </div >
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
                return <Cell key={idx} imgSrc={getBadgeIcon(dice)} />
            })}
        </div>
    )
}

export function Cell({ imgSrc }: { imgSrc: string }) {
    return (
        <div className="cell">
            <img src={imgSrc} />
        </div>
    )
}

function Nobles({ amount }: { amount: number }) {
    function renderNobles() {
        switch (amount) {
            case 0:
                return
            case 1:
                return <Noble />
            case 2:
                return (
                    <>
                        <Noble /><Noble />
                    </>
                )
            default:
                return null
        }
    }
    return (
        <div className="nobles">
            {renderNobles()}
        </div>
    )
}

function Noble() {
    return (
        <div>X</div>
    )
}


export default Game