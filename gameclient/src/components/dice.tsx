import { memo, MouseEventHandler, useState } from "react";
import { DiceResult } from "./common";
import { BoardCell } from "./board";

export const DiceSection = memo(function DiceSection({
  dices,
  handleDiceSelect,
  playerId,
}: {
  dices: DiceResult[] | null;
  handleDiceSelect: MouseEventHandler;
  playerId: number;
}) {

  return (
    <div className="text-center mx-auto">
      <div className="text-lg text-center">
        {dices && <h2 className="font-bold text-2xl">Dice</h2>}
        {dices?.map((diceResult, idx) => {
          return (
            <DiceSelectCell
              key={idx}
              diceResult={diceResult}
              id={idx.toString()}
              onClick={shouldBeClickable(diceResult, playerId) ? handleDiceSelect : () => { }}
              playerId={playerId}
              isSelected={true}
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
}: {
  diceResult: DiceResult;
  id: string;
  onClick: MouseEventHandler;
  playerId: number;
  isSelected: boolean;
}) {
  return (
    <BoardCell
      className={
        calculateClass(diceResult.isPlaced, diceResult.isPicked, diceResult.playerId === playerId)
      }
      onClick={onClick}
      name={diceResult.dice.name}
      nobles={diceResult.dice.nobles}
      id={id}
      nobleColor={GetNobleColor(
        playerId,
        diceResult.playerId,
        diceResult.isPicked
      )}
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

function calculateClass(isPlaced: boolean, isPicked: boolean, belongToCurrentPlayer: boolean) {
  const baseClass = "m-auto"

  if (!isPicked || (belongToCurrentPlayer && !isPlaced)) {
    return `${baseClass} hover:scale-110 hover:bg-gray-700 duration-100`
  }

  if (belongToCurrentPlayer && isPlaced) {
    return `${baseClass} grayscale`
  }
}

function shouldBeClickable(diceResult: DiceResult, playerId: number) {
  return !diceResult.isPicked || playerId === diceResult.playerId
}