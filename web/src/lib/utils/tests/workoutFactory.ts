import type {GetWorkoutsResponse} from "$lib/gen/zelvy/workout/get_workouts_response";

export const mockGetWorkoutsResponse = (): GetWorkoutsResponse => {
    return {
        workouts: [
            {
                id: "1",
                name: "Wrkt1",
                activityType: "STRENGTH",
                duration: "3600",
                doneAt: new Date().toISOString(),
                kcalBurned: 429,
                picto: "picto"
            },
            {
                id: "2",
                name: "Wrkt2",
                activityType: "CYLING",
                duration: "1800",
                doneAt: new Date().toISOString(),
                kcalBurned: 329,
                picto: "picto"
            }
        ]
    };
};
