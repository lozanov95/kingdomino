import badgeEmpty from "../assets/empty.svg"
import badgeCastle from "../assets/castle.svg"
import badgeChecked from "../assets/checked.svg"
import badgeFilled from "../assets/filled.svg"
import badgeDot from "../assets/dot.svg"
import badgeDoubleDot from "../assets/doubledot.svg"
import badgeLine from "../assets/line.svg"
import badgeDoubleLine from "../assets/doubleline.svg"
import badgeQuestion from "../assets/question.svg"


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