<script lang="ts">
    import {WorkoutViewModel} from "$lib/gen/zelvy/workout/workout_view_model";
    import Strength from "$lib/assets/img/strength.png";
    import Cardio from "$lib/assets/img/cardio.png";
    import {WorkoutActivityType, workoutActivityTypeToJSON} from "$lib/gen/zelvy/workout/workout_activity_type_enum";

    type Props = {
        workout: WorkoutViewModel
        index: number
        dataTestId: string
    }
    const {workout, index, dataTestId}: Props = $props();
    const picto = workout.activityType === workoutActivityTypeToJSON(WorkoutActivityType.STRENGTH) ? Strength : Cardio;
    const date = new Date(workout.doneAt);
    const time = `${date.getHours()}h${date.getMinutes()}`;
</script>

<li data-testid={dataTestId} class="gap-2 justify-center items-center w-5/6">
    {#if index !== 0}
        <hr class="bg-primary/50"/>
    {/if}
    <div class="timeline-start">
        <time class="text-base-content">{time}</time>
    </div>
    <div class="timeline-middle bg-primary rounded-full h-4 w-4">
    </div>
    <div class="timeline-end flex gap-2">
        <img
                src={picto}
                alt="{workout.activityType} picto"
                class="w-10 h-10 rounded-md"
        />
        <div class="flex flex-col">
            <p data-testid="workoutsTemplateWorkoutName" class="text-base-content">{workout.name}</p>
            <p data-testid="workoutsTemplateWorkoutDuration" class="text-base-content/60">{workout.duration}</p>
            <div>
            </div>
        </div>
    </div>
    {#if index === 0}
        <hr class="bg-primary/50"/>
    {/if}
</li>


