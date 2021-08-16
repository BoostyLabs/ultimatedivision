// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.


/** class for our getters to get label and value while mapping */
export class CardField {
    /** label and value for mapping */
    constructor(
        public label: string,
        public value: string | number,
    ) { }
}

/* eslint-disable */
/** player stats implementation */
export class CardStats {
    /** main stat with substats */
    constructor(
        public title: string = '',
        public abbreviated: string = '',
        public fields: CardField[] = []
    ) { }
    public average: number = this.fields
        .map(item => +item.value)
        .reduce((prev, current) => prev + current) / this.fields.length;
    /** abbreviated title of card stat name */
    get abbr(): string {
        return this.title.slice(0, 3);
    }
    /** stat giagram color depend on avarage stat value */
    get color(): string {
        const STATISTIC_UPPER_BOUND = 90;
        const STATISTIC_LOWER_BOUND = 50;

        const STATISTIC_UPPER_BOUND_COLOR = '#3CCF5D';
        const STATISTIC_MEDIUM_BOUND_COLOR = '#E8EC16';
        const STATISTIC_LOWER_BOUND_COLOR = '#FF4200';

        switch (true) {
            case this.average >= STATISTIC_UPPER_BOUND:
                return STATISTIC_UPPER_BOUND_COLOR;
            case this.average >= STATISTIC_LOWER_BOUND:
                return STATISTIC_MEDIUM_BOUND_COLOR;
            default:
                return STATISTIC_LOWER_BOUND_COLOR;
        }
    }
}

/** Card base implementation */
export class Card {
    /** Card fields */
    constructor(
        public id: string,
        public playerName: number,
        public quality: number,
        public pictureType: number,
        public height: number,
        public weight: number,
        public skinColor: number,
        public hairStyle: number,
        public hairColor: number,
        public accessories: string,
        public dominantFoot: number,
        public isTatoos: boolean,
        public userId: string,
        public tactics: number,
        public positioning: number,
        public composure: number,
        public aggression: number,
        public vision: number,
        public awareness: number,
        public crosses: number,
        public physique: number,
        public acceleration: number,
        public runningSpeed: number,
        public reactionSpeed: number,
        public agility: number,
        public stamina: number,
        public strength: number,
        public jumping: number,
        public balance: number,
        public technique: number,
        public dribbling: number,
        public ballControl: number,
        public weakFoot: number,
        public skillMoves: number,
        public finesse: number,
        public curve: number,
        public volleys: number,
        public shortPassing: number,
        public longPassing: number,
        public forwardPass: number,
        public offense: number,
        public finishingAbility: number,
        public shotPower: number,
        public accuracy: number,
        public distance: number,
        public penalty: number,
        public freeKicks: number,
        public corners: number,
        public headingAccuracy: number,
        public defence: number,
        public offsideTrap: number,
        public sliding: number,
        public tackles: number,
        public ballFocus: number,
        public interceptions: number,
        public vigilance: number,
        public goalkeeping: number,
        public reflexes: number,
        public diving: number,
        public handling: number,
        public sweeping: number,
        public throwing: number,
    ) { }
    /** Using in footballerCard in info block */
    get mainInfo() {
        return [
            new CardField('name', this.playerName),
            // To do: at this momenty nation does not exist
            new CardField('nation', 'this.nation'),
            new CardField('skills', '5'),
            new CardField('weak foot', this.weakFoot),
            new CardField('intl. rep', '5'),
            new CardField('foot', this.dominantFoot),
            new CardField('height', this.height),
            new CardField('nation', this.weight),
            // To do: at this momenty revision does not exist or it is designer mistake or it is quality
            new CardField('revision', 'rare'),
            // To do: create method to convert attack and defence values into this scale
            new CardField('def. wr', 'low'),
            new CardField('arr. wr', 'high'),
            // next fields does not exist in card at this moment
            new CardField('added on', '2020-09-10'),
            new CardField('origin', 'na'),
            new CardField('r. Face', 'low'),
            new CardField('b. type', ''),
            new CardField('age', '36 years old'),
        ];
    }

    /** Using in diagramm area in footballerCard */
    get diagramArea() {
        // To do: need to get real min and max values to convert into diagram value
        // To do: this fields does not exist
        return [
            new CardField('physical', 688),
            new CardField('mental', 688),
            new CardField('skill', 688),
            new CardField('cham. style', 688),
            new CardField('base stats', 688),
            new CardField('in game stats', 688),
        ];
    }

    /** returns fields for card stats area in footballerCard */
    get statsArea() {
        return [
            new CardStats('tactics', 'tac', [
                new CardStatsField('positioning', this.positioning),
                new CardStatsField('composure', this.composure,),
                new CardStatsField('aggression', this.aggression),
                new CardStatsField('vision', this.vision),
                new CardStatsField('awareness', this.awareness),
                new CardStatsField('crosses', this.crosses),
            ]),
            new CardStats('physique', 'phy', [
                new CardStatsField('acceleration', this.acceleration),
                new CardStatsField('running speed', this.runningSpeed),
                new CardStatsField('reaction speed', this.reactionSpeed),
                new CardStatsField('agility', this.agility),
                new CardStatsField('stamina', this.stamina),
                new CardStatsField('strength', this.strength),
                new CardStatsField('jumping', this.jumping),
                new CardStatsField('balance', this.jumping),
            ]),
            new CardStats('technique', 'tec', [
                new CardStatsField('dribbing', this.dribbling),
                new CardStatsField('ball fontrol', this.ballControl),
                new CardStatsField('weak foot', this.weakFoot),
                new CardStatsField('skill moves', this.skillMoves),
                new CardStatsField('finesse', this.finesse),
                new CardStatsField('curve', this.curve),
                new CardStatsField('volleys', this.volleys),
                new CardStatsField('short passing', this.shortPassing),
                new CardStatsField('long passing', this.longPassing),
                new CardStatsField('forward pass', this.forwardPass),
            ]),
            new CardStats('offence', 'off', [
                new CardStatsField('finishing ability', this.finishingAbility),
                new CardStatsField('shot power', this.shotPower),
                new CardStatsField('accuracy', this.accuracy),
                new CardStatsField('distance', this.distance),
                new CardStatsField('penalty', this.penalty),
                new CardStatsField('free Kicks', this.freeKicks),
                new CardStatsField('corners', this.corners),
                new CardStatsField('heading accuracy', this.headingAccuracy),
            ]),
            new CardStats('defence', 'def', [
                new CardStatsField('offside trap', this.offsideTrap),
                new CardStatsField('tackles', this.tackles),
                new CardStatsField('ball focus', this.ballFocus),
                new CardStatsField('interceptions', this.interceptions),
                new CardStatsField('vigilance', this.vigilance),
            ]),
            new CardStats('goalkeeping', 'gk', [
                new CardStatsField('diving', this.diving),
                new CardStatsField('handling', this.handling),
                new CardStatsField('sweeping', this.sweeping),
                new CardStatsField('throwing', this.throwing),
            ]),
        ];
    }
}
