<script lang="ts">
    import type {PageProps} from "./$types";
    import {page} from "$app/state";
    import {navigationStates} from "$lib/state.svelte";
    import HeatmapItem from "$lib/ui/overview/HeatmapItem.svelte";
    import OverviewStatCard from "$lib/ui/overview/OverviewStatCard.svelte";
    import {formatSuccessRate, formatWinner, formatLongestStreak} from "$lib/utils/formatOverviewStats";
    import {getNewPeriodTitle, changeOverviewPeriod} from "$lib/utils/periodChanger";
    import {parseWorkouts} from "$lib/utils/chartFormatter";
    import Bar from "$lib/ui/charts/Bar.svelte";
    import {ArrowLeft, ArrowRight} from "@lucide/svelte";

    const {data}: PageProps = $props();
    const rowsNumber = $derived(data.hmRes.items.length > 50 ? 7 : 4);
    const rowString = $derived(rowsNumber === 7 ? "grid-rows-7" : "grid-rows-4");
    const gridTemplate = $derived(`${rowString} grid-cols-${Math.ceil(data.hmRes.items.length / rowsNumber)}`);

    const period = $derived(getNewPeriodTitle(data.hmRes.items));
    const workoutRadarData = $derived(parseWorkouts(data.wkrRes.workouts));
</script>

<svelte:head>
    <title>Zelvy dashboard - Overview</title>
    <meta name="description" content="Stats overview"/>
</svelte:head>

{#if navigationStates.isLoading}
    <div class="flex w-full items-center justify-center">
    <span class="loading loading-spinner loading-md">

    </span>
    </div>
{:else}
    <div class="flex flex-col gap-6">
        <section class="flex flex-row justify-between items-center">
            <h2 class="text-xl md:text-3xl" data-testid="overviewPeriod">{period}</h2>
            <div class="flex flex-row gap-2">
                <button class="btn btn-square" aria-label="previous period"
                        onclick={() => {changeOverviewPeriod(page.url, "previous")}}
                        data-testid="overviewPreviousButton">
                    <ArrowLeft data-testid="overviewPreviousIcon"/>
                </button>
                <button class="btn btn-square" aria-label="next period"
                        onclick={() => {changeOverviewPeriod(page.url, "next")}} data-testid="overviewNextButton">
                    <ArrowRight/>
                </button>
            </div>
        </section>
        <section class="flex flex-row justify-start lg:justify-between flex-wrap gap-3">
            <OverviewStatCard {...formatWinner(data.winRes.winners)} dataTestId="overviewWinnerCard"/>
            <OverviewStatCard {...formatSuccessRate(data.hmRes.items)} dataTestId="overviewSuccessRateCard"/>
            <OverviewStatCard {...formatLongestStreak(data.hmRes.items)} dataTestId="overviewLongestStreakCard"/>
        </section>
        <div class="flex flex-row flex-wrap md:items-start justify-center md:justify-between gap-6 md:gap-2">
            <section class="w-full flex flex-col gap-4 md:w-[49%] bg-base-200 rounded-lg p-6">
                <h3 class="text-lg md:text-2xl">Heatmap</h3>
                <div class="grid grid-flow-col {gridTemplate} gap-1 overflow-auto" data-testId="overviewHeatmapGrid">
                    {#each data.hmRes.items as item, i (item.date)}
                        <HeatmapItem item={item} dataTestId={`overviewHeatmapGridItem-${i}`}/>
                    {/each}
                </div>
            </section>
            <section class="flex flex-col gap-2 w-full md:w-[49%] bg-base-200 rounded-lg p-6">
                <h3 class="text-lg md:text-2xl">Workouts</h3>
                <Bar data={workoutRadarData} dataTestId="overviewBarChart"/>
            </section>
        </div>
    </div>
{/if}
