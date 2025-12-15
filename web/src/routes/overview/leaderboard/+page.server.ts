import type {PageServerLoad} from "./$types";
import type {GetWinnersResponse} from "$lib/gen/zelvy/user/get_winners_response";
import {getFormattedDateParams, setDefaultDateParams,} from "$lib/utils/setDefaultDateParams";
import {getWinners} from "$lib/server/grpc/user";


export const load: PageServerLoad = async ({url}): Promise<{
    winRes: GetWinnersResponse,
} | null> => {

    const {formattedSD, formattedED, isInvalid} = getFormattedDateParams(url);
    if (isInvalid) {
        setDefaultDateParams(url, true);
    }

    const winnersResponse = await getWinners(formattedSD, formattedED);

    return {
        winRes: winnersResponse
    };
};
