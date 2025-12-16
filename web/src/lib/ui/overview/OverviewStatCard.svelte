<script lang="ts">
    import {ArrowRight} from "@lucide/svelte";

    type Props = {
        picto: string;
        title: string;
        subtitle: string;
        value: string | number;
        link: string | null;
    }

    const {picto, title, subtitle, value, link}: Props = $props();
</script>

<article class="w-[48%] md:w-[32%] flex flex-col bg-base-200 rounded-lg p-4 pr-0">
    <div class="flex flex-row items-center gap-1">
        <p class="text-sm lg:text-lg">{picto}</p>
        <p class="text-sm lg:text-lg font-bold">{title}</p>
    </div>
    <p class="text-sm lg:text-md text-base-content/60">{subtitle}</p>
    <div class="flex flex-row items-center justify-between gap-1 pr-4">
        <p class="text-sm lg:text-md font-semibold">{value}</p>
        {#if link !== null}
            <a href={link} target="_blank" rel="noopener noreferrer" class="bg-primary/50 rounded-full p-1">
                <ArrowRight size="16"/>
            </a>
        {/if}
    </div>
</article>
