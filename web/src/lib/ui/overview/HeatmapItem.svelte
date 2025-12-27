<script lang="ts">
    import {HeatmapItemViewModel} from "$lib/gen/zelvy/summary/heatmap_item_view_model";

    type Props = {
        item: HeatmapItemViewModel;
        dataTestId: string;
    }

    const {item, dataTestId}: Props = $props();
    let color = $state();
    if (item.success === true) {
        color = item.id !== "" ? "bg-secondary" : "bg-secondary/50";
    } else {
        color = "bg-error/50";
    }

</script>

<a href="/daily?date={item.date.slice(0, 10)}" aria-label="date's details"
   class="min-h-8 max-h-10 aspect-square rounded-md {color}" data-testid={dataTestId}></a>
