import { memo, MouseEventHandler } from "react";
import { Domino } from "./common";
import { Row } from "./board";

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
    <div className="text-center mx-auto">
      <div className="text-lg text-center">
        {dices ? <h2 className="font-bold text-2xl">Available dice</h2> : ""}
        <Row elements={dices} onClick={handleDiceSelect} className="flex-col" />
      </div>
      <div className="text-lg">
        {selectedDice?.length ?? 0 > 0 ? (
          <h2 className="font-bold">Selected dice</h2>
        ) : (
          ""
        )}
        <Row elements={selectedDice} onClick={handleDiceSelect} />
      </div>
    </div>
  );
});
