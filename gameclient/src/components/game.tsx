import { MouseEventHandler, useEffect, useState } from "react"
import { getBoard, getDices } from "../api/api"
import BonusBoard from "./bonusboard"
import { Bonus, Domino, getBadgeIcon, Badge, GameState } from "./common"


function Game() {
    const [gameState, setGameState] = useState(WebSocket.CLOSED)
    const [statusMsg, setStatusMsg] = useState("")
    const [playerName, setPlayerName] = useState("")
    const [gameBoard, setGameBoard] = useState<Domino[][] | undefined>(undefined)
    const [bonusCard, setBonusCard] = useState<Bonus[] | undefined>(undefined)
    const [dices, setDices] = useState<Domino[] | undefined>(undefined)


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
                setGameBoard(board)
                setBonusCard(bonusCard)
                setStatusMsg(message)
                setDices(dices)
            }
        }
    }

    return (
        <>
            <div className="game">
                {statusMsg !== "" ? <StatusPane message={statusMsg} /> : ""}
                {gameState === WebSocket.OPEN && gameBoard !== undefined ? <>
                    <Board board={gameBoard} />
                    <BonusBoard bonusCard={bonusCard} />
                    <DiceSection dices={dices} />
                </> : ""}
                {gameState !== WebSocket.OPEN ? <Connect connectHandler={handleConnect} playerName={playerName} setPlayerName={setPlayerName} /> : ""}
            </div>
        </>
    )
}

function Board({ board }: { board?: Domino[][] }) {
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
            {props.elements.map(({ name, nobles }, idx) => {
                return <BoardCell key={idx} nobles={nobles} name={name} />
            })}
        </div>
    )
}

function BoardCell({ name, nobles, onClick }: { name: Badge, nobles: number, onClick?: MouseEventHandler }) {
    return (
        <div className="boardCell">
            <Nobles amount={nobles} />
            <Cell imgSrc={getBadgeIcon(name)} onClick={onClick} />
        </div >
    )
}

function DiceSection({ dices }: { dices: Domino[] | undefined }) {

    return (
        <div className="dice-section">
            {dices?.map(({ name, nobles }, idx) => {
                return <BoardCell key={idx} name={name} nobles={nobles} />
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