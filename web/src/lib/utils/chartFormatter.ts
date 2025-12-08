import type {WorkoutViewModel} from "$lib/gen/zelvy/workout/workout_view_model";

type SetValue = {
    [key: string]: number;
}


export const parseWorkouts = (workouts: WorkoutViewModel[]) => {
    const labels: string[] = [];
    workouts.forEach(workout => {
        if (!labels.includes(workout.activityType)) {
            labels.push(workout.activityType);
        }
    });

    const set = workouts.reduce((acc, w) => {
        acc[w.activityType] = (acc[w.activityType] || 0) + 1;
        return acc;
    }, {} as SetValue);

    return {
        labels,
        values: labels.map(label => set[label]),
    };
};
