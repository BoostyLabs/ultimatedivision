export class DropDownStyles {
    constructor(
        public state: boolean = false,
    ) { }
    get style(): string {
        let style: string ='';
        return style;
    }
}
export class TriangleStyle extends DropDownStyles {
    constructor(
        public state: boolean = false
    ) {
        super(state)
    }
    get style() {
        let style = this.state ? 'rotate(-90deg)' : 'rotate(0deg)';
        return style;
    }
}
export class ListStyle extends DropDownStyles {
    constructor(
        public state: boolean = false
    ) {
        super(state)
    }
    get style() {
        let style = this.state ? '0' : '90px';
        return style;
    }
}