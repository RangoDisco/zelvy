<script lang="ts">
    type Props = {
        picto: string;
        title: string;
        subtitle: string;
        value: string | number;
    }

    const {picto, title, subtitle, value}: Props = $props();

</script>

<article class="w-[48%] md:w-[32%] flex flex-col bg-base-200 rounded-lg p-4 pr-0">
    <div class="flex flex-row items-center gap-1">
        <p class="text-sm lg:text-md">{picto}</p>
        <p class="text-sm lg:text-md font-bold">{title}</p>
    </div>
    <p class="text-sm text-base-content/60">{subtitle}</p>
    <p class="text-sm font-semibold">{value}</p>
</article>
