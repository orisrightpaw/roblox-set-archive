<script lang="ts">
    import './layout.css';
    import 'remixicon/fonts/remixicon.css';

    import favicon from '$lib/assets/favicon.svg';
    import Navlink from '$lib/components/Navlink.svelte';

    let { children } = $props();
</script>

<svelte:head>
    <link rel="icon" href={favicon} />
</svelte:head>

<div class="container mx-auto px-32 min-h-screen">
    <nav class="bg-zinc-800 text-xl p-4 rounded-b-xl">
        <div class="text-2xl font-semibold inline-block mr-2">
            <a href="/">
                <img class="h-7 inline-block -translate-y-0.5 mr-0.5" src={favicon} alt="Website icon" />
                Roblox Set Archive
            </a>
        </div>
        <Navlink href="/" name="Home"></Navlink>
        <Navlink href="/sets" name="Sets"></Navlink>
        <Navlink href="/users" name="Users"></Navlink>
    </nav>

    <div class="my-5 px-4">
        {@render children()}
    </div>
</div>

<div class="bg-zinc-950/30 p-4 rounded-b-xl text-zinc-500">
    <div class="container mx-auto px-32">
        <p>made by me :)</p>
    </div>
</div>
