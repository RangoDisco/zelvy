import type {GetSummaryHeatmapResponse} from "$lib/gen/zelvy/summary/get_summary_heatmap_response";

export const mockGetHeatmapResponse = (): GetSummaryHeatmapResponse => {
    return {
        items: [
            {
                date: new Date().toISOString(),
                success: true,
                id: "1"
            }
        ]
    };
};
