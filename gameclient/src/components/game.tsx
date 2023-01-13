import { MouseEventHandler, useEffect, useState } from "react"
import { getBoard, getDices } from "../api/api"
import { Domino, getBadgeIcon } from "./common"


function Game() {
    return (
        <div className="game">
            <Board />
            <DiceSection />
        </div>
    )
}

function Board() {
    const [board, setBoard] = useState<Domino[][] | null>(null)

    useEffect(() => {
        setBoard(() => getBoard())
    }, [])

    return (
        <div className="board center" >
            {board?.map((el, idx) => {
                return <Row key={idx} elements={el} />
            })}
        </div>
    )
}

function Row(props: { elements: Domino[] }) {
    return (
        <div className="row">
            {props.elements.map(({ badge, nobles }, idx) => {
                return <BoardCell key={idx} nobles={nobles} badge={getBadgeIcon(badge)} />
            })}
        </div>
    )
}

function BoardCell({ badge, nobles, onClick }: { badge: string, nobles: number, onClick?: MouseEventHandler }) {
    return (
        <div className="boardCell">
            <Nobles amount={nobles} />
            <Cell imgSrc={badge} onClick={onClick} />
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

export function Cell({ imgSrc, onClick }: { imgSrc: string, onClick?: MouseEventHandler }) {
    return (
        <div className="cell" onClick={onClick}>
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