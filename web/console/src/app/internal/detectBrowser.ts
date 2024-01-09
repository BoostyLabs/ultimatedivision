const NO_INDEX_FOR_BROWSER = -1;
export const detectBrowser = () => {
    if (navigator.userAgent.indexOf('Firefox') !== NO_INDEX_FOR_BROWSER) {
        return 'Mozilla Firefox';
    } else if (navigator.userAgent.indexOf('Chrome') !== NO_INDEX_FOR_BROWSER) {
        return 'Google Chrome';
    } else if (navigator.userAgent.indexOf('Edge') !== NO_INDEX_FOR_BROWSER) {
        return 'Microsoft Edge';
    } else if (navigator.userAgent.indexOf('Opera') !== NO_INDEX_FOR_BROWSER) {
        return 'Opera';
    }

    return '';
};
export const openWalletShopPage = (walletPageLink: string) => {
    if (walletPageLink) {
        window.open(
            walletPageLink,
            '_blank'
        );
    }
};
export enum CasperWalletShopPage {
    'Mozilla Firefox' = 'https://addons.mozilla.org/en-US/firefox/addon/casper-wallet/',
    'Opera' = '',
    'Microsoft Edge' = 'https://microsoftedge.microsoft.com/addons/detail/casper-wallet/dfmbcapkkeejcpmfhpnglndfkgmalhik',
    'Google Chrome' = 'https://chrome.google.com/webstore/detail/casper-wallet/abkahkcbhngaebpcgfmhkoioedceoigp',
    '' = ''
}
export enum ConcordiumWalletShopPage {
    'Mozilla Firefox' = '',
    'Opera' = 'https://chromewebstore.google.com/detail/concordium-wallet/mnnkpffndmickbiakofclnpoiajlegmg?hl=en-US',
    'Google Chrome' = 'https://chromewebstore.google.com/detail/concordium-wallet/mnnkpffndmickbiakofclnpoiajlegmg?hl=en-US',
    'Microsoft Edge' = '',
    '' = ''
}
