import {SummaryServiceClient} from "$lib/gen/zelvy/summary/summary_service";
import {credentials} from "@grpc/grpc-js";
import {GetSummaryHeatmapRequest} from "$lib/gen/zelvy/summary/get_summary_heatmap_request";
import type {GetSummaryHeatmapResponse} from "$lib/gen/zelvy/summary/get_summary_heatmap_response";
import {createMetadataWithAuth} from "$lib/server/grpc/metadata";

export const getSummaryHeatmap = async (formattedSD: string, formattedED: string) => {
    const client = getClient();
    const heatmapReq = GetSummaryHeatmapRequest.create();
    heatmapReq.startDate = formattedSD;
    heatmapReq.endDate = formattedED;

    return await new Promise((resolve, reject) => {
        client.getSummaryHeatmap(heatmapReq, createMetadataWithAuth(), (err, response) => {
            if (err) reject(err);
            else resolve(response);
        });
    }) as GetSummaryHeatmapResponse;
};

function getClient() {
    return new SummaryServiceClient("localhost:50051", credentials.createInsecure());

}
