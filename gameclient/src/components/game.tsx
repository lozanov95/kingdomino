import { MouseEventHandler, useEffect, useState } from "react"
import { getBoard, getDices } from "../api/api"
import BonusBoard from "./bonusboard"
import { Bonus, Domino, getBadgeIcon, Badge } from "./common"


function Game() {
    const [gameState, setGameState] = useState(WebSocket.CLOSED)
    const [statusMsg, setStatusMsg] = useState("")
    const [playerName, setPlayerName] = useState("")
    const [gameBoard, setGameBoard] = useState<Domino[][] | undefined>(undefined)
    const [bonusCard, setBonusCard] = useState<Bonus[] | undefined>(undefined)


    function handleConnect(ev: SubmitEvent) {
        ev.preventDefault()
        setGameState(WebSocket.CONNECTING)
        const ws = new WebSocket("ws://localhost:8080/join")
        setStatusMsg("Connecting...")

        ws.onopen = () => {
            setGameState(ws.readyState)
            ws.send(JSON.stringify({ name: playerName }))

            setStatusMsg("")
        }

        ws.onerror = () => {
            setGameState(WebSocket.CLOSED)
            setStatusMsg("Connection to the server failed.")
        }

        ws.onclose = () => {
            setGameState(ws.readyState)
            setStatusMsg("The connection was closed.")
        }

        ws.onmessage = ({ data }: { data: string }) => {
            const d: string = data
            if (d.length > 0) {
                const { board, bonusCard } = JSON.parse(d)
                setGameBoard(board)
                setBonusCard(bonusCard)
            }
        }
    }


    return (
        <div className="game">
            <StatusPane message={statusMsg} />
            {(() => {
                switch (gameState) {
                    case WebSocket.OPEN:
                        return (<>
                            <Board board={gameBoard} />
                            <BonusBoard bonusCard={bonusCard} />
                            <DiceSection />
                        </>)
                    default:
                        return <Connect connectHandler={handleConnect} playerName={playerName} setPlayerName={setPlayerName} />
                }
            })()}
        </div>
    )
}

function Board({ board }: { board?: Domino[][] }) {
    // const [board, setBoard] = useState<Domino[][] | null>(null)

    useEffect(() => {

    }, [board])

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

function DiceSection() {
    const [dices, setDices] = useState<Domino[] | null>(null)

    useEffect(() => {
        setDices(getDices())
    }, [])
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
        <form onSubmitCapture={connectHandler}>
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