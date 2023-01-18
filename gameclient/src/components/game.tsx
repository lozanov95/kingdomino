import React, { MouseEventHandler, useEffect, useState, MouseEvent, useMemo } from "react"
import BonusBoard from "./bonusboard"
import { Bonus, Domino, getBadgeIcon, Badge, GameState } from "./common"

function GS() {

}

function Game() {
    const [gameState, setGameState] = useState(WebSocket.CLOSED)
    const [wsConn, setWsConn] = useState<WebSocket | null>(null)
    const [statusMsg, setStatusMsg] = useState("")
    const [playerName, setPlayerName] = useState("")
    const [gameBoard, setGameBoard] = useState<Domino[][] | undefined>(undefined)
    const [bonusCard, setBonusCard] = useState<Bonus[] | undefined>(undefined)
    const [dices, setDices] = useState<Domino[] | undefined>(undefined)
    const [pTurn, setPTurn] = useState(-1)


    function clearGameState(ws: WebSocket) {
        setGameState(ws.readyState)
        setBonusCard(undefined)
        setGameBoard(undefined)
        setDices(undefined)
    }

    function handleConnect(ev: SubmitEvent) {
        ev.preventDefault()
        setGameState(WebSocket.CONNECTING)
        const ws = new WebSocket("ws://localhost:8080/join")
        setStatusMsg("Connecting...")

        ws.onopen = () => {
            setGameState(ws.readyState)
            ws.send(JSON.stringify({ name: playerName }))
            setStatusMsg("Waiting for opponent.")
            setWsConn(ws)
        }

        ws.onerror = () => {
            clearGameState(ws)
            setStatusMsg("Connection to the server failed.")
        }

        ws.onclose = () => {
            clearGameState(ws)
            setStatusMsg("The connection was closed.")
        }

        ws.onmessage = ({ data }: { data: string }) => {
            const d: string = data
            if (d.length > 0) {
                const { board, bonusCard, message, dices }: GameState = JSON.parse(d)
                board !== null ? setGameBoard(board) : ""
                bonusCard !== null ? setBonusCard(bonusCard) : ""
                message !== "" ? setStatusMsg(message) : ""
                dices !== null ? setDices(dices) : ""
            }
        }
    }

    function handleDiceSelect(ev: MouseEvent<HTMLElement>) {
        wsConn?.send(ev.currentTarget.id)
    }

    return (
        <>
            <div className="game">
                {statusMsg !== "" ? <StatusPane message={statusMsg} /> : ""}
                {gameState === WebSocket.OPEN && gameBoard !== undefined ? <>
                    <Board board={gameBoard} />
                    <BonusBoard bonusCard={bonusCard} />
                    <DiceSection dices={dices} handleDiceSelect={handleDiceSelect} />
                </> : ""}
                {gameState !== WebSocket.OPEN ? <Connect connectHandler={handleConnect} playerName={playerName} setPlayerName={setPlayerName} /> : ""}
            </div>
        </>
    )
}


const Board = React.memo(
    function Board({ board }: { board?: Domino[][] }) {

        const currentBoard = useMemo(() => {
            return board
        }, [board])

        return (
            <div className="board center" >
                {currentBoard?.map((el, idx) => {
                    return <Row key={idx} elements={el} />
                })}
            </div>
        )
    }
)


function Row(props: { elements: Domino[] }) {
    return (
        <div className="row">
            {props.elements.map(({ name, nobles }, idx) => {
                return <BoardCell id={idx.toString()} key={idx} nobles={nobles} name={name} />
            })}
        </div>
    )
}

function BoardCell({ id, name, nobles, onClick }: { id: string, name: Badge, nobles: number, onClick?: MouseEventHandler }) {
    return (
        <div className="boardCell">
            <Nobles amount={nobles} />
            <Cell id={id} imgSrc={getBadgeIcon(name)} onClick={onClick} />
        </div >
    )
}

const DiceSection = React.memo(
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

export function Cell({ id, imgSrc, onClick }: { id: string, imgSrc: string, onClick?: MouseEventHandler }) {
    return (
        <div className="cell" id={id} onClick={onClick}>
            <img src={imgSrc} alt="badge icon" />
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

function Connect({ connectHandler, playerName, setPlayerName }: { connectHandler: any, playerName: string, setPlayerName: any }) {

    return (
        <form onSubmitCapture={connectHandler} className="connectForm">
            <h2>Enter your name</h2>
            <input placeholder="name" minLength={3} value={playerName} onChange={e => setPlayerName(e.target.value)} />
            <button>Connect</button>
        </form>
    )
}

function StatusPane({ message }: { message: string }) {
    return (
        <div className="status">{message}</div>
    )
}

export default Game