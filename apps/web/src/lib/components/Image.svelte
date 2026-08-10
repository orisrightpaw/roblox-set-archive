<script lang="ts">
    // TODO
    // lazy loading, show spinner when loading and error when uhhh lets see Mayybeee uhh when it errors ok yea

    let { src, alt }: { src: string; alt?: string } = $props();
</script>

{#key src}
    <img class="rounded-md border border-zinc-700 aspect-square w-full mb-1" {alt} {src} />
{/key}
