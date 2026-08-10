<script lang="ts">
    // TODO
    // lazy loading, show spinner when loading and error when uhhh lets see Mayybeee uhh when it errors ok yea

    let {
        src,
        alt,
        className
    }: {
        src: string;
        alt?: string;
        className?: string;
    } = $props();
</script>

{#key src}
    <img class="rounded-md border border-zinc-700 aspect-square w-full mb-1 {className}" {alt} {src} />
{/key}
