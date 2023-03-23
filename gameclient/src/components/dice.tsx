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
    <div className="max-w-screen-lg">
      <div className="text-lg max-w-max text-center">
        {dices ? <h2>Available dice</h2> : ""}
        {dices?.map(({ name, nobles }, idx) => {
          return (
            <BoardCell
              id={idx.toString()}
              key={idx}
              name={name}
              nobles={nobles}
              onClick={handleDiceSelect}
            />
          );
        })}
      </div>
      <div>
        {selectedDice?.length ?? 0 > 0 ? <h2>Selected</h2> : ""}
        <Row elements={selectedDice} onClick={handleDiceSelect} />
      </div>
    </div>
  );
});
