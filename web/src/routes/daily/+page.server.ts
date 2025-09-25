import type {PageLoad} from "./$types";
import {credentials} from "@grpc/grpc-js";
import {SummaryServiceClient} from "$lib/gen/zelvy/summary/summary_service";
import {GetSummaryResquest} from "$lib/gen/zelvy/summary/get_summary_request";
import type {GetSummaryResponse} from "$lib/gen/zelvy/summary/get_summary_response";
import {getMetadataWithAuth} from "$lib/server/grpc";

export const csr = false;

export const load: PageLoad = async ({params}) => {

    const client = new SummaryServiceClient("localhost:50051", credentials.createInsecure());
    const req = GetSummaryResquest.create();
    const res: GetSummaryResponse = await new Promise((resolve, reject) => {
        client.getSummary(req, getMetadataWithAuth(), (err, response) => {
            if (err) reject(err);
            else resolve(response);
        });
    });

    return {
        summary: res
    };
};
