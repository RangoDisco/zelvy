<script lang="ts">
	import type { PageProps } from './$types';
	import ViewSelector from '$lib/ui/home/ViewSelector.svelte';
	import HeatmapItem from '$lib/ui/overview/HeatmapItem.svelte';
	import OverviewStatCard from '$lib/ui/overview/OverviewStatCard.svelte';
	import type { HeatmapItemViewModel } from '$lib/gen/zelvy/summary/heatmap_item_view_model';

	const { data }: PageProps = $props();
	const { winRes, hmRes } = data;
	const rowsNumber = hmRes.items.length > 50 ? 7 : 4;
	const gridTemplate = `grid-rows-${rowsNumber} grid-cols-${Math.ceil(hmRes.items.length / rowsNumber)}`;
	const winner = winRes.winners[0];

	const successNumber = hmRes.items.filter((item: HeatmapItemViewModel) => item.success === true).length;
	const successRate = successNumber / hmRes.items.length * 100;

</script>

<svelte:head>
	<title>Zelvy dashboard - Overview</title>
	<meta name="description" content="Stats overview" />
</svelte:head>
<section class="flex flex-col gap-6">
	<ViewSelector />
	<section class="flex flex-row justify-center flex-wrap gap-3">
		<OverviewStatCard picto="👑" title={winner.username} subtitle="Most wins" value={winner.wins} />
		<OverviewStatCard picto="🎯" title={`${successRate}%`} subtitle="Success rate"
											value={`${successNumber}/${hmRes.items.length}`} />
		<OverviewStatCard picto="O" title="KanaPei" subtitle="Most wins" value="35" />
		<OverviewStatCard picto="O" title="KanaPei" subtitle="Most wins" value="35" />
	</section>
	<section
		class="min-h-12 bg-base-200 grid grid-flow-col {gridTemplate} gap-1 rounded-lg overflow-scroll p-4">
		{#each hmRes.items as item (item.date)}
			<HeatmapItem item={item} />
		{/each}
	</section>
</section>
