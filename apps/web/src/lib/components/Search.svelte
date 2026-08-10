<script lang="ts">
    import type { PaginatedResult } from '$lib';
    import { onMount, type Snippet } from 'svelte';

    let {
        endpoint,
        snippet,
        placeholder
    }: {
        endpoint: string;
        snippet: (result: any) => ReturnType<Snippet>;
        placeholder?: string;
    } = $props();

    let page = $state(0);
    let keyword = $state('');

    let pages = $state(0);
    let total = $state(0);
    let results: any[] = $state([]);

    async function search(k: string, p: number) {
        const response = await fetch(`${endpoint}?keyword=${k}&page=${p}`);
        if (!response.ok) return;

        const body = (await response.json()) as PaginatedResult<any>;

        pages = body.pages;
        total = body.total;
        results = body.results;
    }

    onMount(() => search('', 0));

    $effect(() => {
        search(keyword, page);
    });
</script>

<form class="flex w-full gap-2 mb-4">
    <input class="rounded-md border border-zinc-700 text-lg grow p-2" type="text" {placeholder} bind:value={keyword} />
    <button class="bg-zinc-800 rounded-lg px-4 cursor-pointer" title="Search">
        <i class="ri-search-line"></i>
    </button>
</form>

<div class="grid grid-cols-8 gap-4 mb-4">
    {#each results as result}
        {@render snippet(result)}
    {/each}
</div>

<div class="w-full flex">
    <div>
        <p class="text-sm text-zinc-400">
            Showing {(page * 24 + 1).toLocaleString()} to {((page + 1) * 24).toLocaleString()} of {total.toLocaleString()} results
        </p>
    </div>
</div>
