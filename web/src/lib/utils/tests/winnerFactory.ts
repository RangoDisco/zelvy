import type {GetWinnersResponse} from "$lib/gen/zelvy/user/get_winners_response";

export const mockGetWinnerResponse = (): GetWinnersResponse => {
    return {
        winners: [
            {
                username: "user1",
                wins: 2,
                picture: "https://example.com"
            },
            {
                username: "user2",
                wins: 1,
                picture: "https://example.com"
            }
        ],
    };
};



