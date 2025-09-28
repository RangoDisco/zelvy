import type {PageServerLoad} from "./$types";
import {credentials} from "@grpc/grpc-js";
import {UserServiceClient} from "$lib/gen/zelvy/user/user_service";
import {GetWinnersRequest} from "$lib/gen/zelvy/user/get_winners_request";
import type {GetWinnersResponse} from "$lib/gen/zelvy/user/get_winners_response";
import {createMetadataWithAuth} from "$lib/server/grpc";


const LIMIT = 10;

export const load: PageServerLoad = async ({params}): Promise<GetWinnersResponse> => {
    const endDate = new Date();
    const startDate = new Date(endDate.getFullYear(), endDate.getMonth() - 1, endDate.getDate());

    const client = new UserServiceClient("localhost:50051", credentials.createInsecure());
    const req = GetWinnersRequest.create();
    req.endDate = endDate.toISOString().slice(0, 10);
    req.startDate = startDate.toISOString().slice(0, 10);
    req.limit = LIMIT;

    const res: GetWinnersResponse = await new Promise((resolve, reject) => {
        client.getWinners(req, createMetadataWithAuth(), (err, response) => {
            if (err) reject(err);
            else resolve(response);
        });
    });

    return res;
};
