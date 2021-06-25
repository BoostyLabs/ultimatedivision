export class CardStats {
    average: number = this.fields
        .map(item => item.value)
        .reduce((prev, current) => prev + current) / this.fields.length;
    constructor(public title: string, public fields: CardStatsField[]) {
    }
    get color(): string {
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

export class CardStatsField {
    constructor(public label: string, public value: number) { }
}

export class CardInfoField {
    constructor(
        public label: string,
        public value: string | number | boolean,
        public icon?: string | undefined
    ) { }
}
export class CardPriceField {
    constructor(
        public label: string,
        public value: number | string
    ) { }
}
export class CardPricePRP {
    constructor(
        public label: string,
        public value: number
    ) { }
}

export class CardPrice {
    constructor(
        public id: CardPriceField,
        public price: CardPriceField,
        public prp: CardPricePRP,
        public updated: CardPriceField,
        public pr: CardPriceField,
    ) { }

    get color() {
        switch (true) {
            case (this.prp.value >= 80):
                return '#1898D7';
            case (this.prp.value >= 70):
                return '#3CCF5D';
            case (this.prp.value >= 50):
                return '#E86C27';
            default:
                return '#FF4200';
        }
    }
}

export class Diagram {
    constructor(
        public id: string,
        public name: string,
        public min: number,
        public max: number,
        public value: number,
    ) { }
}
