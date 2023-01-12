import { Badge } from "../components/common"

export function getBoard() {
    const board = [
        [0, 2, 3, 4, 5, 6, 7],
        [0, 2, 3, 4, 5, 6, 7],
        [0, 0, 0, 1, 0, 0, 0],
        [0, 2, 3, 4, 5, 6, 7],
        [0, 2, 3, 4, 5, 6, 7],
    ]

    return board
}

export function getDices() {
    const dices = [2, 3, 8, 4]

    return dices
}

export function getBonus() {
    const bonus = [
        {
            badge: Badge.DOT, requiredChecks: 5, currentChecks: 4, eligible: true
        },
        {
            badge: Badge.LINE, requiredChecks: 5, currentChecks: 1, eligible: true
        },
        {
            badge: Badge.DOUBLEDOT, requiredChecks: 4, currentChecks: 0, eligible: true
        },
        {
            badge: Badge.DOUBLELINE, requiredChecks: 4, currentChecks: 2, eligible: false
        },
        {
            badge: Badge.CHECKED, requiredChecks: 3, currentChecks: 3, eligible: false
        },
        {
            badge: Badge.FILLED, requiredChecks: 3, currentChecks: 1, eligible: true
        },
    ]

    return bonus
}