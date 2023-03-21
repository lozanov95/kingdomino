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
    <div className="board center">
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
}: {
  id?: string;
  elements: Domino[] | null;
  onClick?: MouseEventHandler;
}) {
  return (
    <div className="row" id={id}>
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
    <div className="boardCell">
      <Nobles amount={nobles} />
      <Cell id={id} imgSrc={getBadgeIcon(name)} onClick={onClick} />
    </div>
  );
}

export default Board;
