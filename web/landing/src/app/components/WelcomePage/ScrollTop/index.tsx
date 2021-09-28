// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

import './index.scss';

export const ScrollTop: React.FC = () => {

    const handleScroolToTop = () => {
        const scrollBtn = document.querySelector('#scroll-to-top');
        
        document.onscroll = function() {
            if (document.body.scrollTop > 1000 && document.body.scrollTop < (document.body.scrollHeight - 1500)) {
                scrollBtn?.setAttribute('style', 'dispaly:block;');
            } else {
                scrollBtn?.setAttribute('style', 'dispaly:none;');
            }
        };

        window.scrollTo( {
            top: 0,
            behavior: 'smooth'
        });
    };
  
    return (
        <div className="scroll-to-top" id="scroll-to-top" onClick={() => handleScroolToTop()}
            data-aos="fade-bottom"
            data-aos-duration="600"
            data-aos-easing="ease-in-out-cubic"
        >
            <svg width="66" height="66" viewBox="0 0 66 66" fill="none" xmlns="http://www.w3.org/2000/svg">
                <rect x="0.5" y="0.5" width="65" height="65" rx="7.5" fill="#022261" stroke="#3B50BD"/>
                <path d="M32.2223 25.9622C32.6225 25.467 33.3775 25.467 33.7777 25.9622L43.808 38.3714C44.3365 39.0253 43.8711 40 43.0302 40H22.9698C22.1289 40 21.6635 39.0253 22.192 38.3714L32.2223 25.9622Z" fill="#0CE255"/>
            </svg>
        </div>
    );
};
