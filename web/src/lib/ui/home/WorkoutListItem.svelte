<script lang="ts">
    import {WorkoutViewModel} from "$lib/gen/zelvy/workout/workout_view_model";
    import Strength from "$lib/assets/img/strength.png";
    import Cardio from "$lib/assets/img/cardio.png";
    import {WorkoutActivityType, workoutActivityTypeToJSON} from "$lib/gen/zelvy/workout/workout_activity_type_enum";

    type Props = {
        workout: WorkoutViewModel
    }
    const {workout}: Props = $props();
    const picto = workout.activityType === workoutActivityTypeToJSON(WorkoutActivityType.STRENGTH) ? Strength : Cardio;
</script>

<article data-testid="workoutsTemplateWorkout"
         class="flex items-center justify-between w-full bg-base-200 rounded-lg h-[70px] px-4">
    <div class="w-full flex gap-4 items-center">
        <img
                src={picto}
                alt="{workout.activityType} picto"
                class="w-10 h-10 rounded-md"
        />
        <div>
            <p data-testid="workoutsTemplateWorkoutName" class="font-bold text-base-content">{workout.name}</p>
            <p data-testid="workoutsTemplateWorkoutDuration" class="text-base-content">{workout.duration}</p>
        </div>
    </div>
    <div data-testid="workoutsTemplateWorkoutActivityType"
         class="self-end pb-2 text-sm text-base-content/60">{ workout.activityType.toLowerCase()}</div>
</article>


