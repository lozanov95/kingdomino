import { useState, MouseEvent, memo } from "react"
import BonusBoard from "./bonusboard"
import { Bonus, Domino, GameState } from "./common"
import { Board } from "./board"
import { DiceSection } from "./dice"

function Game() {
    const [gameState, setGameState] = useState(WebSocket.CLOSED)
    const [wsConn, setWsConn] = useState<WebSocket | null>(null)
    const [statusMsg, setStatusMsg] = useState("")
    const [playerName, setPlayerName] = useState("")
    const [gameBoard, setGameBoard] = useState<Domino[][] | undefined>(undefined)
    const [bonusCard, setBonusCard] = useState<Bonus[] | undefined>(undefined)
    const [dices, setDices] = useState<Domino[] | null>(null)
    const [selectedDice, setSelectedDice] = useState<Domino[] | null>(null)


    function clearGameState(ws: WebSocket) {
        setGameState(ws.readyState)
        setBonusCard(undefined)
        setGameBoard(undefined)
        setDices(null)
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
                const { board, bonusCard, message, dices, selectedDice }: GameState = JSON.parse(d)
                board !== null ? setGameBoard(board) : ""
                bonusCard !== null ? setBonusCard(bonusCard) : ""
                message !== "" ? setStatusMsg(message) : ""
                dices !== null ? setDices(dices) : ""
                setSelectedDice(selectedDice)
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
                    <DiceSection selectedDice={selectedDice} dices={dices} handleDiceSelect={handleDiceSelect} />
                </> : ""}
                {gameState !== WebSocket.OPEN ? <Connect connectHandler={handleConnect} playerName={playerName} setPlayerName={setPlayerName} /> : ""}
            </div>
        </>
    )
}

function Connect({ connectHandler, playerName, setPlayerName }: { connectHandler: any, playerName: string, setPlayerName: any }) {

    return (
        <form onSubmitCapture={connectHandler} className="connectForm">
            <h2>Enter your name</h2>
            <input placeholder="name" minLength={3} value={playerName} onChange={e => setPlayerName(e.target.value)} required />
            <button>Connect</button>
        </form>
    )
}

const StatusPane = memo(
    function StatusPane({ message }: { message: string }) {
        return (
            <div className="status">{message}</div>
        )
    }
)

export default Game