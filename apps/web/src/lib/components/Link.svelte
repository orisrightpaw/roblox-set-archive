<script lang="ts">
    import type { Snippet } from 'svelte';

    let {
        href,
        children,
        className
    }: {
        href: string;
        children: Snippet;
        className?: string;
    } = $props();
</script>

<a class="text-blue-400 hover:text-blue-300 group-hover:text-blue-300 transition-colors {className}" {href}>
    {@render children()}
</a>
