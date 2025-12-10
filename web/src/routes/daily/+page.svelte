<script lang="ts">
    import type {PageProps} from "./$types";
    import {page} from "$app/state";
    import GoalCard from "$lib/ui/home/GoalCard.svelte";
    import WorkoutListItem from "$lib/ui/home/WorkoutListItem.svelte";
    import {changeDailyPeriod} from "$lib/utils/periodChanger";
    import {ArrowLeft, ArrowRight} from "@lucide/svelte";
    import {navigationStates} from "$lib/state.svelte";

    const {data}: PageProps = $props();
</script>

<svelte:head>
    <title>Zelvy dashboard - Daily</title>
    <meta name="description" content="Today's stats from Zelvy"/>
</svelte:head>

{#if navigationStates.isLoading}
    <div class="flex w-full items-center justify-center">
    <span class="loading loading-spinner loading-md">

    </span>
    </div>
{:else}
    <div class="flex flex-col gap-6">
        <div class="flex flex-row justify-between items-start">
            <div class="flex flex-col gap-0">
                <h3 class="text-xl font-medium">{data.summary?.day ?? page.url.searchParams.get("date")}</h3>
                {#if data.summary}
                    <div class="flex gap-2">
                        <span class="text-base-content/60">Winner:</span>
                        <span class="text-base">{data.summary.winner?.name}</span>
                    </div>
                {/if}
            </div>
            <div class="flex flex-row gap-2">
                <button class="btn btn-square" aria-label="previous period" disabled={navigationStates.isLoading}
                        onclick={() => {changeDailyPeriod(page.url, "previous")}}>
                    <ArrowLeft/>
                </button>
                <button class="btn btn-square" aria-label="next period" disabled={navigationStates.isLoading}
                        onclick={() => {changeDailyPeriod(page.url, "next")}}>
                    <ArrowRight/>
                </button>
            </div>
        </div>
        {#if data.summary}
            <section class="carousel w-full gap-3 bg-base-100">
                {#each data.summary.goals as goal, i (goal.name)}
                    <GoalCard goal={goal} index={i}/>
                {/each}
            </section>
            <section class="w-full flex flex-col p-4 gap-4 bg-base-200 rounded-lg md:w-1/3">
                <h1 class="text-2xl">Activity timeline</h1>
                <ul class="timeline timeline-snap-icon timeline-vertical">
                    {#each data.summary.workouts as workout, i (workout.id)}
                        <WorkoutListItem workout={workout} index={i}/>
                    {/each}
                </ul>
            </section>
        {:else}
            <section class="flex flex-col gap-6">
                <p class="text-xl">No data</p>
            </section>
        {/if}
    </div>
{/if}
