<script lang="ts">
    import { page } from '$app/state';
    import Heading from '$lib/components/Heading.svelte';
    import Image from '$lib/components/Image.svelte';
    import Link from '$lib/components/Link.svelte';
    import Meta from '$lib/components/Meta.svelte';
    import Set from '$lib/components/Set.svelte';

    let { data } = $props();
</script>

{#if data.user}
    <Meta title="{data.user?.user_name!}'s Sets" />

    <div class="flex flex-col gap-4">
        <div class=" flex gap-4">
            <div class="w-48">
                <Image src="https://sets.starfall.wtf/api/users/{data.user.id}/thumbnail"></Image>
            </div>
            <div>
                <Heading>{data.user.user_name}</Heading>
                <div class="flex gap-3">
                    <Link href="https://www.roblox.com/users/{data.user.id}/profile">View Roblox Profile</Link>
                    <Link href="{page.url.origin}/Game/Tools/InsertAsset.ashx?nsets=20&type=user&userid={data.user.id}">
                        See LuaWebService XML
                    </Link>
                </div>
            </div>
        </div>
        <div class="mb-2">
            <Heading>Owned Sets</Heading>
            <div class="border-t border-t-zinc-600 pt-4">
                {#if data.user.owned.length > 0}
                    <div class="grid grid-cols-8 gap-4">
                        {#each data.user.owned as set}
                            <Set {set}></Set>
                        {/each}
                    </div>
                {:else}
                    <p>{data.user.user_name} has does not own any sets.</p>
                {/if}
            </div>
        </div>
        <div class="mb-2">
            <Heading>Subscribed Sets</Heading>
            <div class="border-t border-t-zinc-600 pt-4">
                {#if data.user.subscribed.length > 0}
                    <div class="grid grid-cols-8 gap-4">
                        {#each data.user.subscribed as set}
                            <Set {set}></Set>
                        {/each}
                    </div>
                {:else}
                    <p>{data.user.user_name} has not subscribed to any sets.</p>
                {/if}
            </div>
        </div>
    </div>
{:else if data.limited}
    <Meta title="Rate Limited" />

    <div class="w-fit m-auto text-center text-zinc-400">
        <p class="text-3xl mt-2"><i class="ri-slow-down-line"></i> Rate Limited</p>
        <p class="text-lg">Slow down! You're sending too many requests.</p>
    </div>
{:else}
    <Meta title="Not Found" />

    <div class="w-fit m-auto text-center text-zinc-400">
        <p class="text-3xl mt-2"><i class="ri-question-line"></i> Not Found</p>
        <p class="text-lg">That user was not found in the database.</p>
    </div>
{/if}
