<script lang="ts">
    import type {WinnerViewModel} from "$lib/gen/zelvy/user/winner_view_model";

    type Props = {
        winner: WinnerViewModel
        index: number;
    }

    const {winner, index}: Props = $props();
    let color;
    switch (index) {
        case 0:
            color = "text-secondary";
            break;
        case 1:
            color = "text-base-content";
            break;
        case 2:
            color = "text-orange-400/60";
            break;
        default:
            color = "text-base-content/60";
            break;
    }
</script>

<article class="flex items-center justify-between w-full bg-base-200 rounded-lg h-[70px] px-4">
    <div class="w-full flex gap-4 items-center">
        <p class="font-bold items-start w-[20px] text-center {color}">{index + 1}</p>
        <p class="font-bold">{winner.username}</p>
    </div>
    <div class="font-medium items-center justify-self-end">{winner.wins}</div>
</article>
