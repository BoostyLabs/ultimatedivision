import diamond from '../img/MarketPlacePage/marketPlaceCardsGroup/diamond2.png';
import gold from '../img/MarketPlacePage/marketPlaceCardsGroup/gold2.png';
import silver from '../img/MarketPlacePage/marketPlaceCardsGroup/silver2.png';
import wood from '../img/MarketPlacePage/marketPlaceCardsGroup/wood2.png';

export const getBackgroundTypePlayer = () => {

    /*
    * bakgroundtype picture that depend on quality
    */

    const listOfQualities = [
        diamond, gold, silver, wood
    ];
    let background = listOfQualities[Math.floor(Math.random()
        * listOfQualities.length)];
    return background;
};
