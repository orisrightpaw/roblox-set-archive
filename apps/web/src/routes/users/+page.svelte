<script lang="ts">
    import Meta from '$lib/components/Meta.svelte';
    import Search from '$lib/components/Search.svelte';
    import User from '$lib/components/User.svelte';
</script>

<Meta title="Search Users"></Meta>

{#snippet UserSnippet(result: any)}
    <User user={result}></User>
{/snippet}

<Search endpoint={'http://127.0.0.1:3000/api/users/search'} placeholder={'Search Users'} snippet={UserSnippet}></Search>
