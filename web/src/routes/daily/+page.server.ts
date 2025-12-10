import type {PageServerLoad} from "./$types";
import type {GetSummaryResponse} from "$lib/gen/zelvy/summary/get_summary_response";
import setDefaultDateRangeParams from "$lib/utils/setDefaultDateParams";
import {isDate} from "node:util/types";
import {getSummary} from "$lib/server/grpc/summary";

export const csr = false;

export const load: PageServerLoad = async ({url}): Promise<{ summary: GetSummaryResponse | null } | null> => {
    const dateParam = url.searchParams.get("date");

    if (dateParam === null) {
        setDefaultDateRangeParams(url, false);
        return null;
    }

    const date = new Date(dateParam);
    if (!isDate(date) || new Date(date) > new Date()) {
        setDefaultDateRangeParams(url, false);
        return null;
    }

    const parsedDate = date.toISOString().slice(0, 10);
    try {
        const res = await getSummary(parsedDate);

        return {
            summary: res
        };
    } catch (error) {
        return {
            summary: null
        };
    }
};
