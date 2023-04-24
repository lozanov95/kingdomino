import { memo, MouseEventHandler } from "react";
import { DiceResult } from "../helpers/types";
import { BoardCell } from "./board";

export const DiceSection = memo(function DiceSection({
  dices,
  handleDiceSelect,
  playerId,
  selectedDie,
}: {
  dices: DiceResult[] | null;
  handleDiceSelect: MouseEventHandler;
  playerId: number;
  selectedDie: number;
}) {
  if (dices === null) {
    return <></>;
  }

  return (
    <div className="text-center mx-auto">
      <div className="text-lg text-center">
        <h2 className="font-bold text-2xl">Dice</h2>
        {dices.map((diceResult, idx) => {
          return (
            <DiceSelectCell
              key={idx}
              diceResult={diceResult}
              id={idx.toString()}
              onClick={
                shouldBeClickable(diceResult, playerId)
                  ? handleDiceSelect
                  : () => {}
              }
              playerId={playerId}
              isSelected={selectedDie === idx}
            />
          );
        })}
      </div>
    </div>
  );
});

function DiceSelectCell({
  diceResult,
  id,
  onClick,
  playerId,
  isSelected,
}: {
  diceResult: DiceResult;
  id: string;
  onClick: MouseEventHandler;
  playerId: number;
  isSelected: boolean;
}) {
  return (
    <BoardCell
      className={calculateClass(
        diceResult.isPlaced,
        diceResult.isPicked,
        diceResult.playerId === playerId
      )}
      onClick={onClick}
      name={diceResult.dice.name}
      nobles={diceResult.dice.nobles}
      id={id}
      nobleColor={GetNobleColor(
        playerId,
        diceResult.playerId,
        diceResult.isPicked
      )}
      selected={isSelected}
    />
  );
}

function GetNobleColor(
  playerId: number,
  dicePlayerId: number,
  isPicked: boolean
) {
  if (!isPicked) {
    return "bg-gray-600";
  }

  if (playerId === dicePlayerId) {
    return "bg-green-700";
  }

  return "bg-red-700";
}

function calculateClass(
  isPlaced: boolean,
  isPicked: boolean,
  belongToCurrentPlayer: boolean
) {
  const classList = ["m-auto"];

  if (!isPicked || (belongToCurrentPlayer && !isPlaced)) {
    classList.push("hover:scale-110 hover:bg-gray-700 duration-100");
  }

  if (belongToCurrentPlayer && isPlaced) {
    classList.push("grayscale");
  }

  return classList.join(" ");
}

function shouldBeClickable(diceResult: DiceResult, playerId: number) {
  return !diceResult.isPicked || playerId === diceResult.playerId;
}
