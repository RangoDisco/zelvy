<script lang="ts">
    type Props = {
        picto: string;
        title: string;
        subtitle: string;
        value: string | number;
    }

    const {picto, title, subtitle, value}: Props = $props();

</script>

<article class="w-[48%] flex flex-col bg-base-200 rounded-lg p-4">
    <div class="flex flex-row flex-wrap items-center gap-1">
        <p>{picto}</p>
        <p class="font-bold">{title}</p>
    </div>
    <p class="text-sm text-base-content/60">{subtitle}</p>
    <p class="text-sm font-semibold">{value}</p>
</article>
