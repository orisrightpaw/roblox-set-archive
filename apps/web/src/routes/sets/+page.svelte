<script lang="ts">
    import Meta from '$lib/components/Meta.svelte';
    import Search from '$lib/components/Search.svelte';
    import Set from '$lib/components/Set.svelte';
</script>

<Meta title="Search Sets"></Meta>

{#snippet AssetSet(result: any)}
    <Set set={result}></Set>
{/snippet}

<Search endpoint={'http://127.0.0.1:3000/api/sets/search'} placeholder={'Search Sets'} snippet={AssetSet}></Search>
