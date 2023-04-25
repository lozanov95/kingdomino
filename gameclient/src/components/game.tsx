import { useState, MouseEvent, memo } from "react";
import BonusBoard from "./bonusboard";
import {
  Bonus,
  Dice,
  GameState,
  ServerPayload,
  PlayerPower,
  Scoreboard,
  DiceResult,
  BoardPosition,
  GameTurn,
} from "../helpers/types";
import { Board } from "./board";
import { DiceSection } from "./dice";
import { PowerPrompt } from "./powerprompt";
import { RulesSection } from "./rules";
import { ScoreSection } from "./scoreboard";
import { ModalPrompt } from "./modal";
import { isReadyToSubmit, SendServerData } from "../helpers/gamestate";

function Game() {
  const DOMAIN = "192.168.1.2";

  const [gameState, setGameState] = useState<number>(WebSocket.CLOSED);
  const [wsConn, setWsConn] = useState<WebSocket | null>(null);
  const [statusMsg, setStatusMsg] = useState("");
  const [gameTurn, setGameTurn] = useState<GameTurn>(GameTurn.Disconnected);
  const [playerName, setPlayerName] = useState("");
  const [gameBoard, setGameBoard] = useState<Dice[][] | null>(null);
  const [bonusCard, setBonusCard] = useState<Bonus[] | null>(null);
  const [scoreboards, setScoreboards] = useState<Scoreboard[] | null>(null);
  const [dices, setDices] = useState<DiceResult[] | null>(null);
  const [power, setPower] = useState<PlayerPower>({
    type: 0,
    description: "",
    use: false,
    confirmed: false,
  });
  const [playerId, setPlayerId] = useState<number>(0);

  function clearGameState(ws: WebSocket) {
    setGameState(ws.readyState);
    setBonusCard(null);
    setGameBoard(null);
    setDices(null);
    setPlayerId(0);
  }

  function handleConnect(ev: SubmitEvent) {
    ev.preventDefault();
    setGameState(WebSocket.CONNECTING);
    const ws = new WebSocket(`ws://${DOMAIN}:8080/join`);
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
      if (data.length === 0) {
        return;
      }

      const {
        board,
        bonusCard,
        message,
        dices,
        playerPower,
        scoreboards,
        id,
        gameTurn,
      }: GameState = JSON.parse(data);
      board !== null && setGameBoard(board);
      bonusCard !== null && setBonusCard(bonusCard);
      message !== "" && setStatusMsg(message);
      dices !== null && setDices(dices);
      setPower(playerPower);
      setScoreboards(scoreboards);
      id !== 0 ? setPlayerId(id) : "";
      gameTurn !== 0 && setGameTurn(gameTurn);
    };
  }

  function handleSendServerData(payload: ServerPayload) {
    SendServerData(wsConn, payload);
  }

  function handlePowerChoice(use: boolean) {
    const pwr: PlayerPower = { ...power, use: use, confirmed: true };
    setPower(pwr);
    SendServerData(wsConn, { playerPower: pwr });
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
        <GameSection
          display={gameState === WebSocket.OPEN && gameBoard !== undefined}
          dices={dices}
          handleSendServerData={handleSendServerData}
          playerId={playerId}
          gameTurn={gameTurn}
          handlePowerChoice={handlePowerChoice}
          statusMsg={statusMsg}
          bonusCard={bonusCard ?? []}
          gameBoard={gameBoard ?? []}
          power={power}
        />
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

function GameSection({
  display,
  dices,
  playerId,
  gameTurn,
  statusMsg,
  power,
  gameBoard,
  bonusCard,
  handleSendServerData,
  handlePowerChoice,
}: {
  display: boolean;
  dices: DiceResult[] | null;
  playerId: number;
  gameTurn: GameTurn;
  statusMsg: string;
  power: PlayerPower;
  gameBoard: Dice[][];
  bonusCard: Bonus[];
  handleSendServerData: (payload: ServerPayload) => void;
  handlePowerChoice: (use: boolean) => void;
}) {
  const [selectedDie, setSelectedDie] = useState<number>(-1);
  const [boardPosition, setBoardPosition] = useState<BoardPosition>({
    row: -1,
    cell: -1,
  });

  function handleDiceSelect(ev: MouseEvent<HTMLElement>) {
    const id = Number(ev.currentTarget.id);
    const die = dices?.at(id);
    if (die?.isPicked && !die?.isPlaced && die?.playerId === playerId) {
      setSelectedDie(id);
      return;
    }

    handleSendServerData({
      selectedDie: id,
    });
  }

  function handleBoardClick(ev: MouseEvent<HTMLElement>) {
    if (selectedDie === -1 && gameTurn !== GameTurn.HandlePlayerPower) {
      return;
    }

    const row = Number(ev.currentTarget.parentElement?.parentElement?.id);
    const cell = Number(ev.currentTarget.id);
    switch (gameTurn) {
      case GameTurn.PlaceDice:
        setBoardPosition({ row, cell });
        break;
      case GameTurn.HandlePlayerPower:
        handleSendServerData({ boardPosition: { cell, row } });
        break;
      default:
        break;
    }
  }

  function handlePlaceDie(place: boolean) {
    if (!place) {
      setSelectedDie(-1);
      setBoardPosition({ cell: -1, row: -1 });
      return;
    }

    handleSendServerData({
      boardPosition,
      selectedDie,
    });

    setSelectedDie(-1);
    setBoardPosition({ cell: -1, row: -1 });
  }

  return (
    <div className="lg:text-3xl">
      {statusMsg !== "" && <StatusPane message={statusMsg} />}
      <div className="grid grid-cols-5 grid-rows-1">
        {display && (
          <>
            <DiceSection
              dices={dices}
              handleDiceSelect={handleDiceSelect}
              playerId={playerId}
              selectedDie={selectedDie}
            />
            {power.type !== 0 && !power.confirmed && (
              <PowerPrompt
                handlePowerChoice={handlePowerChoice}
                power={power}
              />
            )}
            {isReadyToSubmit(boardPosition, selectedDie) && (
              <ModalPrompt
                prompt={"Do you want to place the die?"}
                onClick={handlePlaceDie}
              />
            )}
            <Board
              board={gameBoard}
              handleOnClick={handleBoardClick}
              boardPosition={boardPosition}
            />
            <BonusBoard bonusCard={bonusCard} />
          </>
        )}
      </div>
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
