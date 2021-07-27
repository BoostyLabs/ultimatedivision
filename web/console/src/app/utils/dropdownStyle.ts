export class DropdownStyle {
    constructor(
        public vilibility: boolean,
        public height: number
    ) { }
    get triangleRotate() {
        return this.vilibility ? 'rotate(0deg)' : 'rotate(-90deg)';
    }
    get listHeight () {
        return this.vilibility ? `${this.height}px` : '0';
    }
}