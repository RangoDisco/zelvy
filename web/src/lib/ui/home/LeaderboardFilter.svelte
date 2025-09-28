<script lang="ts">
    import {ArrowUpDown, CalendarDays} from "@lucide/svelte";

    type Props = {
        label: string;
        picto: string;
    }

    const {label, picto}: Props = $props();
    let isOpen = $state(null);
</script>

<details class="dropdown" open={isOpen}>
    <summary class="btn rounded-full bg-primary/60">
        {#if picto === "calendar"}
            <CalendarDays size="20"/>
        {:else}
            <ArrowUpDown size="20"/>
        {/if}
        {label}
    </summary>
    <ul class="menu dropdown-content bg-base-200 rounded-box z-1 w-52 shadow-sm">
        <li>
            <button disabled class="btn">Soon</button>
        </li>
    </ul>

</details>
