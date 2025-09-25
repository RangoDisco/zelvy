<script lang="ts">
    import type {PageProps} from "./$types";
    import GoalCard from "$lib/ui/home/GoalCard.svelte";
    import WorkoutListItem from "$lib/ui/home/WorkoutListItem.svelte";

    const {data}: PageProps = $props();
</script>


<section class="flex flex-col gap-6">
    <p>{data.summary.day}</p>
    <section class="carousel w-full bg-base-900 gap-2 md:gap-0 md:rounded-lg">
        {#each data.summary.goals as goal}
            <GoalCard goal={goal}/>
        {/each}
    </section>
    <section class="flex flex-col gap-2 w-full md:mx-auto">
        {#each data.summary.workouts as workout (workout.id)}
            <WorkoutListItem workout={workout}/>
        {/each}
    </section>
</section>
