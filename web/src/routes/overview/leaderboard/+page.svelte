<script lang="ts">
    import type {PageProps} from "./$types";
    import LeaderboardFilter from "$lib/ui/home/LeaderboardFilter.svelte";
    import WinnerListItem from "$lib/ui/home/WinnerListItem.svelte";

    const {data}: PageProps = $props();
</script>

<svelte:head>
    <title>Zelvy dashboard - Leaderboard</title>
    <meta name="description" content="Leaderboard"/>
</svelte:head>
<section class="w-full self-start flex gap-2">
    <div class="flex flex-col gap-6">
        <h1 class="text-xl md:text-3xl">Last 3 months</h1>
        <LeaderboardFilter/>
    </div>
</section>

<section class="flex flex-col gap-3 w-full md:mx-auto">
    {#each data.winRes.winners as winner, i (winner.username)}
        <WinnerListItem winner={winner} index={i}/>
    {/each}
</section>
