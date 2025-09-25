import type {PageServerLoad} from "./$types";
import {credentials} from "@grpc/grpc-js";
import {SummaryServiceClient} from "$lib/gen/zelvy/summary/summary_service";
import {GetSummaryResquest} from "$lib/gen/zelvy/summary/get_summary_request";
import type {GetSummaryResponse} from "$lib/gen/zelvy/summary/get_summary_response";
import {createMetadataWithAuth} from "$lib/server/grpc";

export const csr = false;

export const load: PageServerLoad = async ({params}): Promise<{ summary: GetSummaryResponse }> => {

    const client = new SummaryServiceClient("localhost:50051", credentials.createInsecure());
    const req = GetSummaryResquest.create();
    const res: GetSummaryResponse = await new Promise((resolve, reject) => {
        client.getSummary(req, createMetadataWithAuth(), (err, response) => {
            if (err) reject(err);
            else resolve(response);
        });
    });

    return {
        summary: res
    };
};
