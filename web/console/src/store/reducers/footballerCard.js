/*
Copyright (C) 2021 Creditor Corp. Group.
See LICENSE for copying information.
*/

class CardStats {
    constructor(fields) {
        this.fields = fields;
        this.average = fields
            .map(item => item.value)
            .reduce((prev, current) => prev + current) / fields.length;
    }
    get color() {
        switch (true) {
        case (this.average >= 90):
            return '#3CCF5D';
        case (this.average >= 50):
            return '#E8EC16';
        default:
            return '#FF4200';
        }
    }
}
class CardStatsField {
    constructor(key, value) {
        this.label = key;
        this.value = value;
    }
}
class Card {
    tactics = new CardStats([
        new CardStatsField('positioning', 100),
        new CardStatsField('composure', 95,),
        new CardStatsField('aggression', 98),
        new CardStatsField('vision', 98),
        new CardStatsField('awareness', 99),
        new CardStatsField('crosses', 98),
    ]);
    physique = new CardStats([
        new CardStatsField('acceleration', 26),
        new CardStatsField('running speed', 25),
        new CardStatsField('reaction speed', 45),
        new CardStatsField('agility', 31),
        new CardStatsField('stamina', 40),
        new CardStatsField('strength', 35),
        new CardStatsField('jumping', 28),
        new CardStatsField('balance', 42),
    ]);
    technique = new CardStats([
        new CardStatsField('dribbing', 26),
        new CardStatsField('ball fontrol', 26),
        new CardStatsField('weak foot', 26),
        new CardStatsField('skill moves', 26),
        new CardStatsField('finesse', 26),
        new CardStatsField('curve', 26),
        new CardStatsField('volleys', 26),
        new CardStatsField('short passing', 26),
        new CardStatsField('long passing', 26),
        new CardStatsField('forward pass', 26),
    ]);
    offence = new CardStats([
        new CardStatsField('finishing ability', 42),
        new CardStatsField('shot power', 42),
        new CardStatsField('accuracy', 42),
        new CardStatsField('distance', 42),
        new CardStatsField('penalty', 42),
        new CardStatsField('free Kicks', 42),
        new CardStatsField('corners', 42),
        new CardStatsField('heading accuracy', 42),
    ]);
    defence = new CardStats([
        new CardStatsField('offside trap', 74),
        new CardStatsField('tackles', 74),
        new CardStatsField('ball focus', 74),
        new CardStatsField('interceptions', 74),
        new CardStatsField('vigilance', 74),
    ]);
    goalkeeping = new CardStats([
        new CardStatsField('diving', 84),
        new CardStatsField('handling', 84),
        new CardStatsField('sweeping', 84),
        new CardStatsField('throwing', 84),
    ]);
}

function cardlist(count) {
    let quantity = count;
    const list = [];

    while (quantity > 0) {
        list.push(new Card());
        quantity--;
    }

    return list;
}
/* eslint-disable */
export const cardStatsReducer = (cardState = cardlist(3), action) => {
    return cardState;
};