import badgeEmpty from "../assets/empty.svg"
import badgeCastle from "../assets/castle.svg"
import badgeChecked from "../assets/checked.svg"
import badgeFilled from "../assets/filled.svg"
import badgeDot from "../assets/dot.svg"
import badgeDoubleDot from "../assets/doubledot.svg"
import badgeLine from "../assets/line.svg"
import badgeDoubleLine from "../assets/doubleline.svg"
import badgeQuestion from "../assets/question.svg"
import { MouseEventHandler } from "react"


export type Domino = {
    name: Badge,
    nobles: number
}

export type Bonus = {
    name: number
    requiredChecks: number,
    currentChecks: number,
    eligible: boolean
}

export type GameState = {
    bonusCard: Bonus[],
    message: string,
    board: Domino[][],
    dices: Domino[],
    selectedDice: Domino[],
}

export enum Badge {
    EMPTY = 0,
    CASTLE,
    DOT,
    LINE,
    DOUBLEDOT,
    DOUBLELINE,
    FILLED,
    CHECKED,
    QUESTIONMARK,
}

export function getBadgeIcon(id: number) {
    switch (id) {
        case Badge.EMPTY:
            return badgeEmpty
        case Badge.CASTLE:
            return badgeCastle
        case Badge.DOT:
            return badgeDot
        case Badge.LINE:
            return badgeLine
        case Badge.DOUBLEDOT:
            return badgeDoubleDot
        case Badge.DOUBLELINE:
            return badgeDoubleLine
        case Badge.FILLED:
            return badgeFilled
        case Badge.CHECKED:
            return badgeChecked
        case Badge.QUESTIONMARK:
            return badgeQuestion
        default:
            return ""
    }
}

export function Cell({ id, imgSrc, onClick }: { id: string, imgSrc: string, onClick?: MouseEventHandler }) {
    return (
        <div className="cell" id={id} onClick={onClick}>
            <img src={imgSrc} alt="badge icon" />
        </div>
    )
}

export function Nobles({ amount }: { amount: number }) {
    function renderNobles() {
        switch (amount) {
            case 0:
                return
            case 1:
                return <Noble />
            case 2:
                return (
                    <>
                        <Noble /><Noble />
                    </>
                )
            default:
                return null
        }
    }
    return (
        <div className="nobles">
            {renderNobles()}
        </div>
    )
}

export function Noble() {
    return (
        <div>X</div>
    )
}
