import type {GetSummaryHeatmapResponse} from "$lib/gen/zelvy/summary/get_summary_heatmap_response";
import type {GetSummaryResponse} from "$lib/gen/zelvy/summary/get_summary_response";
import {GoalType, goalTypeToJSON} from "$lib/gen/zelvy/goal/goal_type_enum";
import {
    WorkoutActivityType,
    workoutActivityTypeToJSON
} from "$lib/gen/zelvy/workout/workout_activity_type_enum";

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

export const mockGetSummaryResponse = (): GetSummaryResponse => {
    return {
        id: "1",
        day: new Date().toISOString(),
        winner: {
            name: "winner",
            discordId: "123"
        },
        workouts: [
            {
                id: "1",
                name: "Gym",
                kcalBurned: 295,
                picto: "",
                doneAt: new Date().toISOString(),
                duration: "3784",
                activityType: workoutActivityTypeToJSON(WorkoutActivityType.STRENGTH)
            }
        ],
        goals: [
            {
                type: goalTypeToJSON(GoalType.KCAL_BURNED),
                name: "Kcal burned",
                isOff: false,
                value: 578,
                displayValue: "578",
                threshold: 500,
                displayThreshold: "500",
                progression: 100,
                picto: "",
                isSuccessful: true
            }
        ]
    };

};
