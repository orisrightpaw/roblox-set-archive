<script lang="ts">
    import type { AssetSet } from '$lib';
    import Image from './Image.svelte';
    import Link from './Link.svelte';

    let { set }: { set: AssetSet } = $props();
</script>

<div>
    <a class="group" href="/sets/{set.id}">
        <Image src="https://web.archive.org/web/0im_/https://assetdelivery.roblox.com/v1/asset?id={set.image_asset_id}"></Image>
        <Link href="/sets/{set.id}">
            {set.name}
        </Link>
    </a>
    <p class="text-zinc-400 text-sm text-nowrap truncate">
        Creator: <Link href="/users/{set.creator_id}">{set.creator_name}</Link>
    </p>
    <p class="text-zinc-400 text-sm">
        Subscribers: <span class="text-red-400">{set.subscriber_count?.toLocaleString()}</span>
    </p>
</div>
