import { EventHandler, FormEventHandler, useEffect, useState } from "react"

enum Player {
    You = 0,
    Other,
}

type Message = { content: string, player: Player }

function Chat() {
    const [connected, setConnected] = useState(false)
    const [wsConnection, setWsConnection] = useState<WebSocket | null>(null)
    const [messages, setMessages] = useState<Message[]>([])

    function handleChatConnect() {
        connect()
    }

    useEffect(() => { }, [messages])

    const connect = () => {
        const ws = new WebSocket("ws://localhost:8080/ws")
        ws.onopen = () => {
            setConnected(true)
            setWsConnection(ws)
        }

        ws.onclose = () => {
            setConnected(false)
            setMessages([])
        }

        ws.onerror = () => {
            setConnected(false)
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
        <div>
            <h1>Chat</h1>

            {connected ?
                <div><MessageContainer messages={messages} /><MessageSendContainer handleSend={sendMessage} /></div> :
                <button hidden={connected} onClick={handleChatConnect}>Connect</button>
            }

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

