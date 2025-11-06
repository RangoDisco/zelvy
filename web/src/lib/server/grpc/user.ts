import {UserServiceClient} from "$lib/gen/zelvy/user/user_service";
import {credentials} from "@grpc/grpc-js";
import {GetWinnersRequest} from "$lib/gen/zelvy/user/get_winners_request";
import type {GetWinnersResponse} from "$lib/gen/zelvy/user/get_winners_response";
import {createMetadataWithAuth} from "$lib/server/grpc/metadata";

const LIMIT = 10;

export const getWinners = async (formattedSD: string, formattedED: string) => {
    const client = getClient();
    const winnerReq = GetWinnersRequest.create();
    winnerReq.endDate = formattedED;
    winnerReq.startDate = formattedSD;
    winnerReq.limit = LIMIT;

    return await new Promise((resolve, reject) => {
        console.log(formattedSD);
        client.getWinners(winnerReq, createMetadataWithAuth(), (err, response) => {
            if (err) reject(err);
            else resolve(response);
        });
    }) as GetWinnersResponse;
};

const getClient = () => {
    return new UserServiceClient("localhost:50051", credentials.createInsecure());
};
