import { memo, MouseEventHandler } from "react";
import { Dice, Badge, getBadgeIcon } from "./common";
import { Nobles, Cell } from "./common";

export const Board = memo(function Board({
  board,
  handleOnClick,
}: {
  board: Dice[][] | null;
  handleOnClick: MouseEventHandler;
}) {
  return (
    <div className="flex flex-col max-w-fit col-start-2 col-end-4 lg:col-end-5 m-auto bg-neutral-500">
      {board?.map((el, idx) => {
        return (
          <Row
            key={idx}
            id={idx.toString()}
            elements={el}
            onClick={handleOnClick}
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
}: {
  id?: string;
  elements: Dice[] | null;
  onClick?: MouseEventHandler;
  className?: string;
}) {
  return (
    <div
      className={["flex flex-row max-w-fit m-auto", className].join(" ")}
      id={id}
    >
      {elements?.map(({ name, nobles }, idx) => {
        return (
          <BoardCell
            id={idx.toString()}
            key={idx}
            nobles={nobles}
            name={name}
            onClick={onClick}
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
}: {
  id: string;
  name: Badge;
  nobles: number;
  onClick?: MouseEventHandler;
  className?: string;
  nobleColor?: string;
}) {
  return (
    <div className={["flex max-w-fit p-1", className].join(" ")}>
      <Nobles amount={nobles} color={nobleColor} />
      <Cell id={id} imgSrc={getBadgeIcon(name)} onClick={onClick} />
    </div>
  );
}

export default Board;
