import type {PageServerLoad} from "./$types";
import type {GetWinnersResponse} from "$lib/gen/zelvy/user/get_winners_response";
import type {GetSummaryHeatmapResponse} from "$lib/gen/zelvy/summary/get_summary_heatmap_response";
import {getFormattedDateParams, setDefaultDateParams} from "$lib/utils/setDefaultDateParams";
import {getSummaryHeatmap} from "$lib/server/grpc/summary";
import {getWinners} from "$lib/server/grpc/user";
import {getWorkouts} from "$lib/server/grpc/workout";
import type {GetWorkoutsResponse} from "$lib/gen/zelvy/workout/get_workouts_response";


export const load: PageServerLoad = async ({url}): Promise<{
    winRes: GetWinnersResponse,
    hmRes: GetSummaryHeatmapResponse,
    wkrRes: GetWorkoutsResponse
} | null> => {
    const {formattedSD, formattedED, isInvalid} = getFormattedDateParams(url);

    if (isInvalid) {
        setDefaultDateParams(url, true);
    }

    const winnersResponse = await getWinners(formattedSD, formattedED);
    const heatmapResponse = await getSummaryHeatmap(formattedSD, formattedED);
    const workoutsResponse = await getWorkouts(formattedSD, formattedED);

    return {
        winRes: winnersResponse,
        hmRes: heatmapResponse,
        wkrRes: workoutsResponse
    };
};
