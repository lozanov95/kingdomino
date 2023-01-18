import { useEffect, useState } from "react"

enum Player {
    You = 0,
    Other,
}

type Message = { content: string, player: Player }

function Chat() {
    const [connectionState, setConnectionState] = useState<number>(WebSocket.CLOSED)
    const [wsConnection, setWsConnection] = useState<WebSocket | null>(null)
    const [messages, setMessages] = useState<Message[]>([])

    function handleChatConnect() {
        connect()
    }

    useEffect(() => {
        console.log(wsConnection)
    }, [wsConnection?.readyState])

    const connect = () => {
        const ws = new WebSocket("ws://localhost:8080/ws")
        setConnectionState(WebSocket.CONNECTING)
        ws.onopen = () => {
            setConnectionState(WebSocket.OPEN)
            setWsConnection(ws)
        }

        ws.onclose = () => {
            setConnectionState(WebSocket.CLOSED)
            setMessages([])
        }

        ws.onerror = () => {
            setConnectionState(WebSocket.CLOSED)
            setMessages([])
        }

        ws.onmessage = (ev) => {
            setMessages((msg) => [...msg, { content: ev.data, player: Player.Other }])
        }
    }

    function sendMessage(newMsg: string) {
        if (newMsg.length == 0) {
            return
        }
        wsConnection?.send(newMsg)
        setMessages((msg) => [...msg, { content: newMsg, player: Player.You }])
    }

    return (
        <div className="chat">
            <h1>Chat</h1>
            {connectionState === WebSocket.CONNECTING ? <div>Connecting to chat...</div > : ""}
            {connectionState === WebSocket.OPEN ? <div><MessageContainer messages={messages} /><MessageSendContainer handleSend={sendMessage} /></div> : ""}
            {connectionState === WebSocket.CLOSED ? <button onClick={handleChatConnect}>Connect</button> : ""}
        </div>
    )
}

function MessageSendContainer(props: { handleSend: (msg: string) => void }) {
    const [msg, setMsg] = useState("")

    function handleSubmit(ev: any) {
        ev.preventDefault()
        props.handleSend(msg)
        setMsg("")
    }

    return (
        <form onSubmit={e => handleSubmit(e)}>
            <input value={msg} onChange={e => setMsg(e.target.value)} />
            <button>Send</button>
        </form>

    )
}

function MessageContainer({ messages }: { messages: Message[] }) {
    return (
        <div className="msg-container">

            {messages.map((msg, idx) => {

                if (msg.player === Player.You) {
                    return <MessageBox key={idx} msg={msg} cls="chat-msg chat-you" />
                }
                else {
                    return <MessageBox key={idx} msg={msg} cls="chat-msg chat-other" />
                }
            })}
        </div>
    )
}

function MessageBox({ msg, cls }: { msg: Message, cls: string }) {
    const { content } = msg

    return (
        <div>
            < div className={cls} > {content}</div >
        </div>
    )
}

export default Chat

