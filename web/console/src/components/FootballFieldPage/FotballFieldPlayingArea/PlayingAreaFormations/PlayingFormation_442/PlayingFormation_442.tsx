import React from 'react';
import './PlayingFormation_442.scss';

export const PlayingFormation_442 = () => {
    const cardsPositions = [
        {
            id: 0,
            top: 0,
            left: 0
        },
        {
            id: 1,
            top: 0,
            left: 0
        },
        {
            id: 2,
            top: 0,
            left: 0
        },
        {
            id: 3,
            top: 0,
            left: 0
        },
        {
            id: 4,
            top: 0,
            left: 0
        },
        {
            id: 5,
            top: 0,
            left: 0
        },
        {
            id: 6,
            top: 0,
            left: 0
        },
        {
            id: 7,
            top: 0,
            left: 0
        },
        {
            id: 8,
            top: 0,
            left: 0
        },
        {
            id: 9,
            top: 0,
            left: 0
        },
        {
            id: 10,
            top: 0,
            left: 0
        },
    ]
    return (
        <div className="playing-formation-442">
            {cardsPositions.map(card => (
                <div
                    className="playing-formation-442__card"
                >
                    {card.id}
                </div>
            ))}
        </div>
    )
}
