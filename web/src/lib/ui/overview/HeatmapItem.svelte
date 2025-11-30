<script lang="ts">
	import { HeatmapItemViewModel } from '$lib/gen/zelvy/summary/heatmap_item_view_model';

	type Props = {
		item: HeatmapItemViewModel;
	}

	const { item }: Props = $props();
	let color = $state();
	if (item.success === true) {
		color = item.id !== '' ? 'bg-secondary' : 'bg-secondary/50';
	} else {
		color = 'bg-red-500';
	}

</script>

<a href="/daily?date={item.date}" aria-label="date's details"
	 class="min-h-8 max-h-10 aspect-square bg-secondary rounded-md {color}"></a>
