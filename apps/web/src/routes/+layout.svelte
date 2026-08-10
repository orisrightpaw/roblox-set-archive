<script lang="ts">
    import './layout.css';
    import favicon from '$lib/assets/favicon.svg';
    import { page } from '$app/state';
    import Navlink from '$lib/components/Navlink.svelte';

    let { children } = $props();

    page.url.pathname;
</script>

<svelte:head>
    <link rel="icon" href={favicon} />
</svelte:head>

<div class="container mx-auto px-32">
    <nav class="bg-zinc-800 text-xl p-4 rounded-b-xl">
        <div class="text-2xl inline-block mr-2">
            <a href="/">
                <img class="h-7 inline-block -translate-y-0.5 mr-0.5" src={favicon} alt="Website icon" />
                Roblox Set Archive
            </a>
        </div>
        <Navlink href="/" name="Home"></Navlink>
        <Navlink href="/sets" name="Sets"></Navlink>
        <Navlink href="/users" name="Users"></Navlink>
    </nav>

    <div class="my-5 px-3">
        {@render children()}
    </div>
</div>
