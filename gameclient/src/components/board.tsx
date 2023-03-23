import { memo, MouseEventHandler, MouseEvent } from "react";
import { Domino, Badge, getBadgeIcon } from "./common";
import { Nobles, Cell } from "./common";

export const Board = memo(function Board({
  board,
  handleOnClick,
}: {
  board: Domino[][] | null;
  handleOnClick: MouseEventHandler;
}) {
  return (
    <div className="flex flex-col max-w-fit col-start-2 col-end-4">
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
  elements: Domino[] | null;
  onClick?: MouseEventHandler;
  className?: string;
}) {
  return (
    <div className={["flex flex-row", className].join(" ")} id={id}>
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
}: {
  id: string;
  name: Badge;
  nobles: number;
  onClick?: MouseEventHandler;
}) {
  return (
    <div className="flex border-solid first-of-type:border-t-2 border-x-2 border-b-2 max-w-fit p-1 m-0">
      <Nobles amount={nobles} />
      <Cell id={id} imgSrc={getBadgeIcon(name)} onClick={onClick} />
    </div>
  );
}

export default Board;
