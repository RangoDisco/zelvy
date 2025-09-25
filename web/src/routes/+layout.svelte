<script lang="ts">
    import "../app.css";
    import favicon from "$lib/assets/favicon.svg";
    import ViewSelector from "$lib/ui/ViewSelector.svelte";

    let {children} = $props();
</script>

<svelte:head>
    <link rel="icon" href={favicon}/>
</svelte:head>
<section class="bg-base-100 min-h-screen container mx-auto p-6 md:mx-none flex flex-col gap-6">
    <ViewSelector/>
    {@render children?.()}
</section>
