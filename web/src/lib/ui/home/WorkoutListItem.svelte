<script lang="ts">
    import {WorkoutViewModel} from "$lib/gen/zelvy/workout/workout_view_model";

    type Props = {
        workout: WorkoutViewModel
    }

    let {workout}: Props = $props();
</script>

<article data-testid="workoutsTemplateWorkout"
         class="flex items-center justify-between w-full bg-base-200 rounded-lg h-[70px] px-4">
    <div class="w-full flex gap-4 items-center">
        <!--        TODO: HANDLE PICTURE-->
        <img
                src=""
                alt="wor"
                class="w-10 h-10 rounded-md"
        />
        <div>
            <p data-testid="workoutsTemplateWorkoutName" class="text-lg text-base-content">{workout.name}</p>
            <p data-testid="workoutsTemplateWorkoutDuration" class="text-md text-base-content">{workout.duration}</p>
        </div>
    </div>
    <div data-testid="workoutsTemplateWorkoutActivityType"
         class="self-end pb-2 text-base">{ workout.activityType.toLowerCase()}</div>
</article>


