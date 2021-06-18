import { diagramColor } from '../../utils/fotballerCard'

class Card {
    overalInfo = {
        'Name': 'Albert Ronalculus',
        'Nation': 'Portugal 🇵🇹',
        'Skills': '5',
        'Weak foot': '4',
        'Intl. Rep': '5',
        'Foot': 'Right',
        'Height': '187',
        'Nation?': '83',
        'Revision': 'Rare',
        'Def. WR': 'Low',
        'Att. WR': 'High',
        'Added on': '2020-09-10',
        'Origin': 'NA',
        'R.Face': 'Low',
        'B.Type': true,
        'Age': '36 years old',
    }
    tactics = {
        'tactics': 98,
        'positioning': 70,
        'composure': 70,
        'aggression': 70,
        'vision': 70,
        'awareness': 70,
        'crosses': 70,
        get color() {
            return diagramColor(this.tactics)
        }
    }
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
        get color() {
            return diagramColor(this.physique)
        }
    }
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
        get color() {
            return diagramColor(this.technique)
        }
    }
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
        get color() {
            return diagramColor(this.offence)
        }
    }
    defence = {
        'defence': 74,
        'offside trap': 70,
        'tackles': 70,
        'ball focus': 70,
        'interceptions': 70,
        'vigilance': 70,
        get color() {
            return diagramColor(this.defence)
        }
    }
    goalkeeping = {
        'goalkeeping': 84,
        'diving': 70,
        'handling': 70,
        'sweeping': 70,
        'throwing': 70,
        get color() {
            return diagramColor(this.goalkeeping)
        }
    }
};

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
export const cardReducer = (cardState = cardlist(15), action) => {
    return cardState;
};