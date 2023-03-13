import { useState, MouseEvent, memo } from "react"
import BonusBoard from "./bonusboard"
import { Bonus, Domino, GameState, ServerPayload, GameTurn, PlayerPower } from "./common"
import { Board } from "./board"
import { DiceSection } from "./dice"
import { PowerPrompt } from "./powerprompt"

function Game() {
    const [gameState, setGameState] = useState(WebSocket.CLOSED)
    const [wsConn, setWsConn] = useState<WebSocket | null>(null)
    const [statusMsg, setStatusMsg] = useState("")
    const [playerName, setPlayerName] = useState("")
    const [gameBoard, setGameBoard] = useState<Domino[][] | null>(null)
    const [bonusCard, setBonusCard] = useState<Bonus[] | null>(null)
    const [dices, setDices] = useState<Domino[] | null>(null)
    const [selectedDice, setSelectedDice] = useState<Domino[] | null>(null)
    const [gameTurn, setGameTurn] = useState<GameTurn>(0)
    const [power, setPower] = useState<PlayerPower>({
        type: 0,
        description: "",
        use: false,
        confirmed: false,
    })


    function SendServerData(payload: ServerPayload) {
        wsConn?.send(JSON.stringify(payload))
    }

    function clearGameState(ws: WebSocket) {
        setGameState(ws.readyState)
        setBonusCard(null)
        setGameBoard(null)
        setDices(null)
        setSelectedDice(null)
    }

    function handleConnect(ev: SubmitEvent) {
        ev.preventDefault()
        setGameState(WebSocket.CONNECTING)
        const ws = new WebSocket("ws://localhost:8080/join")
        setStatusMsg("Connecting...")

        ws.onopen = () => {
            setGameState(ws.readyState)
            const payload: ServerPayload = {
                name: playerName
            }
            ws.send(JSON.stringify(payload))

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
                const { board, bonusCard, message, dices, selectedDice, playerPower }: GameState = JSON.parse(d)
                board !== null && setGameBoard(board)
                bonusCard !== null && setBonusCard(bonusCard)
                message !== "" && setStatusMsg(message)
                dices !== null && setDices(dices)
                setSelectedDice(selectedDice)
                setPower(playerPower)
            }
        }
    }

    function handleDiceSelect(ev: MouseEvent<HTMLElement>) {
        const payload: ServerPayload = {
            selectedDie: Number(ev.currentTarget.id)
        }

        SendServerData(payload)
    }

    function handleBoardClick(ev: MouseEvent<HTMLElement>) {
        const payload: ServerPayload = {
            boardPosition: {
                row: Number(ev.currentTarget.parentElement?.parentElement?.id),
                cell: Number(ev.currentTarget.id)
            }
        }
        SendServerData(payload)
    }

    function handlePowerChoice(use: boolean) {
        const pwr: PlayerPower = { ...power, use: use, confirmed: true }
        setPower(pwr)
        SendServerData({ playerPower: pwr })
    }

    return (
        <>
            <div className="game">
                {statusMsg !== "" ? <StatusPane message={statusMsg} /> : ""}
                {power.type !== 0 && !power.confirmed && <PowerPrompt handlePowerChoice={handlePowerChoice} power={power} />}
                {gameState === WebSocket.OPEN && gameBoard !== undefined ? <>
                    <Board board={gameBoard} handleOnClick={handleBoardClick} />
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