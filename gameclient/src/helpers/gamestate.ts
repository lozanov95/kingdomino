import { BoardPosition, ServerPayload } from "../helpers/types";

export function isReadyToSubmit(
  boardPosition: BoardPosition,
  selectedDie: number
) {
  return (
    boardPosition.cell !== -1 && boardPosition.row !== -1 && selectedDie !== -1
  );
}

export function SendServerData(
  wsConn: WebSocket | null,
  payload: ServerPayload
) {
  wsConn?.send(JSON.stringify(payload));
}
