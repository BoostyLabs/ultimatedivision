// Copyright (C) 2021 Creditor Corp. Group.
// See LICENSE for copying information.

export const PaginatorBlockPages: React.FC<{
    blockPages: number[],
    onPageChange: (type: string, pageNumber?: number) => void,
    currentPage: number,
}> = ({ blockPages, onPageChange, currentPage }) => {
    const getPageClassName = (page: number) => {
        let pageClassName = 'ultimatedivision-paginator__pages__item';

        if (currentPage === page) {
            pageClassName = 'ultimatedivision-paginator__pages__item-active';
        };

        return pageClassName;
    };

    return <ul className="ultimatedivision-paginator__pages">
        {blockPages.map((page, index) =>
            <li
                className={getPageClassName(page)}
                key={index}
                onClick={() => onPageChange('change page', page)}
            >
                {page}
            </li>,
        )}
    </ul>;
}
