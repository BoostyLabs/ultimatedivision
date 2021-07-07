export interface DropDownStyles {
    state: boolean;
    get style(): string;
}
export class TriangleStyle implements DropDownStyles {
    constructor(
        public state: boolean = false
    ) { }
    get style() {
        return this.state ? 'rotate(-90deg)' : 'rotate(0deg)';
    }
}
export class ListStyle implements DropDownStyles {
    constructor(
        public state: boolean = false
    ) { }
    get style() {
        return this.state ? '0' : '90px';
    }
}

export class FootballCardStyle implements DropDownStyles {
    constructor (
        public state: boolean
        ) { }
        get style() {
            return this.state ? 'block' : 'none';
        }
}