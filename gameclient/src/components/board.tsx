import { memo, MouseEventHandler } from "react";
import { getBadgeIcon } from "./common";
import { Dice, Badge, BoardPosition } from "../helpers/types";
import { Nobles, Cell } from "./common";

export const Board = memo(function Board({
  board,
  handleOnClick,
  boardPosition,
}: {
  board: Dice[][];
  handleOnClick: MouseEventHandler;
  boardPosition: BoardPosition;
}) {
  return (
    <div className="flex flex-col max-w-fit col-start-2 col-end-4 lg:col-end-5 m-auto bg-neutral-500 mt-9">
      {board.map((el, idx) => {
        return (
          <Row
            key={idx}
            id={idx}
            elements={el}
            onClick={handleOnClick}
            boardPosition={boardPosition}
          />
        );
      })}
    </div>
  );
});

export function Row({
  elements,
  id,
  onClick,
  className,
  boardPosition,
}: {
  id?: number;
  elements: Dice[] | null;
  onClick?: MouseEventHandler;
  className?: string;
  boardPosition: BoardPosition;
}) {
  return (
    <div
      className={["flex flex-row max-w-fit m-auto", className].join(" ")}
      id={id?.toString()}
    >
      {elements?.map(({ name, nobles }, idx) => {
        return (
          <BoardCell
            id={idx.toString()}
            key={idx}
            nobles={nobles}
            name={name}
            onClick={onClick}
            selected={id === boardPosition.row && idx === boardPosition.cell}
          />
        );
      })}
    </div>
  );
}

export function BoardCell({
  id,
  name,
  nobles,
  onClick,
  className,
  nobleColor,
  selected,
}: {
  id: string;
  name: Badge;
  nobles: number;
  onClick?: MouseEventHandler;
  className?: string;
  nobleColor?: string;
  selected: boolean;
}) {
  const classList = ["flex max-w-fit p-1", className];

  if (selected) {
    classList.push("bg-green-300");
  }

  return (
    <div className={classList.join(" ")}>
      <Nobles amount={nobles} color={nobleColor} />
      <Cell id={id} imgSrc={getBadgeIcon(name)} onClick={onClick} />
    </div>
  );
}

export default Board;
