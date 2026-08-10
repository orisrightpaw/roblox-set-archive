<script lang="ts">
    import Image from './Image.svelte';

    let { user }: { user: { id: number; user_name: string } } = $props();
</script>

<div class="group truncate text-blue-400 hover:text-blue-300">
    <a href="/users/{user.id}">
        <Image src="https://sets.starfall.wtf/api/users/{user.id}/thumbnail"></Image>
    </a>
    <a class="group-hover:text-blue-300 transition-colors" href="/users/{user.id}">
        {user.user_name}
    </a>
</div>
