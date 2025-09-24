<script lang="ts">
    import {page} from "$app/state";
</script>

<section class="join w-full bg-base-200 rounded-full">
    <a
            href="/daily"
            class="join-item btn rounded-full w-1/2 h-8 {page.url.pathname === '/daily' ? 'bg-primary' : ''}"
    >
        Daily
    </a>
    <a
            href="/overview"
            class="join-item btn rounded-full w-1/2 h-8 {page.url.pathname === '/overview' ? 'bg-primary' : ''}"
    >
        Overview
    </a>
</section>
