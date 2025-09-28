<script lang="ts">
    import type {PageProps} from "./$types";
    import GoalCard from "$lib/ui/home/GoalCard.svelte";
    import WorkoutListItem from "$lib/ui/home/WorkoutListItem.svelte";
    import ViewSelector from "$lib/ui/home/ViewSelector.svelte";

    const {data}: PageProps = $props();
</script>

<svelte:head>
    <title>Zelvy dashboard - Daily</title>
    <meta name="description" content="Today's stats from Zelvy"/>
</svelte:head>

<section class="flex flex-col gap-6">
    <section class="flex flex-col gap-0">
        <h3 class="text-xl font-medium">{data.summary.day}</h3>
        <div class="flex gap-2">
            <span class="text-base-content/60">Winner:</span>
            <span class="text-base">{data.summary.winner?.name}</span>
        </div>
    </section>
    <ViewSelector/>
    <section class="carousel w-full bg-base-900 gap-3 md:gap-0 md:rounded-lg">
        {#each data.summary.goals as goal}
            <GoalCard goal={goal}/>
        {/each}
    </section>
    <section class="flex flex-col gap-2 w-full md:mx-auto md:flex-row">
        {#each data.summary.workouts as workout (workout.id)}
            <WorkoutListItem workout={workout}/>
        {/each}
    </section>
</section>
