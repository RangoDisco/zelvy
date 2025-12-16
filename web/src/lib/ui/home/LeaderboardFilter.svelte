<script lang="ts">
    import {ArrowUpDown} from "@lucide/svelte";
    import {page} from "$app/state";
    import {WinnerFilterType, winnerFilterTypeToJSON} from "$lib/gen/zelvy/user/get_winners_request";
    import {afterNavigate, goto, invalidateAll} from "$app/navigation";

    let isOpen = $state(false);
    const initialFilter = $derived(page.url.searchParams.get("filter"));
    let label = $state();

    afterNavigate(() => {
        isOpen = false;
        switch (initialFilter) {
            case winnerFilterTypeToJSON(WinnerFilterType.RELEVANT):
                label = "Relevant";
                break;
            case winnerFilterTypeToJSON(WinnerFilterType.IRRELEVANT):
                label = "Irrelevant";
                break;
            default:
                label = "Overall";
        }
    });

    const handleFilterChange = (value: WinnerFilterType | null) => {
        if (value === initialFilter) {
            return;
        }

        if (value) {
            const filterValue = winnerFilterTypeToJSON(value);
            page.url.searchParams.set("filter", filterValue);
        } else {
            page.url.searchParams.delete("filter");
        }

        goto(page.url).then(() => {
            invalidateAll().then(() => {
            });
        });
    };
</script>

<details class="dropdown" open={isOpen}>
    <summary class="btn rounded-full bg-primary/60">
        <ArrowUpDown size="20"/>
        {label}
    </summary>
    <ul class="menu dropdown-content bg-base-200 rounded-box z-1 w-52 shadow-sm gap-2">
        <li>
            <button class="btn" onclick={() => handleFilterChange(null)}>Overall</button>
        </li>
        <li>
            <button class="btn" onclick={() => handleFilterChange(WinnerFilterType.RELEVANT)}>Relevant (real
                wins)
            </button>
        </li>
        <li>
            <button class="btn" onclick={() => handleFilterChange(WinnerFilterType.IRRELEVANT)}>Irrelevant
                (no money)
            </button>
        </li>
    </ul>

</details>
