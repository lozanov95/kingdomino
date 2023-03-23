import { memo, MouseEventHandler } from "react";
import { Domino } from "./common";
import { BoardCell, Row } from "./board";

export const DiceSection = memo(function DiceSection({
  dices,
  selectedDice,
  handleDiceSelect,
}: {
  dices: Domino[] | null;
  selectedDice: Domino[] | null;
  handleDiceSelect: MouseEventHandler;
}) {
  return (
    <div className="max-w-[20%] text-center">
      <div className="text-lg">
        {dices ? <h2 className="font-bold">Available dice</h2> : ""}
        <Row elements={dices} onClick={handleDiceSelect} />
      </div>
      <div className="text-lg">
        {selectedDice?.length ?? 0 > 0 ? (
          <h2 className="font-bold">Selected</h2>
        ) : (
          ""
        )}
        <Row elements={selectedDice} onClick={handleDiceSelect} />
      </div>
    </div>
  );
});
