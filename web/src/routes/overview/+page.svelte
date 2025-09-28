<script lang="ts">
    import type {PageProps} from "./$types";
    import WinnerListItem from "$lib/ui/home/WinnerListItem.svelte";
    import LeaderboardFilter from "$lib/ui/home/LeaderboardFilter.svelte";
    import ViewSelector from "$lib/ui/home/ViewSelector.svelte";

    const {data}: PageProps = $props();

</script>

<svelte:head>
    <title>Zelvy dashboard - Overview</title>
    <meta name="description" content="Stats overview"/>
</svelte:head>

<section class="flex flex-col gap-6">
    <ViewSelector/>
    <section class="w-full self-start flex gap-2">
        <LeaderboardFilter label="This month" picto="calendar"/>
        <LeaderboardFilter label="Overall" picto="sort"/>
    </section>
    <section class="flex flex-col gap-3 w-full md:mx-auto">
        {#each data.winners as winner, i}
            <WinnerListItem winner={winner} index={i}/>
        {/each}
    </section>
</section>
