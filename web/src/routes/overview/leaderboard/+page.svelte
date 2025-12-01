<script lang="ts">

	import LeaderboardFilter from '$lib/ui/home/LeaderboardFilter.svelte';
	import WinnerListItem from '$lib/ui/home/WinnerListItem.svelte';
</script>

<svelte:head>
	<title>Zelvy dashboard - Leaderboard</title>
	<meta name="description" content="Leaderboard" />
</svelte:head>
<section class="w-full self-start flex gap-2">
	<LeaderboardFilter label="This month" picto="calendar" />
	<LeaderboardFilter label="Overall" picto="sort" />
</section>

<section class="flex flex-col gap-3 w-full md:mx-auto">
	{#each winRes.winners as winner, i (winner.username)}
		<WinnerListItem winner={winner} index={i} />
	{/each}
</section>
