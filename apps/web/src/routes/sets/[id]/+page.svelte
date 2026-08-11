<script lang="ts">
    import { page } from '$app/state';
    import Heading from '$lib/components/Heading.svelte';
    import Image from '$lib/components/Image.svelte';
    import Link from '$lib/components/Link.svelte';
    import Meta from '$lib/components/Meta.svelte';

    let { data } = $props();
</script>

{#if data.set}
    <Meta title="{data.set?.name!}, a Set by {data.set?.creator_name!}" />

    <div class="flex flex-col gap-4">
        <div class="flex gap-4">
            <div class="w-48">
                <Image
                    src="https://web.archive.org/web/0im_/https://assetdelivery.roblox.com/v1/asset?id={data.set
                        .image_asset_id}"
                ></Image>
            </div>
            <div>
                <Heading>{data.set.name}</Heading>
                <div class="flex flex-col">
                    <Link href="{page.url.origin}/Game/Tools/InsertAsset.ashx?sid={data.set.id}">See LuaWebService XML</Link>
                    <p class="text-zinc-400">Owner: <Link href="/users/{data.set.creator_id}">{data.set.creator_name}</Link></p>
                    <p class="text-zinc-400">
                        Last Image Update: <span class="text-zinc-100"
                            >{data.set.image_asset_updated
                                ? new Date(data.set.image_asset_updated * 1000).toLocaleString()
                                : 'Unknown'}</span
                        >
                    </p>
                </div>
                <p class="mt-2 whitespace-pre-wrap">
                    {data.set.description}
                </p>
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
        <p class="text-lg">That set was not found in the database.</p>
    </div>
{/if}
