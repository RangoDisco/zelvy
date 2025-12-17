<script lang="ts">
    import "../app.css";
    import favicon from "$lib/assets/favicon.ico";
    import ViewSelector from "$lib/ui/home/ViewSelector.svelte";
    import {afterNavigate, beforeNavigate} from "$app/navigation";
    import {navigationStates} from "$lib/state.svelte";

    let {children} = $props();

    beforeNavigate(() => {
        navigationStates.isLoading = true;
    });

    afterNavigate(() => {
        navigationStates.isLoading = false;
    });
</script>

<svelte:head>
    <link rel="icon" href={favicon}/>
    <title>Zelvy dashboard</title>
    <meta name="description" content="Zelvy dashboard"/>
</svelte:head>

<section class="bg-base-100 min-h-screen container mx-auto p-6 md:max-w-[1200px] flex flex-col gap-6">
    <ViewSelector/>
    {@render children?.()}
</section>
