<script lang="ts">
	import type { PageProps } from './$types';
	import WinnerListItem from '$lib/ui/home/WinnerListItem.svelte';
	import LeaderboardFilter from '$lib/ui/home/LeaderboardFilter.svelte';
	import ViewSelector from '$lib/ui/home/ViewSelector.svelte';
	import HeatmapItem from '$lib/ui/overview/HeatmapItem.svelte';

	const { data }: PageProps = $props();
	const { winRes, hmRes } = data;
	const rowsNumber = hmRes.items.length > 50 ? 7 : 4;
	const gridTemplate = `grid-rows-${rowsNumber} grid-cols-${Math.ceil(hmRes.items.length / rowsNumber)}`;
</script>

<svelte:head>
	<title>Zelvy dashboard - Overview</title>
	<meta name="description" content="Stats overview" />
</svelte:head>

<section class="flex flex-col gap-6">
	<ViewSelector />
	<section class="w-full self-start flex gap-2">
		<LeaderboardFilter label="This month" picto="calendar" />
		<LeaderboardFilter label="Overall" picto="sort" />
	</section>
	<section
		class="bg-base-200 grid grid-flow-col {gridTemplate} gap-1 rounded-lg overflow-scroll p-4">
		{#each hmRes.items as item (item.date)}
			<HeatmapItem item={item} />
		{/each}
	</section>
	<section class="flex flex-col gap-3 w-full md:mx-auto">
		{#each winRes.winners as winner, i (winner.username)}
			<WinnerListItem winner={winner} index={i} />
		{/each}
	</section>
</section>
