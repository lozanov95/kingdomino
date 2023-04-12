import { memo, MouseEventHandler } from "react";
import { DiceResult, Domino } from "./common";
import { BoardCell } from "./board";

export const DiceSection = memo(function DiceSection({
  dices,
  handleDiceSelect,
}: {
  dices: DiceResult[] | null;
  selectedDice: Domino[] | null;
  handleDiceSelect: MouseEventHandler;
}) {
  return (
    <div className="text-center mx-auto">
      <div className="text-lg text-center">
        {dices && <h2 className="font-bold text-2xl">Available dice</h2>}
        {dices?.map((diceResult, idx) => {
          return <DiceSelectCell diceResult={diceResult} id={idx.toString()} onClick={handleDiceSelect} />
        })}
      </div>
    </div>
  );
});

function DiceSelectCell({ diceResult, id, onClick }: { diceResult: DiceResult, id: string, onClick: MouseEventHandler }) {

  return (
    <BoardCell onClick={onClick} name={diceResult.dice.name} nobles={diceResult.dice.nobles} id={id} nobleColor={GetNobleColor(diceResult.isSelected, diceResult.playerId)} />
  )
}

function GetNobleColor(isSelected: boolean, playerId: number) {
  if (!isSelected) {
    return "bg-gray-600"
  }

  if (playerId === 0) {
    return "bg-green-700"
  }

  return "bg-red-700"
}