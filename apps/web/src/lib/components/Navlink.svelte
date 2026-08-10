<script lang="ts">
    import { page } from '$app/state';

    let { href, name }: { href: string; name: string } = $props();

    let active = $derived.by(() => {
        return page.url.pathname === href;
    });
</script>

<div class="inline-block mx-2">
    <a class="transition-colors {active ? 'text-zinc-100' : 'text-zinc-400 hover:text-zinc-300'}" {href}>{name}</a>
</div>
