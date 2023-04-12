import { memo, MouseEventHandler } from "react";
import { DiceResult, Domino } from "./common";
import { BoardCell } from "./board";

export const DiceSection = memo(function DiceSection({
  dices,
  handleDiceSelect, playerId
}: {
  dices: DiceResult[] | null;
  handleDiceSelect: MouseEventHandler;
  playerId: number;
}) {
  return (
    <div className="text-center mx-auto">
      <div className="text-lg text-center">
        {dices && <h2 className="font-bold text-2xl">Available dice</h2>}
        {dices?.map((diceResult, idx) => {
          return <DiceSelectCell diceResult={diceResult} id={idx.toString()} onClick={handleDiceSelect} playerId={playerId} />
        })}
      </div>
    </div>
  );
});

function DiceSelectCell({ diceResult, id, onClick, playerId }: { diceResult: DiceResult, id: string, onClick: MouseEventHandler, playerId: number }) {

  return (
    <BoardCell onClick={onClick} name={diceResult.dice.name} nobles={diceResult.dice.nobles} id={id} nobleColor={GetNobleColor(playerId, diceResult.playerId, diceResult.isSelected)} />
  )
}

function GetNobleColor(playerId: number, dicePlayerId: number, isSelected: boolean,) {
  if (!isSelected) {
    return "bg-gray-600"
  }

  if (playerId === dicePlayerId) {
    return "bg-green-700"
  }

  return "bg-red-700"
}