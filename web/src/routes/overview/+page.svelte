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
    import {ArrowLeft, ArrowRight} from "@lucide/svelte";

    const {data}: PageProps = $props();
    const rowsNumber = $state(data.hmRes.items.length > 50 ? 7 : 4);
    const rowString = $state(rowsNumber === 7 ? "grid-rows-7" : "grid-rows-4");
    const gridTemplate = $state(`${rowString} grid-cols-${Math.ceil(data.hmRes.items.length / rowsNumber)}`);

    const period = $derived(getNewPeriodTitle(data.hmRes.items));
    const workoutRadarData = $derived(parseWorkouts(data.wkrRes.workouts));
</script>

<svelte:head>
    <title>Zelvy dashboard - Overview</title>
    <meta name="description" content="Stats overview"/>
</svelte:head>
<div class="flex flex-col gap-6">
    <ViewSelector/>
    <section class="flex flex-row justify-between items-center">
        <h2 class="text-xl md:text-3xl">{period}</h2>
        <div class="flex flex-row gap-2">
            <button class="btn btn-square" aria-label="previous period"
                    onclick={() => {handlePeriodChange(page.url, "previous")}}>
                <ArrowLeft/>
            </button>
            <button class="btn btn-square" aria-label="next period"
                    onclick={() => {handlePeriodChange(page.url, "next")}}>
                <ArrowRight/>
            </button>
        </div>
    </section>
    <section class="flex flex-col gap-2">
        <h3 class="text-lg md:text-2xl">Stats</h3>
        <div class="flex flex-row justify-center md:justify-start lg:justify-between flex-wrap gap-3">
            <OverviewStatCard {...formatWinner(data.winRes.winners)}/>
            <OverviewStatCard {...formatSuccessRate(data.hmRes.items)}/>
            <OverviewStatCard {...formatLongestStreak(data.hmRes.items)}/>
            <OverviewStatCard picto="O" title="KanaPei" subtitle="Most wins" value="35"/>
        </div>
    </section>
    <div class="flex flex-row flex-wrap md:items-start justify-center md:justify-between gap-6 md:gap-2">
        <section class="w-full flex flex-col gap-2 md:w-[49%]">
            <h3 class="text-lg md:text-2xl">Heatmap</h3>
            <div
                    class="bg-base-200 grid grid-flow-col {gridTemplate} gap-1 rounded-lg overflow-auto p-6">
                {#each data.hmRes.items as item (item.date)}
                    <HeatmapItem item={item}/>
                {/each}
            </div>
        </section>
        <section class="flex flex-col gap-2 w-full md:w-[49%] p-1">
            <h3 class="text-lg md:text-2xl">Workouts</h3>
            <Bar data={workoutRadarData}/>
        </section>
    </div>
</div>
