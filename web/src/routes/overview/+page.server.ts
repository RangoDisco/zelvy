import type {PageServerLoad} from "./$types";
import type {GetWinnersResponse} from "$lib/gen/zelvy/user/get_winners_response";
import type {GetSummaryHeatmapResponse} from "$lib/gen/zelvy/summary/get_summary_heatmap_response";
import setDefaultDateRangeParams from "$lib/utils/setDefaultDateRangeParams";
import {isDate} from "node:util/types";
import {getSummaryHeatmap} from "$lib/server/grpc/summary";
import {getWinners} from "$lib/server/grpc/user";


export const load: PageServerLoad = async ({url}): Promise<{
    winRes: GetWinnersResponse,
    hmRes: GetSummaryHeatmapResponse
} | null> => {
    const eDParam = url.searchParams.get("end_date");
    const sDParam = url.searchParams.get("start_date");

    if (eDParam === null || sDParam === null) {
        setDefaultDateRangeParams(url);
        return null;
    }

    const endDate = new Date(eDParam);
    const startDate = new Date(sDParam);

    if (startDate > endDate || !isDate(endDate) || !isDate(startDate)) {
        setDefaultDateRangeParams(url);
        return null;
    }

    const formattedED = endDate.toISOString().slice(0, 10);
    const formattedSD = startDate.toISOString().slice(0, 10);

    const winnersResponse = await getWinners(formattedSD, formattedED);
    const heatmapResponse = await getSummaryHeatmap(formattedSD, formattedED);

    return {
        winRes: winnersResponse,
        hmRes: heatmapResponse
    };
};
