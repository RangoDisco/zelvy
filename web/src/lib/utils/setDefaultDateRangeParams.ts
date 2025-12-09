import {redirect} from "@sveltejs/kit";

export default function setDefaultDateRangeParams(url: URL) {
    const endDate = new Date();
    const startDate = new Date();
    startDate.setMonth(endDate.getMonth() - 3);

    url.searchParams.set("end_date", endDate.toISOString().slice(0, 10));
    url.searchParams.set("start_date", startDate.toISOString().slice(0, 10));

    redirect(307, `${url.pathname}?${url.searchParams.toString()}`);
}
