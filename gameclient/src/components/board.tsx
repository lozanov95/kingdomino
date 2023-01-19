import { memo, MouseEventHandler } from "react"
import { Domino, Badge, getBadgeIcon } from "./common"
import { Nobles, Cell } from "./common"

export const Board = memo(
    function Board({ board }: { board?: Domino[][] }) {

        return (
            <div className="board center" >
                {board?.map((el, idx) => {
                    return <Row key={idx} elements={el} />
                })}
            </div>
        )
    }
)

function Row(props: { elements: Domino[] }) {
    return (
        <div className="row">
            {props.elements.map(({ name, nobles }, idx) => {
                return <BoardCell id={idx.toString()} key={idx} nobles={nobles} name={name} />
            })}
        </div>
    )
}


export function BoardCell({ id, name, nobles, onClick }: { id: string, name: Badge, nobles: number, onClick?: MouseEventHandler }) {
    return (
        <div className="boardCell">
            <Nobles amount={nobles} />
            <Cell id={id} imgSrc={getBadgeIcon(name)} onClick={onClick} />
        </div >
    )
}

export default Board
