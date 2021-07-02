export class Styles {
    constructor(
        public state: boolean = false
    ) { }
    get listStyle() {
        return this.state ? '0' : '90px';
    }
    get triangleStyle() {
        return this.state ? 'rotate(-90deg)' : 'rotate(0deg)';
    }
}