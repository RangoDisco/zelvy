<script lang="ts">
    import {page} from "$app/state";
</script>

<section class="join w-full md:w-2/5 bg-base-200 rounded-full">
    <a
            href="/daily"
            class="text-base font-medium join-item btn md:btn-sm rounded-full w-1/2 {page.url.pathname === '/daily' ? 'bg-primary' : ''}"
    >
        Daily
    </a>
    <a
            href="/overview"
            class="text-base font-medium join-item btn md:btn-sm rounded-full w-1/2 {page.url.pathname.includes('/overview') ? 'bg-primary' : ''}"
    >
        Overview
    </a>
</section>
