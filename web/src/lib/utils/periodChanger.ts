import {goto, invalidateAll} from "$app/navigation";
import type {HeatmapItemViewModel} from "$lib/gen/zelvy/summary/heatmap_item_view_model";

export const changeOverviewPeriod = (url: URL, direction: "previous" | "next") => {
    const params = url.searchParams;
    const rawED = params.get("end_date");

    if (!rawED) {
        return null;
    }

    let endDate = new Date(rawED);
    endDate.setMonth(direction === "next" ? endDate.getMonth() + 1 : endDate.getMonth() - 3);
    if (new Date() < endDate) {
        endDate = new Date();
    }

    const startDate = new Date(endDate.getFullYear(), endDate.getMonth() - 3);

    url.searchParams.set("end_date", endDate.toISOString().slice(0, 10));
    url.searchParams.set("start_date", startDate.toISOString().slice(0, 10));
    goto(url).then(() => {
        invalidateAll().then(() => {
        });
    });
};

export const changeDailyPeriod = (url: URL, direction: string) => {
    const params = url.searchParams;
    const dateParam = params.get("date");

    if (!dateParam) {
        return null;
    }

    let date = new Date(dateParam);
    date.setDate(direction === "next" ? date.getDate() + 1 : date.getDate() - 1);
    if (date > new Date()) {
        date = new Date();
    }

    url.searchParams.set("date", date.toISOString().slice(0, 10));
    goto(url).then(() => {
        invalidateAll().then(() => {
        });
    });
};

export const getNewPeriodTitle = (days: HeatmapItemViewModel[]) => {
    const length = days.length;
    const firstDay = days[0].date.slice(0, 10);
    const lastDay = days[length - 1].date.slice(0, 10);

    return `${firstDay} - ${lastDay}`;
};
