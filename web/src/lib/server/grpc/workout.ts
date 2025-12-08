import {createMetadataWithAuth} from "$lib/server/grpc/metadata";
import type {GetSummaryHeatmapResponse} from "$lib/gen/zelvy/summary/get_summary_heatmap_response";
import {credentials} from "@grpc/grpc-js";
import {GetWorkoutsRequest} from "$lib/gen/zelvy/workout/get_workouts_request";
import {WorkoutServiceClient} from "$lib/gen/zelvy/workout/workout_service";
import type {GetWorkoutsResponse} from "$lib/gen/zelvy/workout/get_workouts_response";

export const getWorkouts = async (formattedSD: string, formattedED: string) => {
    const client = getClient();
    const req = GetWorkoutsRequest.create();
    req.startDate = formattedSD;
    req.endDate = formattedED;

    return await new Promise((resolve, reject) => {
        client.getWorkouts(req, createMetadataWithAuth(), (err, response) => {
            if (err) reject(err);
            else resolve(response);
        });
    }) as GetWorkoutsResponse;
};

function getClient() {
    return new WorkoutServiceClient("localhost:50051", credentials.createInsecure());
}
