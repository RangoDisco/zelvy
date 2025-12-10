import {redirect} from "@sveltejs/kit";

export default function validateOrUpdateDateParams(url: URL, isRange: boolean | null = false) {
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
}
