export function diagramColor(value, callplace) {

    if (callplace === 'priceArea') {
        switch (true) {
            case (+value >= 80):
                return '#1898D7';
            case (+value >= 70):
                return '#3CCF5D';
            case (+value >= 50):
                return '#E86C27';
            default:
                return '#FF4200';
        }
    }

    switch (true) {
        case (+value >= 90):
            return '#3CCF5D';
        case (+value >= 50):
            return '#E8EC16';
        default:
            return '#FF4200';
    }
}