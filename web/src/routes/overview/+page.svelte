<script lang="ts">
    import type {PageProps} from "./$types";
    import {page} from "$app/state";
    import ViewSelector from "$lib/ui/home/ViewSelector.svelte";
    import HeatmapItem from "$lib/ui/overview/HeatmapItem.svelte";
    import OverviewStatCard from "$lib/ui/overview/OverviewStatCard.svelte";
    import {formatSuccessRate, formatWinner, formatLongestStreak} from "$lib/utils/formatOverviewStats";
    import {getNewPeriodTitle, handlePeriodChange} from "$lib/utils/periodChanger";
    import {parseWorkouts} from "$lib/utils/chartFormatter";
    import Bar from "$lib/ui/charts/Bar.svelte";

    const {data}: PageProps = $props();
    let rowsNumber = $state(data.hmRes.items.length > 50 ? 7 : 4);
    let rowString = $state(rowsNumber === 7 ? "grid-rows-7" : "grid-rows-4");
    let gridTemplate = $state(`${rowString} grid-cols-${Math.ceil(data.hmRes.items.length / rowsNumber)}`);
    let period = $derived(getNewPeriodTitle(data.hmRes.items));
    const workoutRadarData = $derived(parseWorkouts(data.wkrRes.workouts));
</script>

<svelte:head>
    <title>Zelvy dashboard - Overview</title>
    <meta name="description" content="Stats overview"/>
</svelte:head>
<section class="flex flex-col gap-6">
    <ViewSelector/>
    <section class="flex flex-row justify-between items-center">
        <h2 class="text-xl md:text-3xl">{period}</h2>
        <div class="flex flex-row gap-2">
            <button class="btn btn-circle" aria-label="previous period"
                    onclick={() => {handlePeriodChange(page.url, "previous")}}>
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2.5"
                     stroke="currentColor" class="size-[1.2em]">
                    <path stroke-linecap="round" stroke-linejoin="round"
                          d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12Z"/>
                </svg>
            </button>
            <button class="btn btn-circle" aria-label="next period"
                    onclick={() => {handlePeriodChange(page.url, "next")}}>
                <svg xmlns="http://www.w3.org/2000/svg" fill="none" viewBox="0 0 24 24" stroke-width="2.5"
                     stroke="currentColor" class="size-[1.2em]">
                    <path stroke-linecap="round" stroke-linejoin="round"
                          d="M21 8.25c0-2.485-2.099-4.5-4.688-4.5-1.935 0-3.597 1.126-4.312 2.733-.715-1.607-2.377-2.733-4.313-2.733C5.1 3.75 3 5.765 3 8.25c0 7.22 9 12 9 12s9-4.78 9-12Z"/>
                </svg>
            </button>
        </div>
    </section>
    <section class="flex flex-col gap-2">
        <h3 class="text-lg md:text-2xl">Stats</h3>
        <section class="flex flex-row justify-center md:justify-between flex-wrap gap-3">
            <OverviewStatCard {...formatWinner(data.winRes.winners)}/>
            <OverviewStatCard {...formatSuccessRate(data.hmRes.items)}/>
            <OverviewStatCard {...formatLongestStreak(data.hmRes.items)}/>
            <OverviewStatCard picto="O" title="KanaPei" subtitle="Most wins" value="35"/>
        </section>
    </section>
    <section class="flex flex-row flex-wrap justify-center md:justify-between gap-6 md:gap-2">
        <section class="w-full flex flex-col gap-2 md:w-[49%]">
            <h3 class="text-lg md:text-2xl">Heatmap</h3>
            <section
                    class="bg-base-200 grid grid-flow-col {gridTemplate} gap-1 rounded-lg overflow-auto p-6">
                {#each data.hmRes.items as item (item.date)}
                    <HeatmapItem item={item}/>
                {/each}
            </section>
        </section>
        <section class="flex flex-col gap-2 md:w-[49%] p-1">
            <h3 class="text-lg md:text-2xl">Workouts</h3>
            <Bar data={workoutRadarData}/>
        </section>
    </section>
</section>
