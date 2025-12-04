<script lang="ts">
    import type {PageProps} from "./$types";
    import ViewSelector from "$lib/ui/home/ViewSelector.svelte";
    import HeatmapItem from "$lib/ui/overview/HeatmapItem.svelte";
    import OverviewStatCard from "$lib/ui/overview/OverviewStatCard.svelte";
    import {formatSuccessRate, formatWinner, formatLongestStreak} from "$lib/utils/formatOverviewStats";

    const {data}: PageProps = $props();
    const {winRes, hmRes} = data;
    const rowsNumber = hmRes.items.length > 50 ? 7 : 4;
    const rowString = rowsNumber === 7 ? "grid-rows-7" : "grid-rows-4";
    const gridTemplate = `${rowString} grid-cols-${Math.ceil(hmRes.items.length / rowsNumber)}`;

</script>

<svelte:head>
    <title>Zelvy dashboard - Overview</title>
    <meta name="description" content="Stats overview"/>
</svelte:head>
<section class="flex flex-col gap-6">
    <ViewSelector/>
    <section class="flex flex-row justify-center flex-wrap gap-3">
        <OverviewStatCard {...formatWinner(winRes.winners)}/>
        <OverviewStatCard {...formatSuccessRate(hmRes.items)}/>
        <OverviewStatCard {...formatLongestStreak(hmRes.items)}/>
        <OverviewStatCard picto="O" title="KanaPei" subtitle="Most wins" value="35"/>
    </section>
    <section
            class="bg-base-200 grid grid-flow-col {gridTemplate} gap-1 rounded-lg overflow-scroll p-4">
        {#each hmRes.items as item (item.date)}
            <HeatmapItem item={item}/>
        {/each}
    </section>
</section>
