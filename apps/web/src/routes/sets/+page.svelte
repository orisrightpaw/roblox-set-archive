<script lang="ts">
    import { API_HOST } from '$lib';
    import Meta from '$lib/components/Meta.svelte';
    import Search from '$lib/components/Search.svelte';
    import Set from '$lib/components/Set.svelte';
</script>

<Meta title="Search Sets"></Meta>

{#snippet AssetSet(result: any)}
    <Set set={result}></Set>
{/snippet}

<Search endpoint="{API_HOST}/api/sets/search" placeholder="Search Sets" snippet={AssetSet}></Search>
