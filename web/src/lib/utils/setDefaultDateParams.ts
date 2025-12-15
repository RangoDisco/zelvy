import {redirect} from "@sveltejs/kit";
import {isDate} from "node:util/types";

export const setDefaultDateParams = (url: URL, isRange: boolean | null = false) => {
    const startDate = new Date();
    const endDate = new Date();

    if (isRange) {
        startDate.setMonth(endDate.getMonth() - 3);
        url.searchParams.set("end_date", endDate.toISOString().slice(0, 10));
        url.searchParams.set("start_date", startDate.toISOString().slice(0, 10));
    } else {
        endDate.setDate(endDate.getDate() - 1);
        url.searchParams.set("date", endDate.toISOString().slice(0, 10));
    }

    redirect(307, `${url.pathname}?${url.searchParams.toString()}`);
};

export const getFormattedDateParams = (url: URL) => {
    const eDParam = url.searchParams.get("end_date");
    const sDParam = url.searchParams.get("start_date");

    if (eDParam === null || sDParam === null) {
        return {
            formattedED: "",
            formattedSD: "",
            isInvalid: true,
        };
    }

    const endDate = new Date(eDParam);
    const startDate = new Date(sDParam);

    if (startDate > endDate || !isDate(endDate) || !isDate(startDate)) {
        return {
            formattedED: "",
            formattedSD: "",
            isInvalid: true,
        };
    }

    const formattedED = endDate.toISOString().slice(0, 10);
    const formattedSD = startDate.toISOString().slice(0, 10);

    return {
        formattedED,
        formattedSD,
        isInvalid: false
    };
};
