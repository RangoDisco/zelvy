<script lang="ts">
    import type {PageProps} from "./$types";
    import ViewSelector from "$lib/ui/ViewSelector.svelte";

    let {data}: PageProps = $props();
</script>


<main>
    <ViewSelector/>
    <p>{data.summary.id}</p>
</main>
