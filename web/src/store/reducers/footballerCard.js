import player from '../../img/MarketPlacePage/marketPlaceCardsGroup/player.png';
import price from '../../img/MarketPlacePage/marketPlaceCardsGroup/price.png';
import { getBackgroundTypePlayer } from '../../utils/getBackgroundTypePlayer';

class Card {
    mainInfo = {
        price: 1000000,
        get backgroundType() {
            return getBackgroundTypePlayer();
        },
        facePicture: player,
        pricePicture: price
    };
    overalInfo = {
        'name': 'Albert Ronalculus',
        'nation': 'Portugal',
        'skills': '5',
        'weak foot': '4',
        'intl. Rep': '5',
        'foot': 'Right',
        'height': '187',
        'nation?': '83',
        'revision': 'Rare',
        'def. WR': 'Low',
        'att. WR': 'High',
        'added on': '2020-09-10',
        'origin': 'NA',
        'r.Face': 'Low',
        'r.Type': true,
        'age': '36',
    };
    tactics = {
        'tactics': 98,
        'positioning': 70,
        'composure': 70,
        'aggression': 70,
        'vision': 70,
        'awareness': 70,
        'crosses': 70,
    };
    physique = {
        'physique': 34,
        'acceleration': 70,
        'running speed': 70,
        'reaction speed': 70,
        'agility': 70,
        'stamina': 70,
        'strength': 70,
        'jumping': 70,
        'balance': 70,
    };
    technique = {
        'technique': 26,
        'dribbing': 70,
        'ball Control': 70,
        'weak Foot': 70,
        'skill Moves': 70,
        'finesse': 70,
        'curve': 70,
        'volleys': 70,
        'short passing': 70,
        'long passing': 70,
        'forward pass': 70,
    };
    offence = {
        'offence': 42,
        'finishing ability': 70,
        'shot power': 70,
        'accuracy': 70,
        'distance': 70,
        'penalty': 70,
        'free Kicks': 70,
        'corners': 70,
        'heading accuracy': 70,
    };
    defence = {
        'defence': 74,
        'offside trap': 70,
        'tackles': 70,
        'ball focus': 70,
        'interceptions': 70,
        'vigilance': 70,
    };
    goalkeeping = {
        'goalkeeping': 84,
        'reflexes': 70,
        'diving': 70,
        'handling': 70,
        'sweeping': 70,
        'throwing': 70,
    }
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
export const cardReducer = (cardState = cardlist(24), action) => {
    return cardState;
};
