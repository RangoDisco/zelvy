<script lang="ts">
    import "../app.css";
    import favicon from "$lib/assets/favicon.svg";
    import ViewSelector from "$lib/ui/home/ViewSelector.svelte";

    let {children} = $props();
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
