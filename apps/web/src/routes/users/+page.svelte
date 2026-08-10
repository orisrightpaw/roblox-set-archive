<script lang="ts">
    import Image from '$lib/components/Image.svelte';
    import Meta from '$lib/components/Meta.svelte';
    import { onMount } from 'svelte';

    let keyword = $state('');
    let page = $state(0);
    let pages: number = $state(0);
    let results: { id: number; user_name: string }[] = $state([]);

    async function search() {
        const response = await fetch(`http://127.0.0.1:3000/api/users/search?keyword=${keyword}&page=${page}`);
        if (!response.ok) {
            return;
        }

        const body = (await response.json()) as {
            pages: number;
            results: {
                id: number;
                user_name: string;
            }[];
        };

        pages = body.pages;
        results = body.results;
    }

    onMount(() => search());
</script>

<Meta title="Search Users"></Meta>

<form
    class="flex w-full gap-2 mb-4"
    onsubmit={() => {
        page = 0;
        search();
    }}
>
    <input
        class="rounded-md border border-zinc-700 text-lg grow p-2"
        type="text"
        placeholder="Search Users"
        bind:value={keyword}
    />
    <button class="bg-zinc-800 rounded-lg px-4 cursor-pointer" title="Search">
        <i class="ri-search-line"></i>
    </button>
</form>

<div class="grid grid-cols-8 gap-4 mb-4">
    {#each results as { id, user_name }, i}
        <div class="group truncate text-blue-400 hover:text-blue-300">
            <a href="/users/{id}">
                <Image src="https://sets.starfall.wtf/api/users/{id}/thumbnail"></Image>
            </a>
            <a class="group-hover:text-blue-300 transition-colors w-fit" href="/users/{id}">
                {user_name}
            </a>
        </div>
    {/each}
</div>

<div class="text-xl flex mx-auto w-fit">
    <button
        class="cursor-pointer bg-zinc-800 rounded-l-lg p-2 border-r border-zinc-700"
        title="First page"
        onclick={() => {
            page = 0;
            search();
        }}
    >
        <i class="ri-arrow-left-double-line"></i>
    </button>
    <button
        class="cursor-pointer bg-zinc-800 p-2 border-r border-zinc-700"
        title="Previous page"
        onclick={() => {
            page = Math.max(page - 1, 0);
            search();
        }}
    >
        <i class="ri-arrow-left-s-line"></i>
    </button>
    <p class="text-base my-auto bg-zinc-800 p-3 border-r border-zinc-700 font-mono">
        {page + 1} out of {pages + 1}
    </p>
    <button
        class="cursor-pointer bg-zinc-800 p-2 border-r border-zinc-700"
        title="Next page"
        onclick={() => {
            page = Math.min(pages, page + 1);
            search();
        }}
    >
        <i class="ri-arrow-right-s-line"></i>
    </button>
    <button
        class="cursor-pointer bg-zinc-800 rounded-r-lg p-2"
        title="Last page"
        onclick={() => {
            page = pages;
            search();
        }}
    >
        <i class="ri-arrow-right-double-line"></i>
    </button>
</div>
