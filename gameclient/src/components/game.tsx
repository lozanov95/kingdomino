import { useState, MouseEvent, memo } from "react";
import BonusBoard from "./bonusboard";
import {
  Bonus,
  Domino,
  GameState,
  ServerPayload,
  PlayerPower,
  Scoreboard,
  DiceResult,
} from "./common";
import { Board } from "./board";
import { DiceSection } from "./dice";
import { PowerPrompt } from "./powerprompt";
import { RulesSection } from "./rules";
import { ScoreSection } from "./scoreboard";

function Game() {
  const [gameState, setGameState] = useState<number>(WebSocket.CLOSED);
  const [wsConn, setWsConn] = useState<WebSocket | null>(null);
  const [statusMsg, setStatusMsg] = useState("");
  const [playerName, setPlayerName] = useState("");
  const [gameBoard, setGameBoard] = useState<Domino[][] | null>(null);
  const [bonusCard, setBonusCard] = useState<Bonus[] | null>(null);
  const [scoreboards, setScoreboards] = useState<Scoreboard[] | null>(null);
  const [dices, setDices] = useState<DiceResult[] | null>(null);
  const [power, setPower] = useState<PlayerPower>({
    type: 0,
    description: "",
    use: false,
    confirmed: false,
  });
  const [playerId, setPlayerId] = useState<number>(0)

  function SendServerData(payload: ServerPayload) {
    wsConn?.send(JSON.stringify(payload));
  }

  function clearGameState(ws: WebSocket) {
    setGameState(ws.readyState);
    setBonusCard(null);
    setGameBoard(null);
    setDices(null);
    setPlayerId(0)
  }

  function handleConnect(ev: SubmitEvent) {
    ev.preventDefault();
    setGameState(WebSocket.CONNECTING);
    const ws = new WebSocket("ws://192.168.1.2:8080/join");
    setStatusMsg("Connecting...");

    ws.onopen = () => {
      setGameState(ws.readyState);
      const payload: ServerPayload = {
        name: playerName,
      };
      ws.send(JSON.stringify(payload));

      setStatusMsg("Waiting for opponent.");
      setWsConn(ws);
    };

    ws.onerror = () => {
      clearGameState(ws);
      setStatusMsg("Connection to the server failed.");
    };

    ws.onclose = () => {
      clearGameState(ws);
      setStatusMsg("The connection was closed.");
    };

    ws.onmessage = ({ data }: { data: string }) => {
      const d: string = data;
      if (d.length > 0) {
        const {
          board,
          bonusCard,
          message,
          dices,
          playerPower,
          scoreboards,
          id
        }: GameState = JSON.parse(d);
        board !== null && setGameBoard(board);
        bonusCard !== null && setBonusCard(bonusCard);
        message !== "" && setStatusMsg(message);
        dices !== null && setDices(dices);
        setPower(playerPower);
        setScoreboards(scoreboards);
        id !== 0 ? setPlayerId(id) : ""
      }
    };
  }

  function handleDiceSelect(ev: MouseEvent<HTMLElement>) {
    const payload: ServerPayload = {
      selectedDie: Number(ev.currentTarget.id),
    };

    SendServerData(payload);
  }

  function handleBoardClick(ev: MouseEvent<HTMLElement>) {
    const payload: ServerPayload = {
      boardPosition: {
        row: Number(ev.currentTarget.parentElement?.parentElement?.id),
        cell: Number(ev.currentTarget.id),
      },
    };
    SendServerData(payload);
  }

  function handlePowerChoice(use: boolean) {
    const pwr: PlayerPower = { ...power, use: use, confirmed: true };
    setPower(pwr);
    SendServerData({ playerPower: pwr });
  }

  return (
    <>
      {gameState !== WebSocket.OPEN ? (
        <Connect
          connectHandler={handleConnect}
          playerName={playerName}
          setPlayerName={setPlayerName}
        />
      ) : (
        <div className="lg:text-3xl">
          {statusMsg !== "" ? <StatusPane message={statusMsg} /> : ""}
          <div className="grid grid-cols-5 grid-rows-1">
            {gameState === WebSocket.OPEN && gameBoard !== undefined ? (
              <>
                <DiceSection
                  dices={dices}
                  handleDiceSelect={handleDiceSelect}
                  playerId={playerId}
                />
                {power.type !== 0 && !power.confirmed && (
                  <PowerPrompt
                    handlePowerChoice={handlePowerChoice}
                    power={power}
                  />
                )}
                <Board board={gameBoard} handleOnClick={handleBoardClick} />
                <BonusBoard bonusCard={bonusCard} />
              </>
            ) : (
              ""
            )}
          </div>
        </div>
      )}
      {scoreboards !== null && <ScoreSection scoreboards={scoreboards} />}
    </>
  );
}

function Connect({
  connectHandler,
  playerName,
  setPlayerName,
}: {
  connectHandler: any;
  playerName: string;
  setPlayerName: any;
}) {
  const placeholderText = "Enter name here";

  return (
    <div className="text-center">
      <form
        onSubmitCapture={connectHandler}
        className="flex flex-col max-w-fit p-2 m-auto"
      >
        <h2 className="text-2xl">Enter your name</h2>
        <div className="text-md">
          <input
            placeholder={placeholderText}
            minLength={3}
            value={playerName}
            onChange={(e) => setPlayerName(e.target.value)}
            required
            className="rounded text-gray-900 indent-1 p-0.5"
          />
          <button className="m-2 p-1 px-3 rounded-lg bg-gray-700 hover:bg-gray-600">
            Connect
          </button>
        </div>
      </form>
      <RulesSection />
    </div>
  );
}

const StatusPane = memo(function StatusPane({ message }: { message: string }) {
  return (
    <div className="bg-blue-900 text-white text-center py-1 lg:py-2">
      {message}
    </div>
  );
});

export default Game;
